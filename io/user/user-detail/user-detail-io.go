package user

import (
	"errors"
	"timtubeWebAdmin/api"
	user_details "timtubeWebAdmin/domain/user/user-details"
)

const userURL = api.BASE_URL + "user/user-detail/"

func CreateUserDetails(use user_details.UserDetails) (user_details.UserDetails, error) {
	entity := user_details.UserDetails{}
	resp, _ := api.Rest().SetBody(use).Post(userURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserDetails(user user_details.UserDetails) (user_details.UserDetails, error) {
	entity := user_details.UserDetails{}
	resp, _ := api.Rest().SetBody(user).Post(userURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadUserDetailById(id string) (user_details.UserDetails, error) {
	entity := user_details.UserDetails{}
	resp, _ := api.Rest().Get(userURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserDetailByEmail(email string) (user_details.UserDetails, error) {
	entity := user_details.UserDetails{}
	resp, _ := api.Rest().Get(userURL + "getWithEmail/" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func IsExistUserDetailEmail(email string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "isExist-by-email/" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, err
	}
	return entity, nil
}
func IsExistUserDetailId(id string) bool {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "isExist-by-id/" + id)
	if resp.IsError() {
		return entity
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity
	}
	return entity
}
func DeleteUserDetailId(id string) bool {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "delete/" + id)
	if resp.IsError() {
		return entity
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity
	}
	return entity
}
func DeleteUserDetailEmail(id string) bool {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "get-by-email/" + id)
	if resp.IsError() {
		return entity
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity
	}
	return entity
}
func ReadUsers() ([]user_details.UserDetails, error) {
	entity := []user_details.UserDetails{}
	resp, _ := api.Rest().Get(userURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
