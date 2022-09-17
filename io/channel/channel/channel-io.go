package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const ChannelURL = api.BASE_URL + "channel/channel/"

func CreateChannel(use domain.Channel) (domain.Channel, error) {
	entity := domain.Channel{}
	resp, _ := api.Rest().SetBody(use).Post(ChannelURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateChannel(role domain.Channel) (domain.Channel, error) {
	entity := domain.Channel{}
	resp, _ := api.Rest().SetBody(role).Post(ChannelURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadChannel(id string) (domain.Channel, error) {
	entity := domain.Channel{}
	resp, _ := api.Rest().Get(ChannelURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelByUserId(id string) ([]domain.Channel, error) {
	entity := []domain.Channel{}
	resp, _ := api.Rest().Get(ChannelURL + "get-by-user/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadChannelByRegion(id string) ([]domain.Channel, error) {
	entity := []domain.Channel{}
	resp, _ := api.Rest().Get(ChannelURL + "get-by-region/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteChannel(id string) (domain.Channel, error) {
	entity := domain.Channel{}
	resp, _ := api.Rest().Get(ChannelURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannels() ([]domain.Channel, error) {
	entity := []domain.Channel{}
	resp, _ := api.Rest().Get(ChannelURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
