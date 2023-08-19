package ffmpeg

import "github.com/floostack/transcoder"

// Metadata ...
type Metadata struct {
	Format  Format    `json:"format"`
	Streams []Streams `json:"streams"`
}

// Format ...
type Format struct {
	Filename       string
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           Tags   `json:"tags"`
}

// Streams ...
type Streams struct {
	Index              int               `json:"index"`
	ID                 string            `json:"id"`
	CodecName          string            `json:"codec_name"`
	CodecLongName      string            `json:"codec_long_name"`
	Profile            string            `json:"profile"`
	CodecType          string            `json:"codec_type"`
	CodecTimeBase      string            `json:"codec_time_base"`
	CodecTagString     string            `json:"codec_tag_string"`
	CodecTag           string            `json:"codec_tag"`
	Width              int               `json:"width,omitempty"`
	Height             int               `json:"height,omitempty"`
	CodedWidth         int               `json:"coded_width,omitempty"`
	CodedHeight        int               `json:"coded_height,omitempty"`
	ClosedCaptions     int               `json:"closed_captions,omitempty"`
	FilmGrain          int               `json:"film_grain,omitempty"`
	HasBFrames         int               `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string            `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string            `json:"display_aspect_ratio,omitempty"`
	PixFmt             string            `json:"pix_fmt,omitempty"`
	Level              int               `json:"level,omitempty"`
	ColorRange         string            `json:"color_range,omitempty"`
	ColorSpace         string            `json:"color_space,omitempty"`
	ColorTransfer      string            `json:"color_transfer,omitempty"`
	ColorPrimaries     string            `json:"color_primaries,omitempty"`
	ChromaLocation     string            `json:"chroma_location,omitempty"`
	FieldOrder         string            `json:"field_order,omitempty"`
	Refs               int               `json:"refs,omitempty"`
	QuarterSample      string            `json:"quarter_sample,omitempty"`
	DivxPacked         string            `json:"divx_packed,omitempty"`
	IsAvc              string            `json:"is_avc,omitempty"`
	NalLengthSize      string            `json:"nal_length_size,omitempty"`
	RFrameRate         string            `json:"r_frame_rate"`
	AvgFrameRate       string            `json:"avg_frame_rate"`
	TimeBase           string            `json:"time_base"`
	StartPts           int               `json:"start_pts"`
	StartTime          string            `json:"start_time"`
	DurationTs         int               `json:"duration_ts"`
	Duration           string            `json:"duration"`
	BitRate            string            `json:"bit_rate"`
	BitsPerRawSample   string            `json:"bits_per_raw_sample,omitempty"`
	NbFrames           int               `json:"nb_frames"`
	ExtradataSize      int               `json:"extradata_size"`
	Disposition        Disposition       `json:"disposition"`
	Tags               map[string]string `json:"tags,omitempty"`
	SampleFmt          string            `json:"sample_fmt,omitempty"`
	SampleRate         string            `json:"sample_rate,omitempty"`
	Channels           int               `json:"channels,omitempty"`
	ChannelLayout      string            `json:"channel_layout,omitempty"`
	BitsPerSample      int               `json:"bits_per_sample,omitempty"`
}

// Tags ...
type Tags struct {
	Encoder          string `json:"ENCODER"`
	MajorBrand       string `json:"major_brand"`
	MinorVersion     string `json:"minor_version"`
	CompatibleBrands string `json:"compatible_brands"`
	CreationTime     string `json:"creation_time"`
}

// Disposition ...
type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic,omitempty"`
	TimedThumbnails int `json:"timed_thumbnails,omitempty"`
	Captions        int `json:"captions,omitempty"`
	Descriptions    int `json:"descriptions,omitempty"`
	Metadata        int `json:"metadata,omitempty"`
	Dependent       int `json:"dependent,omitempty"`
	StillImage      int `json:"still_image,omitempty"`
}

// GetFormat ...
func (m Metadata) GetFormat() transcoder.Format {
	return m.Format
}

// GetStreams ...
func (m Metadata) GetStreams() (streams []transcoder.Streams) {
	for _, element := range m.Streams {
		streams = append(streams, element)
	}
	return streams
}

// GetFilename ...
func (f Format) GetFilename() string {
	return f.Filename
}

// GetNbStreams ...
func (f Format) GetNbStreams() int {
	return f.NbStreams
}

// GetNbPrograms ...
func (f Format) GetNbPrograms() int {
	return f.NbPrograms
}

// GetFormatName ...
func (f Format) GetFormatName() string {
	return f.FormatName
}

// GetFormatLongName ...
func (f Format) GetFormatLongName() string {
	return f.FormatLongName
}

// GetDuration ...
func (f Format) GetDuration() string {
	return f.Duration
}

// GetSize ...
func (f Format) GetSize() string {
	return f.Size
}

// GetBitRate ...
func (f Format) GetBitRate() string {
	return f.BitRate
}

// GetProbeScore ...
func (f Format) GetProbeScore() int {
	return f.ProbeScore
}

// GetTags ...
func (f Format) GetTags() transcoder.Tags {
	return f.Tags
}

// GetEncoder ...
func (t Tags) GetEncoder() string {
	return t.Encoder
}

// GetIndex ...
func (s Streams) GetIndex() int {
	return s.Index
}

// GetID ...
func (s Streams) GetID() string {
	return s.ID
}

// GetCodecName ...
func (s Streams) GetCodecName() string {
	return s.CodecName
}

// GetCodecLongName ...
func (s Streams) GetCodecLongName() string {
	return s.CodecLongName
}

// GetProfile ...
func (s Streams) GetProfile() string {
	return s.Profile
}

// GetCodecType ...
func (s Streams) GetCodecType() string {
	return s.CodecType
}

// GetCodecTimeBase ...
func (s Streams) GetCodecTimeBase() string {
	return s.CodecTimeBase
}

// GetCodecTagString ...
func (s Streams) GetCodecTagString() string {
	return s.CodecTagString
}

// GetCodecTag ...
func (s Streams) GetCodecTag() string {
	return s.CodecTag
}

// GetWidth ...
func (s Streams) GetWidth() int {
	return s.Width
}

// GetHeight ...
func (s Streams) GetHeight() int {
	return s.Height
}

// GetCodedWidth ...
func (s Streams) GetCodedWidth() int {
	return s.CodedWidth
}

// GetCodedHeight ...
func (s Streams) GetCodedHeight() int {
	return s.CodedHeight
}

// GetHasBFrames ...
func (s Streams) GetHasBFrames() int {
	return s.HasBFrames
}

// GetSampleAspectRatio ...
func (s Streams) GetSampleAspectRatio() string {
	return s.SampleAspectRatio
}

// GetDisplayAspectRatio ...
func (s Streams) GetDisplayAspectRatio() string {
	return s.DisplayAspectRatio
}

// GetPixFmt ...
func (s Streams) GetPixFmt() string {
	return s.PixFmt
}

// GetLevel ...
func (s Streams) GetLevel() int {
	return s.Level
}

// GetChromaLocation ...
func (s Streams) GetChromaLocation() string {
	return s.ChromaLocation
}

// GetRefs ...
func (s Streams) GetRefs() int {
	return s.Refs
}

// GetQuarterSample ...
func (s Streams) GetQuarterSample() string {
	return s.QuarterSample
}

// GetDivxPacked ...
func (s Streams) GetDivxPacked() string {
	return s.DivxPacked
}

// GetRFrameRate ...
func (s Streams) GetRFrameRate() string {
	return s.RFrameRate
}

// GetAvgFrameRate ...
func (s Streams) GetAvgFrameRate() string {
	return s.AvgFrameRate
}

// GetTimeBase ...
func (s Streams) GetTimeBase() string {
	return s.TimeBase
}

// GetDurationTs ...
func (s Streams) GetDurationTs() int {
	return s.DurationTs
}

// GetDuration ...
func (s Streams) GetDuration() string {
	return s.Duration
}

// GetDisposition ...
func (s Streams) GetDisposition() transcoder.Disposition {
	return s.Disposition
}

// GetBitRate ...
func (s Streams) GetBitRate() string {
	return s.BitRate
}

// GetDefault ...
func (d Disposition) GetDefault() int {
	return d.Default
}

// GetDub ...
func (d Disposition) GetDub() int {
	return d.Dub
}

// GetOriginal ...
func (d Disposition) GetOriginal() int {
	return d.Original
}

// GetComment ...
func (d Disposition) GetComment() int {
	return d.Comment
}

// GetLyrics ...
func (d Disposition) GetLyrics() int {
	return d.Lyrics
}

// GetKaraoke ...
func (d Disposition) GetKaraoke() int {
	return d.Karaoke
}

// GetForced ...
func (d Disposition) GetForced() int {
	return d.Forced
}

// GetHearingImpaired ...
func (d Disposition) GetHearingImpaired() int {
	return d.HearingImpaired
}

// GetVisualImpaired ...
func (d Disposition) GetVisualImpaired() int {
	return d.VisualImpaired
}

// GetCleanEffects ...
func (d Disposition) GetCleanEffects() int {
	return d.CleanEffects
}
