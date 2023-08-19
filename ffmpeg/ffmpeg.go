package ffmpeg

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/floostack/transcoder"
	"github.com/floostack/transcoder/utils"
)

// Transcoder ...
type Transcoder struct {
	config           *Config
	input            string
	output           []string
	inputOptions     []string
	outputOptions    [][]string
	metadata         transcoder.Metadata
	inputPipeReader  *io.ReadCloser
	outputPipeWriter *io.ReadWriteCloser
	commandContext   *context.Context
}

// New ...
func New(cfg *Config) transcoder.Transcoder {
	return &Transcoder{config: cfg}
}

// Start ...
func (t *Transcoder) Start() (<-chan transcoder.Progress, error) {

	var stderrIn io.ReadCloser

	out := make(chan transcoder.Progress)

	defer t.closePipes()

	// Validates config
	if err := t.validate(); err != nil {
		return nil, err
	}

	// Get file metadata
	_, err := t.GetMetadata()
	if err != nil {
		return nil, err
	}

	// Append input file and standard options
	var args []string

	if len(t.inputOptions) > 0 {
		args = append(args, t.inputOptions...)
	}

	args = append(args, []string{"-i", t.input}...)
	outputLength := len(t.output)
	outputOptionsLength := len(t.outputOptions)

	if outputLength == 1 && outputOptionsLength == 0 {
		// Just append the 1 output file we've got
		args = append(args, t.output[0])
	} else {
		for index, out := range t.output {
			// Get executable flags
			// If we are at the last output file but still have several options, append them all at once
			if index == outputLength-1 && outputLength < outputOptionsLength {
				for i := index; i < len(t.outputOptions); i++ {
					args = append(args, t.outputOptions[i]...)
				}
				// Otherwise just append the current options
			} else {
				args = append(args, t.outputOptions[index]...)
			}

			// Append output flag
			args = append(args, out)
		}
	}

	// Initialize command
	// If a context object was supplied to this Transcoder before
	// starting, use this context when creating the command to allow
	// the command to be killed when the context expires
	var cmd *exec.Cmd
	if t.commandContext == nil {
		cmd = exec.Command(t.config.FfmpegBinPath, args...)
	} else {
		cmd = exec.CommandContext(*t.commandContext, t.config.FfmpegBinPath, args...)
	}

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

		go func() {
			defer close(out)
			err = cmd.Wait()
		}()
	} else {
		err = cmd.Wait()
	}

	return out, nil
}

// Input ...
func (t *Transcoder) Input(arg string) transcoder.Transcoder {
	t.input = arg
	return t
}

// Output ...
func (t *Transcoder) Output(arg string) transcoder.Transcoder {
	t.output = append(t.output, arg)
	return t
}

// InputPipe ...
func (t *Transcoder) InputPipe(r *io.ReadCloser) transcoder.Transcoder {
	t.input = "pipe:"
	t.inputPipeReader = r
	return t
}

// OutputPipe ...
func (t *Transcoder) OutputPipe(w *io.ReadWriteCloser) transcoder.Transcoder {
	t.output = []string{}
	t.outputPipeWriter = w
	return t
}

// WithInputOptions Sets the options object
func (t *Transcoder) WithInputOptions(opts transcoder.Options) transcoder.Transcoder {
	t.inputOptions = opts.GetStrArguments()
	return t
}

// WithAdditionalInputOptions Appends an additional options object
func (t *Transcoder) WithAdditionalInputOptions(opts transcoder.Options) transcoder.Transcoder {
	t.inputOptions = append(t.inputOptions, opts.GetStrArguments()...)
	return t
}

// WithOutputOptions Sets the options object
func (t *Transcoder) WithOutputOptions(opts transcoder.Options) transcoder.Transcoder {
	t.outputOptions = [][]string{opts.GetStrArguments()}
	return t
}

// WithAdditionalOutputOptions Appends an additional options object
func (t *Transcoder) WithAdditionalOutputOptions(opts transcoder.Options) transcoder.Transcoder {
	t.outputOptions = append(t.outputOptions, opts.GetStrArguments())
	return t
}

// WithContext is to be used on a Transcoder *before Starting* to
// pass in a context.Context object that can be used to kill
// a running transcoder process. Usage of this method is optional
func (t *Transcoder) WithContext(ctx *context.Context) transcoder.Transcoder {
	t.commandContext = ctx
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

	outputLength := len(t.output)

	if outputLength == 0 {
		return errors.New("missing output option")
	}

	// length of output files being greater than length of options would produce an invalid ffmpeg command
	// unless there is only 1 output file, which obviously wouldn't be a problem
	if outputLength > len(t.outputOptions) && outputLength != 1 {
		return errors.New("number of options and output files does not match")
	}

	for index, output := range t.output {
		if output == "" {
			return fmt.Errorf("output at index %d is an empty string", index)
		}
	}

	return nil
}

// GetMetadata Returns metadata for the specified input file
func (t *Transcoder) GetMetadata() (transcoder.Metadata, error) {

	if t.config.FfprobeBinPath != "" {
		var outb, errb bytes.Buffer

		input := t.input

		args := []string{"-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}

		cmd := exec.Command(t.config.FfprobeBinPath, args...)
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		if t.inputPipeReader != nil {
			cmd.Stdin = *t.inputPipeReader
		}
		if t.outputPipeWriter != nil {
			cmd.Stdout = *t.outputPipeWriter
		}

		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("error executing (%s) with args (%s) | error: %s | message: %s %s", t.config.FfprobeBinPath, args, err, outb.String(), errb.String())
		}

		var metadata Metadata

		if t.outputPipeWriter == nil {
			if err = json.Unmarshal(outb.Bytes(), &metadata); err != nil {
				return nil, err
			}
			t.metadata = metadata
		}

		return metadata, nil
	}

	return nil, errors.New("ffprobe binary not found")
}

// progress sends through given channel the transcoding status
func (t *Transcoder) progress(stream io.ReadCloser, out chan transcoder.Progress) {

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

		if strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
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
			dursec, _ := strconv.ParseFloat(t.metadata.GetFormat().GetDuration(), 64)

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
