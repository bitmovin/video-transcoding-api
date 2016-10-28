// Package zencoder provides a implementation of the provider that uses the
// Zencoder API for transcoding media files.
//
// It doesn't expose any public type. In order to use the provider, one must
// import this package and then grab the factory from the provider package:
//
//     import (
//         "github.com/NYTimes/video-transcoding-api/provider"
//         "github.com/NYTimes/video-transcoding-api/provider/zencoder"
//     )
//
//     func UseProvider() {
//         factory, err := provider.GetProviderFactory(Zencoder.Name)
//         // handle err and use factory to get an instance of the provider.
//     }
package zencoder

import (
	"github.com/NYTimes/video-transcoding-api/config"
	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/NYTimes/video-transcoding-api/provider"
	"github.com/brandscreen/zencoder"
)

// Name is the name used for registering the Encoding.com provider in the
// registry of providers.
const Name = "zencoder"

var errZencoderInvalidConfig = provider.InvalidConfigError("missing Zencoder API key. Please define the environment variables ZENCODER_API_KEY or set these values in the configuration file")

func init() {
	provider.Register(Name, zencoderFactory)
}

type zencoderProvider struct {
	config *config.Config
	client *zencoder.Zencoder
}

func (z *zencoderProvider) Transcode(job *db.Job, transcodeProfile provider.TranscodeProfile) (*provider.JobStatus, error) {
	return &provider.JobStatus{}, nil
}

func (z *zencoderProvider) CreatePreset(preset db.Preset) (string, error) {
	return "", nil
}

func (z *zencoderProvider) GetPreset(presetID string) (interface{}, error) {
	return "", nil
}

func (z *zencoderProvider) DeletePreset(presetID string) error {
	return nil
}

func (z *zencoderProvider) JobStatus(job *db.Job) (*provider.JobStatus, error) {
	return &provider.JobStatus{}, nil
}

func (z *zencoderProvider) CancelJob(id string) error {
	return nil
}

func (z *zencoderProvider) Healthcheck() error {
	return nil
}

func (z *zencoderProvider) Capabilities() provider.Capabilities {
	return provider.Capabilities{
		InputFormats:  []string{"prores", "h264"},
		OutputFormats: []string{"mp4", "hls", "webm"},
		Destinations:  []string{"akamai", "s3"},
	}
}

func zencoderFactory(cfg *config.Config) (provider.TranscodingProvider, error) {
	if cfg.Zencoder.APIKey == "" {
		return nil, errZencoderInvalidConfig
	}
	client := zencoder.NewZencoder(cfg.Zencoder.APIKey)
	return &zencoderProvider{client: client, config: cfg}, nil
}
