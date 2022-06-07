package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const roleURL = api.BASE_URL + "user/role/"

func CreateRole(use domain.Role) (domain.Role, error) {
	entity := domain.Role{}
	resp, _ := api.Rest().SetBody(use).Post(roleURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateRole(role domain.Role) (domain.Role, error) {
	entity := domain.Role{}
	resp, _ := api.Rest().SetBody(role).Post(roleURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadRole(id string) (domain.Role, error) {
	entity := domain.Role{}
	resp, _ := api.Rest().Get(roleURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteRole(id string) (domain.Role, error) {
	entity := domain.Role{}
	resp, _ := api.Rest().Get(roleURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadRoles() ([]domain.Role, error) {
	entity := []domain.Role{}
	resp, _ := api.Rest().Get(roleURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
