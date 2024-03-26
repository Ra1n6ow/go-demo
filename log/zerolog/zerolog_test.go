package zerolog_test

import (
	"github.com/l1ghtd/go-demo/log/zerolog"
	"testing"
)

func TestSimple(t *testing.T) {
	zerolog.Sample()
}

func TestContextual(t *testing.T) {
	zerolog.Contextual()
}
