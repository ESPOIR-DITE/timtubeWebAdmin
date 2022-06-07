package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"timtubeWebAdmin/domain"
)

func TestCreateUserVideo(t *testing.T) {
	User := domain.UserVideo{"", "029324", "485345", "today"}
	result, err := CreateUserVideo(User)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadAllOfUserVideo(t *testing.T) {
	result, err := ReadAllOfUserVideo("029324")
	assert.Nil(t, err)
	fmt.Println(result)
}
