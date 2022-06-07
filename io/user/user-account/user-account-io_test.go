package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"timtubeWebAdmin/domain"
)

func TestReadUserAccounts(t *testing.T) {
	result, err := ReadUserAccounts()
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestCreateUserAccount(t *testing.T) {
	result, err := CreateUserAccount(domain.UserAccount{"", "ditekemena@gmai.com", "1234", "", true})
	assert.Nil(t, err)
	fmt.Println(result)
}
