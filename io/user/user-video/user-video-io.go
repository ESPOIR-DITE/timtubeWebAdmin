package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const userVideoURL = api.BASE_URL + "user/user-video/"

func CreateUserVideo(use domain.UserVideo) (domain.UserVideo, error) {
	entity := domain.UserVideo{}
	resp, _ := api.Rest().SetBody(use).Post(userVideoURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserVideo(role domain.UserVideo) (domain.UserVideo, error) {
	entity := domain.UserVideo{}
	resp, _ := api.Rest().SetBody(role).Post(userVideoURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadAllOfUserVideo(id string) ([]domain.UserVideo, error) {
	entity := []domain.UserVideo{}
	resp, _ := api.Rest().Get(userVideoURL + "get-all/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserVideo(id string) (domain.UserVideo, error) {
	entity := domain.UserVideo{}
	resp, _ := api.Rest().Get(userVideoURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserVideo(id string) (domain.UserVideo, error) {
	entity := domain.UserVideo{}
	resp, _ := api.Rest().Get(userVideoURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserVideos() ([]domain.UserVideo, error) {
	entity := []domain.UserVideo{}
	resp, _ := api.Rest().Get(userVideoURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
