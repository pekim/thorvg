package cvalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructs(t *testing.T) {
	for _, size := range sizes {
		t.Run(size.name, func(t *testing.T) {
			assert.Equal(t, size.c, int(size.go_))
		})
	}
}
