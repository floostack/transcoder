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