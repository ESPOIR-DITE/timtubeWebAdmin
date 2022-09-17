package channel_type

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"timtubeWebAdmin/domain"
)

func TestCreateChannelType(t *testing.T) {
	object := domain.ChannelType{"", "Sport", "Relative to Sport news"}
	result, err := CreateChannelType(object)
	assert.Nil(t, err)
	fmt.Println(result)

}
