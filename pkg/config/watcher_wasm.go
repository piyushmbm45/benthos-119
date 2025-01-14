//go:build wasm

package config

import (
	"errors"

	"github.com/nehal119/benthos-119/pkg/bundle"
)

// BeginFileWatching does nothing in WASM builds as it is not supported. Sorry!
func (r *Reader) BeginFileWatching(mgr bundle.NewManagement, strict bool) error {
	return errors.New("file watching is disabled in WASM builds")
}
