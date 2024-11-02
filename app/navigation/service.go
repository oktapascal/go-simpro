package navigation

import "github.com/oktapascal/go-simpro/model"

type Service struct {
	rpo model.NavigationRepository
}

func (svc *Service) GetNavigation(group string) *[]model.Navigation {
	return svc.rpo.GetNavigation(group)
}
