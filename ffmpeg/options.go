package ffmpeg

import (
	"fmt"
	"reflect"
)

// Options defines allowed FFmpeg arguments
type Options struct {
	Aspect                *string           `flag:"-aspect"`
	Resolution            *string           `flag:"-s"`
	VideoBitRate          *string           `flag:"-b:v"`
	VideoBitRateTolerance *int              `flag:"-bt"`
	VideoMaxBitRate       *int              `flag:"-maxrate"`
	VideoMinBitrate       *int              `flag:"-minrate"`
	VideoCodec            *string           `flag:"-c:v"`
	Vframes               *int              `flag:"-vframes"`
	FrameRate             *int              `flag:"-r"`
	AudioRate             *int              `flag:"-ar"`
	KeyframeInterval      *int              `flag:"-g"`
	AudioCodec            *string           `flag:"-c:a"`
	AudioBitrate          *string           `flag:"-ab"`
	AudioChannels         *int              `flag:"-ac"`
	AudioVariableBitrate  *bool             `flag:"-q:a"`
	BufferSize            *int              `flag:"-bufsize"`
	Threadset             *bool             `flag:"-threads"`
	Threads               *int              `flag:"-threads"`
	Preset                *string           `flag:"-preset"`
	Tune                  *string           `flag:"-tune"`
	AudioProfile          *string           `flag:"-profile:a"`
	VideoProfile          *string           `flag:"-profile:v"`
	Target                *string           `flag:"-target"`
	Duration              *string           `flag:"-t"`
	Qscale                *uint32           `flag:"-qscale"`
	Crf                   *uint32           `flag:"-crf"`
	Strict                *int              `flag:"-strict"`
	MuxDelay              *string           `flag:"-muxdelay"`
	SeekTime              *string           `flag:"-ss"`
	SeekUsingTimestamp    *bool             `flag:"-seek_timestamp"`
	MovFlags              *string           `flag:"-movflags"`
	HideBanner            *bool             `flag:"-hide_banner"`
	OutputFormat          *string           `flag:"-f"`
	CopyTs                *bool             `flag:"-copyts"`
	NativeFramerateInput  *bool             `flag:"-re"`
	InputInitialOffset    *string           `flag:"-itsoffset"`
	RtmpLive              *string           `flag:"-rtmp_live"`
	HlsPlaylistType       *string           `flag:"-hls_playlist_type"`
	HlsListSize           *int              `flag:"-hls_list_size"`
	HlsSegmentDuration    *int              `flag:"-hls_time"`
	HlsMasterPlaylistName *string           `flag:"-master_pl_name"`
	HlsSegmentFilename    *string           `flag:"-hls_segment_filename"`
	HTTPMethod            *string           `flag:"-method"`
	HTTPKeepAlive         *bool             `flag:"-multiple_requests"`
	Hwaccel               *string           `flag:"-hwaccel"`
	StreamIds             map[string]string `flag:"-streamid"`
	VideoFilter           *string           `flag:"-vf"`
	AudioFilter           *string           `flag:"-af"`
	SkipVideo             *bool             `flag:"-vn"`
	SkipAudio             *bool             `flag:"-an"`
	CompressionLevel      *int              `flag:"-compression_level"`
	MapMetadata           *string           `flag:"-map_metadata"`
	Metadata              map[string]string `flag:"-metadata"`
	EncryptionKey         *string           `flag:"-hls_key_info_file"`
	Bframe                *int              `flag:"-bf"`
	PixFmt                *string           `flag:"-pix_fmt"`
	WhiteListProtocols    []string          `flag:"-protocol_whitelist"`
	Overwrite             *bool             `flag:"-y"`
	ExtraArgs             map[string]interface{}
}

// GetStrArguments ...
func (opts Options) GetStrArguments() []string {
	f := reflect.TypeOf(opts)
	v := reflect.ValueOf(opts)

	values := []string{}

	for i := 0; i < f.NumField(); i++ {
		flag := f.Field(i).Tag.Get("flag")
		value := v.Field(i).Interface()

		if !v.Field(i).IsNil() {

			if _, ok := value.(*bool); ok {
				values = append(values, flag)
			}

			if vs, ok := value.(*string); ok {
				values = append(values, flag, *vs)
			}

			if va, ok := value.([]string); ok {

				for i := 0; i < len(va); i++ {
					item := va[i]
					values = append(values, flag, item)
				}
			}

			if vm, ok := value.(map[string]string); ok {
				for k, v := range vm {
					values = append(values, flag, fmt.Sprintf("%v:%v", k, v))
				}
			}
			
			if vi, ok := value.(*int); ok {
				values = append(values, flag, fmt.Sprintf("%d", *vi))
			}

		}
	}

	return values
}

// Metadata ...
type Metadata struct {
	Streams []Streams `json:"streams"`
	Format  Format    `json:"format"`
}

// Streams defines allowed stream options
type Streams struct {
	Index              int
	ID                 string      `json:"id"`
	CodecName          string      `json:"codec_name"`
	CodecLongName      string      `json:"codec_long_name"`
	Profile            string      `json:"profile"`
	CodecType          string      `json:"codec_type"`
	CodecTimeBase      string      `json:"codec_time_base"`
	CodecTagString     string      `json:"codec_tag_string"`
	CodecTag           string      `json:"codec_tag"`
	Width              int         `json:"width"`
	Height             int         `json:"height"`
	CodedWidth         int         `json:"coded_width"`
	CodedHeight        int         `json:"coded_height"`
	HasBFrames         int         `json:"has_b_frames"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio"`
	DisplayAspectRatio string      `json:"display_aspect_ratio"`
	PixFmt             string      `json:"pix_fmt"`
	Level              int         `json:"level"`
	ChromaLocation     string      `json:"chroma_location"`
	Refs               int         `json:"refs"`
	QuarterSample      string      `json:"quarter_sample"`
	DivxPacked         string      `json:"divx_packed"`
	RFrameRrate        string      `json:"r_frame_rate"`
	AvgFrameRate       string      `json:"avg_frame_rate"`
	TimeBase           string      `json:"time_base"`
	DurationTs         int         `json:"duration_ts"`
	Duration           string      `json:"duration"`
	Disposition        Disposition `json:"disposition"`
	BitRate            string      `json:"bit_rate"`
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
}

// Format ...
type Format struct {
	Filename       string
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           Tags   `json:"tags"`
}

// Tags ...
type Tags struct {
	Encoder string `json:"ENCODER"`
}
