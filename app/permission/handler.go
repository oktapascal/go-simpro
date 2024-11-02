package permission

import (
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/model"
)

type Handler struct {
	svc      model.PermissionService
	validate *validator.Validate
}
