package user

import (
	"errors"
	"timtubeWebAdmin/api"
	user_details "timtubeWebAdmin/domain/user/user-details"
)

const userURL = api.BASE_URL + "user/user-bank/"

func CreateUserBank(use user_details.UserBank) (user_details.UserBank, error) {
	entity := user_details.UserBank{}
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
func UpdateUserBank(user user_details.UserBank) (user_details.UserBank, error) {
	entity := user_details.UserBank{}
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
func IsExistUserBankEmail(email string) bool {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "isExist-by-email/" + email)
	if resp.IsError() {
		return entity
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity
	}
	return entity
}
func IsExistUserBankId(id string) bool {
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
func ReadUserBankById(id string) (user_details.UserBank, error) {
	entity := user_details.UserBank{}
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
func ReadUserBankByEmail(email string) (user_details.UserBank, error) {
	entity := user_details.UserBank{}
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
func DeleteUserBankById(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(userURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, err
	}
	return entity, nil
}
func DeleteUserBankEmail(id string) bool {
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
func ReadUserBank() ([]user_details.UserBank, error) {
	entity := []user_details.UserBank{}
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
