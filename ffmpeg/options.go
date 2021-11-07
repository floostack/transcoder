package ffmpeg

// Options defines allowed FFmpeg arguments
type Options []string

// GetStrArguments ...
func (opts Options) GetStrArguments() []string {
	return opts
}
