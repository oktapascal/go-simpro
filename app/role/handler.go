package role

import (
	"github.com/oktapascal/go-simpro/model"
)

type Handler struct {
	svc      model.RoleService
	validate *validator.Validate
}
