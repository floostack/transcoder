package transcoder

// Metadata ...
type Metadata interface {
	GetFormat() Format
	GetStreams() []Streams
}

// Format ...
type Format interface {
	GetFilename() string
	GetNbStreams() int
	GetNbPrograms() int
	GetFormatName() string
	GetFormatLongName() string
	GetDuration() string
	GetSize() string
	GetBitRate() string
	GetProbeScore() int
	GetTags() Tags
}

// Streams ...
type Streams interface {
	GetIndex() int
	GetID() string
	GetCodecName() string
	GetCodecLongName() string
	GetProfile() string
	GetCodecType() string
	GetCodecTimeBase() string
	GetCodecTagString() string
	GetCodecTag() string
	GetWidth() int
	GetHeight() int
	GetCodedWidth() int
	GetCodedHeight() int
	GetHasBFrames() int
	GetSampleAspectRatio() string
	GetDisplayAspectRatio() string
	GetPixFmt() string
	GetLevel() int
	GetChromaLocation() string
	GetRefs() int
	GetQuarterSample() string
	GetDivxPacked() string
	GetRFrameRrate() string
	GetAvgFrameRate() string
	GetTimeBase() string
	GetDurationTs() int
	GetDuration() string
	GetDisposition() Disposition
	GetBitRate() string
}

// Tags ...
type Tags interface {
	GetEncoder() string
}

// Disposition ...
type Disposition interface {
	GetDefault() int
	GetDub() int
	GetOriginal() int
	GetComment() int
	GetLyrics() int
	GetKaraoke() int
	GetForced() int
	GetHearingImpaired() int
	GetVisualImpaired() int
	GetCleanEffects() int
}
