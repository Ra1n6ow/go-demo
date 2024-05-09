package zerolog_test

import (
	"github.com/ra1n6ow/go-demo/log/zerolog"
	"testing"
)

func TestSimple(t *testing.T) {
	zerolog.Sample()
}

func TestContextual(t *testing.T) {
	zerolog.Contextual()
}
