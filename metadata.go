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
	GetStartTime() string
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
	GetClosedCaptions() int
	GetFilmGrain() int
	GetHasBFrames() int
	GetSampleAspectRatio() string
	GetDisplayAspectRatio() string
	GetPixFmt() string
	GetLevel() int
	GetChromaLocation() string
	GetRefs() int
	GetQuarterSample() string
	GetDivxPacked() string
	GetIsAvc() string
	GetNalLengthSize() string
	GetRFrameRate() string
	GetAvgFrameRate() string
	GetTimeBase() string
	GetStartPts() int
	GetStartTime() string
	GetDurationTs() int
	GetDuration() string
	GetBitRate() string
	GetBitsPerRawSample() string
	GetNbFrames() int
	GetExtradataSize() int
	GetDisposition() Disposition
	GetSampleFmt() string
	GetSampleRate() string
	GetChannels() int
	GetChannelLayout() string
	GetBitsPerSample() int
}

// Tags ...
type Tags interface {
	GetEncoder() string
	GetMajorBrand() string
	GetMinorVersion() string
	GetCompatibleBrands() string
	GetCreationTime() string
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
	GetAttachedPic() int
	GetTimedThumbnails() int
	GetCaptions() int
	GetMetadata() int
	GetDependent() int
	GetStillImage() int
}
