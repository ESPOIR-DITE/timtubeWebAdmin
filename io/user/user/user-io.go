package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const userURL = api.BASE_URL + "user/user/"

func CreateUser(use domain.User) (domain.User, error) {
	entity := domain.User{}
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
func UpdateUser(user domain.User) (domain.User, error) {
	entity := domain.User{}
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
func ReadUser(id string) (domain.User, error) {
	entity := domain.User{}
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
func DeleteUser(id string) bool {
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
func ReadUsers() ([]domain.User, error) {
	entity := []domain.User{}
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
func ReadSuperAdminUsers() ([]domain.User, error) {
	entity := []domain.User{}
	resp, _ := api.Rest().Get(userURL + "getAllSuperAdmins")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAgentUsers() ([]domain.User, error) {
	entity := []domain.User{}
	resp, _ := api.Rest().Get(userURL + "getAllAgents")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAdminUsers() ([]domain.User, error) {
	entity := []domain.User{}
	resp, _ := api.Rest().Get(userURL + "getAllAdmins")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllUsers() ([]domain.User, error) {
	entity := []domain.User{}
	resp, _ := api.Rest().Get(userURL + "getAllUsers")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
