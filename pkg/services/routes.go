package services

import "auth-service/pkg/models"

func NewRoute(path string, method string, name string, params string) models.Route {
	return models.Route{
		Method: method,
		Path:   path,
		Name:   name,
		Params: params,
	}
}
