package controller

import (
	"fmt"
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
	//var interfaceSlice []interface{} = make([]interface{}, len(models.GetRoles().Data.(models.Role)))
	//rolesData := make(map[string]interface{})
	rolesData := map[string]interface{}{
		"data": Helper.ToSliceOfAny(models.GetRoles().Data.([]models.Role)),
	}
	//models.GetRoles().Data.(models.Role)
	//b := make([]interface{}, len(models.GetRoles().Data.([]models.Role)))
	//rolesData := make(map[string]interface{}, len(models.GetRoles().Data.([]models.Role)))
	fmt.Println("ssss", rolesData)
	/*for i, d := range rolesData {
		fmt.Println("slice", i, d.(models.Role).Title)
	}*/
	responseData := api.Response{
		Status:  "200",
		Message: "Roles retrieved successfully",
		Data:    rolesData,
	}
	api.SuccessRespond(responseData, w)

}
