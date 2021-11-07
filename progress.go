package transcoder

// Progress ...
type Progress interface {
	GetFramesProcessed() string
	GetCurrentTime() string
	GetCurrentBitrate() string
	GetProgress() float64
	GetSpeed() string
}
