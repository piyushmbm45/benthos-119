// Package amqp1 will eventually contain all implementations of AMQP 1
// components (that are currently within ./internal/old)
package amqp1

import (
	"github.com/Azure/go-amqp"

	"github.com/nehal119/benthos-119/pkg/impl/amqp1/shared"
)

func saslToOptFns(s shared.SASLConfig) ([]amqp.ConnOption, error) {
	switch s.Mechanism {
	case "plain":
		return []amqp.ConnOption{
			amqp.ConnSASLPlain(s.User, s.Password),
		}, nil
	case "none":
		return nil, nil
	}
	return nil, shared.ErrSASLMechanismNotSupported(s.Mechanism)
}
