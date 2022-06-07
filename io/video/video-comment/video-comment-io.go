package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const videoCommentURL = api.BASE_URL + "video/video-comment/"

func CreateVideoComment(use domain.VideoComment) (domain.VideoComment, error) {
	entity := domain.VideoComment{}
	resp, _ := api.Rest().SetBody(use).Post(videoCommentURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateVideoComment(role domain.VideoComment) (domain.VideoComment, error) {
	entity := domain.VideoComment{}
	resp, _ := api.Rest().SetBody(role).Post(videoCommentURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadVideoComment(id string) (domain.VideoComment, error) {
	entity := domain.VideoComment{}
	resp, _ := api.Rest().Get(videoCommentURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteVideoComment(id string) (domain.VideoComment, error) {
	entity := domain.VideoComment{}
	resp, _ := api.Rest().Get(videoCommentURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideoComments() ([]domain.VideoComment, error) {
	var entity []domain.VideoComment
	resp, _ := api.Rest().Get(videoCommentURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
