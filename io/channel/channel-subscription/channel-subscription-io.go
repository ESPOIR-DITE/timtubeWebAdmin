package user

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

const ChannelSubscriptionURL = api.BASE_URL + "channel/channel-subscription/"

func CreateChannelSubscription(use domain.ChannelSubscription) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().SetBody(use).Post(ChannelSubscriptionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateChannelSubscription(role domain.ChannelSubscription) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().SetBody(role).Post(ChannelSubscriptionURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func CountChannelSubscriptionByChannelId(id string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get-by-channel/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelSubscription(id string) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelSubscriptionByUserId(id string) ([]domain.ChannelSubscription, error) {
	entity := []domain.ChannelSubscription{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get-by-user/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadChannelSubscriptionByChannel(id string) ([]domain.ChannelSubscription, error) {
	entity := []domain.ChannelSubscription{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get-by-channel/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteChannel(id string) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelSubscriptions() ([]domain.ChannelSubscription, error) {
	entity := []domain.ChannelSubscription{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
