package channel_type

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const ChannelTypeURL = api.BASE_URL + "channel/channel-type/"

func CreateChannelType(use domain.ChannelType) (domain.ChannelType, error) {
	entity := domain.ChannelType{}
	resp, _ := api.Rest().SetBody(use).Post(ChannelTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateChannelType(role domain.ChannelType) (domain.ChannelType, error) {
	entity := domain.ChannelType{}
	resp, _ := api.Rest().SetBody(role).Post(ChannelTypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadChannelType(id string) (domain.ChannelType, error) {
	entity := domain.ChannelType{}
	resp, _ := api.Rest().Get(ChannelTypeURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteChannel(id string) (domain.ChannelType, error) {
	entity := domain.ChannelType{}
	resp, _ := api.Rest().Get(ChannelTypeURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadChannelTypes() ([]domain.ChannelType, error) {
	entity := []domain.ChannelType{}
	resp, _ := api.Rest().Get(ChannelTypeURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
