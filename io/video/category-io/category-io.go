package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const categoryURL = api.BASE_URL + "video/category/"

func CreateCategory(use domain.Category) (domain.Category, error) {
	entity := domain.Category{}
	resp, _ := api.Rest().SetBody(use).Post(categoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCategory(role domain.Category) (domain.Category, error) {
	entity := domain.Category{}
	resp, _ := api.Rest().SetBody(role).Post(categoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadCategory(id string) (domain.Category, error) {
	entity := domain.Category{}
	resp, _ := api.Rest().Get(categoryURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteCategory(id string) (domain.Category, error) {
	entity := domain.Category{}
	resp, _ := api.Rest().Get(categoryURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCategorys() ([]domain.Category, error) {
	entity := []domain.Category{}
	resp, _ := api.Rest().Get(categoryURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
