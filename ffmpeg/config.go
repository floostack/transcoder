package ffmpeg

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Flags ...
type Flags struct {
	Progress bool `json:"progress,omitempty"`
	Verbose  bool `json:"verbose,omitempty"`
	Debug    bool `json:"debug,omitempty"`
}

// Config ...
type Config struct {
	Flags

	// Paths
	FfmpegBinPath  string `json:"ffmpeg_bin_path"`
	FfprobeBinPath string `json:"ffprobe_bin_path"`
}

func mergeFlags(flags ...Flags) (result Flags) {
	for _, v := range flags {
		result.Progress = result.Progress || v.Progress
		result.Verbose = result.Verbose || v.Verbose
		result.Debug = result.Debug || v.Debug
	}
	return
}

func NewAutoConfig(flags ...Flags) (*Config, error) {
	out, which := bytes.Buffer{}, exec.Command("which", "ffmpeg")
	which.Stdout = &out
	if err := which.Run(); err != nil {
		return nil, err
	}
	ffmpegBinPath := strings.TrimSpace(out.String())
	out.Reset()
	if len(ffmpegBinPath) < 1 {
		return nil, errors.New("ffmpeg binary path not found")
	}
	which = exec.Command("which", "ffprobe")
	which.Stdout = &out
	if err := which.Run(); err != nil {
		return nil, err
	}
	ffprobeBinPath := strings.TrimSpace(out.String())
	out.Reset()
	if len(ffprobeBinPath) < 1 {
		return nil, errors.New("ffprobe binary path not found")
	}
	return &Config{
		Flags:          mergeFlags(flags...),
		FfmpegBinPath:  ffmpegBinPath,
		FfprobeBinPath: ffprobeBinPath,
	}, nil
}
