package snowflake

import (
	_ "embed"
	"testing"
)

//go:embed testdata/script.js
var script string

func TestIntegration(t *testing.T) { //nolint:paralleltest
	t.Skip()
}
