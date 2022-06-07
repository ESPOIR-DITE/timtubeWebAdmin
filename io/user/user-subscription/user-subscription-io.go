package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const roleURL = api.BASE_URL + "user/user-subscription/"

func CreateUserSubscription(use domain.UserSubscription) (domain.UserSubscription, error) {
	entity := domain.UserSubscription{}
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
func UpdateUserSubscription(role domain.UserSubscription) (domain.UserSubscription, error) {
	entity := domain.UserSubscription{}
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
func ReadUserSubscription(id string) (domain.UserSubscription, error) {
	entity := domain.UserSubscription{}
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

func DeleteUserSubscription(id string) (domain.UserSubscription, error) {
	entity := domain.UserSubscription{}
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
func ReadUserSubscriptions() ([]domain.UserSubscription, error) {
	entity := []domain.UserSubscription{}
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
