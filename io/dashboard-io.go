package io

import (
	"errors"
	"timtubeWebAdmin/api"
	"timtubeWebAdmin/domain"
)

func GetDashboardData() (domain.Dashboard, error) {
	entity := domain.Dashboard{}
	resp, _ := api.Rest().Get(api.BASE_URL + "board")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
