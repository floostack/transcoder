package transcoder

import (
	"context"
	"io"
)

// Transcoder ...
type Transcoder interface {
	Start() (<-chan Progress, error)
	Input(i string) Transcoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	Output(o string) Transcoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	WithInputOptions(opts Options) Transcoder
	WithAdditionalInputOptions(opts Options) Transcoder
	WithOutputOptions(opts Options) Transcoder
	WithAdditionalOutputOptions(opts Options) Transcoder
	WithContext(ctx *context.Context) Transcoder
	GetMetadata() (Metadata, error)
}
