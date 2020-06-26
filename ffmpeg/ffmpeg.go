package ffmpeg

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/floostack/gotrans"
	"github.com/floostack/gotrans/utils"
)

// Transcoder ...
type Transcoder struct {
	config           *Config
	input            string
	output           string
	options          []string
	metadata         *Metadata
	inputPipeReader  *io.ReadCloser
	outputPipeReader *io.ReadCloser
	inputPipeWriter  *io.WriteCloser
	outputPipeWriter *io.WriteCloser
}

// New ...
func New(cfg *Config) gotrans.Transcoder {
	return &Transcoder{config: cfg}
}

// Start ...
func (t *Transcoder) Start(opts gotrans.Options) (<-chan gotrans.Progress, error) {

	var stderrIn io.ReadCloser

	out := make(chan gotrans.Progress)

	defer t.closePipes()

	// Validates config
	if err := t.validate(); err != nil {
		return nil, err
	}

	// Get file metadata
	_, err := t.getMetadata()
	if err != nil {
		return nil, err
	}

	// Get executable flags
	args := append([]string{"-i", t.input}, opts.GetStrArguments()...)

	// Append output flag
	args = append(args, []string{t.output}...)

	// Initialize command
	cmd := exec.Command(t.config.FfmpegBinPath, args...)

	// cmd.Stderr = os.Stdout

	// If progresss enabled, get stderr pipe and start progress process
	if t.config.ProgressEnabled && !t.config.Verbose {
		stderrIn, err = cmd.StderrPipe()
		if err != nil {
			return nil, fmt.Errorf("Failed getting transcoding progress (%s) with args (%s) with error %s", t.config.FfmpegBinPath, args, err)
		}
	}

	if t.config.Verbose {
		cmd.Stderr = os.Stdout
	}

	// Start process
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("Failed starting transcoding (%s) with args (%s) with error %s", t.config.FfmpegBinPath, args, err)
	}

	if t.config.ProgressEnabled && !t.config.Verbose {
		go func() {
			t.progress(stderrIn, out)
		}()
	}

	go func() {
		defer close(out)
		// Start process
		err = cmd.Wait()
	}()

	return out, nil
}

// Input ...
func (t *Transcoder) Input(arg string) gotrans.Transcoder {
	t.input = arg
	return t
}

// Output ...
func (t *Transcoder) Output(arg string) gotrans.Transcoder {
	t.output = arg
	return t
}

// InputPipe ...
func (t *Transcoder) InputPipe(w *io.WriteCloser, r *io.ReadCloser) gotrans.Transcoder {
	if &t.input == nil {
		t.inputPipeWriter = w
		t.inputPipeReader = r
	}
	return t
}

// OutputPipe ...
func (t *Transcoder) OutputPipe(w *io.WriteCloser, r *io.ReadCloser) gotrans.Transcoder {
	if &t.output == nil {
		t.outputPipeWriter = w
		t.outputPipeReader = r
	}
	return t
}

// WithOptions ...
func (t *Transcoder) WithOptions(opts gotrans.Options) gotrans.Transcoder {
	t.options = opts.GetStrArguments()
	return t
}

// validate ...
func (t *Transcoder) validate() error {
	if t.config.FfmpegBinPath == "" {
		return errors.New("ffmpeg binary path not found")
	}

	if t.input == "" {
		return errors.New("missing input option")
	}

	if t.output == "" {
		return errors.New("missing output option")
	}

	return nil
}

func (t *Transcoder) getMetadata() (metadata *Metadata, err error) {

	if t.config.FfprobeBinPath != "" {
		var outb, errb bytes.Buffer

		input := t.input

		if t.inputPipeReader != nil {
			input = "pipe:"
		}

		args := []string{"-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}

		cmd := exec.Command(t.config.FfprobeBinPath, args...)
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("error executing (%s) with args (%s) | error: %s | message: %s %s", t.config.FfprobeBinPath, args, err, outb.String(), errb.String())
		}

		if err = json.Unmarshal([]byte(outb.String()), &metadata); err != nil {
			return nil, err
		}

		t.metadata = metadata

		return metadata, nil
	}

	return nil, errors.New("ffprobe binary not found")
}

// progress sends through given channel the transcoding status
func (t *Transcoder) progress(stream io.ReadCloser, out chan gotrans.Progress) {

	defer stream.Close()

	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	scanner := bufio.NewScanner(stream)
	scanner.Split(split)

	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		Progress := new(Progress)
		line := scanner.Text()

		if strings.Contains(line, "frame=") && strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
			var re = regexp.MustCompile(`=\s+`)
			st := re.ReplaceAllString(line, `=`)

			f := strings.Fields(st)

			var framesProcessed string
			var currentTime string
			var currentBitrate string
			var currentSpeed string

			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")

				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]

					if fieldname == "frame" {
						framesProcessed = fieldvalue
					}

					if fieldname == "time" {
						currentTime = fieldvalue
					}

					if fieldname == "bitrate" {
						currentBitrate = fieldvalue
					}
					if fieldname == "speed" {
						currentSpeed = fieldvalue
					}
				}
			}

			timesec := utils.DurToSec(currentTime)
			dursec, _ := strconv.ParseFloat(t.metadata.Format.Duration, 64)

			progress := (timesec * 100) / dursec
			Progress.Progress = progress

			Progress.CurrentBitrate = currentBitrate
			Progress.FramesProcessed = framesProcessed
			Progress.CurrentTime = currentTime
			Progress.Speed = currentSpeed

			out <- *Progress
		}
	}
}

// closePipes Closes pipes if opened
func (t *Transcoder) closePipes() {
	if t.inputPipeReader != nil {
		ipr := *t.inputPipeReader
		ipr.Close()
	}

	if t.outputPipeWriter != nil {
		opr := *t.outputPipeWriter
		opr.Close()
	}
}
