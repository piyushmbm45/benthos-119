// Package base brings in only the sql components, but none of the drivers for
// them. It is up to you to import specifically the drivers you want to include.
package base

import (
	// Bring in the internal plugin definitions.
	_ "github.com/nehal119/benthos-119/pkg/impl/sql"
)
