package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const videoURL = api.BASE_URL + "video/video/"

func CreateVideo(use domain.Video) (domain.Video, error) {
	entity := domain.Video{}
	resp, _ := api.Rest().SetBody(use).Post(videoURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateVideo(role domain.Video) (domain.Video, error) {
	entity := domain.Video{}
	resp, _ := api.Rest().SetBody(role).Post(videoURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadVideo(id string) (domain.Video, error) {
	entity := domain.Video{}
	resp, _ := api.Rest().Get(videoURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteVideo(id string) (domain.Video, error) {
	entity := domain.Video{}
	resp, _ := api.Rest().Get(videoURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideos() ([]domain.Video, error) {
	entity := []domain.Video{}
	resp, _ := api.Rest().Get(videoURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
