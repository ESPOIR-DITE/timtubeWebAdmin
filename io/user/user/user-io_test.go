package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"timtubeWebAdmin/domain"
)

func TestReadUsers(t *testing.T) {
	result, err := ReadUsers()
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestCreateUser(t *testing.T) {
	result, err := CreateUser(domain.User{"ditekemena@gmai.com", "Espoir", "Diteekemena", "", "011"})
	assert.Nil(t, err)
	fmt.Println(result)
}
