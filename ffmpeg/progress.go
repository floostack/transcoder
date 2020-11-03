package ffmpeg

// Progress ...
type Progress struct {
	FramesProcessed string
	CurrentTime     string
	CurrentBitrate  string
	Progress        float64
	Speed           string
}

// GetFramesProcessed ...
func (p Progress) GetFramesProcessed() string {
	return p.FramesProcessed
}

// GetCurrentTime ...
func (p Progress) GetCurrentTime() string {
	return p.CurrentTime
}

// GetCurrentBitrate ...
func (p Progress) GetCurrentBitrate() string {
	return p.CurrentBitrate
}

// GetProgress ...
func (p Progress) GetProgress() float64 {
	return p.Progress
}

// GetSpeed ...
func (p Progress) GetSpeed() string {
	return p.Speed
}
