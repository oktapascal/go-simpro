package model

import "net/http"

type (
	Navigation struct {
		Id       string            `json:"id"`
		Name     string            `json:"name"`
		Icon     string            `json:"icon"`
		PathUrl  string            `json:"path_url"`
		Children []NavigationChild `json:"children"`
	}

	NavigationChild struct {
		Id       string            `json:"id"`
		Name     string            `json:"name"`
		PathUrl  string            `json:"path_url"`
		Children []NavigationChild `json:"children"`
	}

	NavigationRepository interface {
		GetNavigation(group string) *[]Navigation
	}

	NavigationService interface {
		GetNavigation(group string) *[]Navigation
	}

	NavigationHandler interface {
		GetNavigation() http.HandlerFunc
	}
)
