package user

import (
	"errors"
	"fmt"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const userAccountURL = api.BASE_URL + "user/user-account/"

func LoginUserAccount(use domain.UserAccount) (domain.UserAccount, error) {
	fmt.Println(userAccountURL)
	entity := domain.UserAccount{}
	resp, _ := api.Rest().SetBody(use).Post(userAccountURL + "login")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserAccount(use domain.UserAccount) (domain.UserAccount, error) {
	entity := domain.UserAccount{}
	resp, _ := api.Rest().SetBody(use).Post(userAccountURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserAccount(role domain.UserAccount) (domain.UserAccount, error) {
	entity := domain.UserAccount{}
	resp, _ := api.Rest().SetBody(role).Post(userAccountURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadUserAccount(id string) (domain.UserAccount, error) {
	entity := domain.UserAccount{}
	resp, _ := api.Rest().Get(userAccountURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserAccountWithEmail(email string) (domain.UserAccount, error) {
	entity := domain.UserAccount{}
	resp, _ := api.Rest().Get(userAccountURL + "getWithEmail/" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserAccount(id string) (domain.UserAccount, error) {
	entity := domain.UserAccount{}
	resp, _ := api.Rest().Get(userAccountURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserAccounts() ([]domain.UserAccount, error) {
	entity := []domain.UserAccount{}
	resp, _ := api.Rest().Get(userAccountURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
