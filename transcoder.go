package transcoder

import (
	"context"
	"io"
)

// Transcoder ...
type Transcoder interface {
	Start(opts Options) (<-chan Progress, error)
	Input(i string) Transcoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	Output(o string) Transcoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	WithOptions(opts Options) Transcoder
	WithAdditionalOptions(opts Options) Transcoder
	WithContext(ctx *context.Context) Transcoder
	WithMetadata(metadata Metadata) Transcoder
	SkipMetadata() Transcoder
	GetMetadata() (Metadata, error)
	Error() error
}
