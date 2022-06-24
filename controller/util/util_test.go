package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsProcessing(t *testing.T) {
	result := IsProcessing("2022-06-09 11:20 39")
	assert.True(t, result)
}
