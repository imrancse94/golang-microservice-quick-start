package controller

import (
	"go.quick.start/Helper"
	"go.quick.start/api"
	"go.quick.start/application"
	"go.quick.start/models"
	"net/http"
)

type RoleController struct {
	application.BaseController
}

func (c *RoleController) GetRole(w http.ResponseWriter, r *http.Request) {
	responseData := api.Response{
		Status:  "200",
		Message: "Roles retrieved successfully",
		Data:    Helper.ToSliceOfAny(models.GetRoles().Data.([]models.Role)),
	}
	api.SuccessRespond(responseData, w)

}
