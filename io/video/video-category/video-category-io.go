package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const videoCategoryURL = api.BASE_URL + "video/video-category/"

func CreateVideoCategory(use domain.VideoCategory) (domain.VideoCategory, error) {
	entity := domain.VideoCategory{}
	resp, _ := api.Rest().SetBody(use).Post(videoCategoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateVideoCategory(role domain.VideoCategory) (domain.VideoCategory, error) {
	entity := domain.VideoCategory{}
	resp, _ := api.Rest().SetBody(role).Post(videoCategoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadVideoCategory(id string) (domain.VideoCategory, error) {
	entity := domain.VideoCategory{}
	resp, _ := api.Rest().Get(videoCategoryURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteVideoCategory(id string) (domain.VideoCategory, error) {
	entity := domain.VideoCategory{}
	resp, _ := api.Rest().Get(videoCategoryURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideoCategories() ([]domain.VideoCategory, error) {
	var entity []domain.VideoCategory
	resp, _ := api.Rest().Get(videoCategoryURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
