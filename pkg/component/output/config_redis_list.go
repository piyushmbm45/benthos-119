package output

import (
	"github.com/nehal119/benthos-119/pkg/batch/policy/batchconfig"
	bredis "github.com/nehal119/benthos-119/pkg/impl/redis/old"
)

// RedisListConfig contains configuration fields for the RedisList output type.
type RedisListConfig struct {
	bredis.Config `json:",inline" yaml:",inline"`
	Key           string             `json:"key" yaml:"key"`
	MaxInFlight   int                `json:"max_in_flight" yaml:"max_in_flight"`
	Batching      batchconfig.Config `json:"batching" yaml:"batching"`
}

// NewRedisListConfig creates a new RedisListConfig with default values.
func NewRedisListConfig() RedisListConfig {
	return RedisListConfig{
		Config:      bredis.NewConfig(),
		Key:         "",
		MaxInFlight: 64,
		Batching:    batchconfig.NewConfig(),
	}
}
