package channel_video

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const ChannelVideoURL = api.BASE_URL + "channel/channel-video/"

func CreateChannelVideo(use domain.ChannelVideos) (domain.ChannelVideos, error) {
	entity := domain.ChannelVideos{}
	resp, _ := api.Rest().SetBody(use).Post(ChannelVideoURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateChannelVideo(role domain.ChannelVideos) (domain.ChannelVideos, error) {
	entity := domain.ChannelVideos{}
	resp, _ := api.Rest().SetBody(role).Post(ChannelVideoURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadChannelVideo(id string) (domain.ChannelVideos, error) {
	entity := domain.ChannelVideos{}
	resp, _ := api.Rest().Get(ChannelVideoURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelVideoByChannelId(id string) ([]domain.ChannelVideos, error) {
	entity := []domain.ChannelVideos{}
	resp, _ := api.Rest().Get(ChannelVideoURL + "get-by-channel/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadChannelVideoByVideoId(id string) (domain.ChannelVideos, error) {
	entity := domain.ChannelVideos{}
	resp, _ := api.Rest().Get(ChannelVideoURL + "get-by-video/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteChannelVideo(id string) (domain.ChannelVideos, error) {
	entity := domain.ChannelVideos{}
	resp, _ := api.Rest().Get(ChannelVideoURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelVideos() ([]domain.ChannelVideos, error) {
	entity := []domain.ChannelVideos{}
	resp, _ := api.Rest().Get(ChannelVideoURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
