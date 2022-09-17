package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsProcessing(t *testing.T) {
	result := IsProcessing("2022-06-09 11:20 39")
	assert.True(t, result)
}
func TestGetCurrentDate(t *testing.T) {
	fmt.Println(GetCurrentDate())
}
func TestGetStringDateTime(t *testing.T) {
	fmt.Println(GetStringDateTime("2022-06-28T00:00:00+02:00"))
}
