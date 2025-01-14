package output

import (
	"github.com/nehal119/benthos-119/pkg/batch/policy/batchconfig"
	sess "github.com/nehal119/benthos-119/pkg/impl/aws/session"
	"github.com/nehal119/benthos-119/pkg/old/util/retries"
)

// KinesisFirehoseConfig contains configuration fields for the KinesisFirehose output type.
type KinesisFirehoseConfig struct {
	SessionConfig  `json:",inline" yaml:",inline"`
	Stream         string `json:"stream" yaml:"stream"`
	MaxInFlight    int    `json:"max_in_flight" yaml:"max_in_flight"`
	retries.Config `json:",inline" yaml:",inline"`
	Batching       batchconfig.Config `json:"batching" yaml:"batching"`
}

// NewKinesisFirehoseConfig creates a new Config with default values.
func NewKinesisFirehoseConfig() KinesisFirehoseConfig {
	rConf := retries.NewConfig()
	rConf.Backoff.InitialInterval = "1s"
	rConf.Backoff.MaxInterval = "5s"
	rConf.Backoff.MaxElapsedTime = "30s"

	return KinesisFirehoseConfig{
		SessionConfig: SessionConfig{
			Config: sess.NewConfig(),
		},
		Stream:      "",
		MaxInFlight: 64,
		Config:      rConf,
		Batching:    batchconfig.NewConfig(),
	}
}
