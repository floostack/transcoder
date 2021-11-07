package ffmpeg

import (
	"encoding/json"
)

// Progress ...
type Progress struct {
	FramesProcessed string  `json:"f"`
	CurrentTime     string  `json:"t"`
	CurrentBitrate  string  `json:"b"`
	Speed           string  `json:"s"`
	Progress        float64 `json:"p"`
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

func (p Progress) String() string {
	data, _ := json.Marshal(&p)
	if data == nil {
		return "{}"
	}
	return string(data)
}
