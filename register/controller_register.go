package register

import (
	"go.quick.start/app/controller"
)

func GetControllers() ControllerRegister {
	return ControllerRegister{
		&controller.UserController{},
		&controller.RoleController{},
	}
}
