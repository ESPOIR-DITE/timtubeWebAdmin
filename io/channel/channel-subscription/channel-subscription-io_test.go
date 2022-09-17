package user

import (
	"fmt"
	"testing"
)

func TestReadChannelSubscriptionByChannel(t *testing.T) {
	result, err := ReadChannelSubscriptionByChannel("UA-260a1663-a62a-4eee-8a23-ffad4cd90647")
	fmt.Println(err)
	fmt.Println(result)
}
