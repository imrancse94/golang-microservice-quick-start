package controller

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.quick.start/Services"
	"go.quick.start/api"
	"go.quick.start/app/requests"
	"go.quick.start/application"
	"go.quick.start/models"
	"net/http"
	"strconv"
)

var DB *gorm.DB

type UserController struct {
	application.BaseController
}

type AuthResponse struct {
	Token   string      `json:"token"`
	Expires string      `json:"expires"`
	User    interface{} `json:"user"`
}

func (c *UserController) RegisterUser(w http.ResponseWriter, request *http.Request) {
	responseData := api.Response{
		Status:  api.Status("SUCCESS"),
		Message: "Dummy data",
		Data:    "",
	}
	api.SuccessRespond(responseData, w)
}

/*func (c *UserController) Index(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})
	data := make(map[string]interface{})
	data["name"] = "test"
	data["email"] = "abc@gmail.com"
	response["data"] = data
	response["status"] = true
	api.SuccessRespond(response, writer)
}*/

func test() {
	fmt.Println("This is for test")
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	input := requests.LoginRequestToStruct(r)
	userData, message := Services.Login(input)

	if userData != nil && message == "" {
		//expireTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
		/*data := AuthResponse{
			Expires: strconv.Itoa(expireTime * 60),
			User:    userData,
		}*/

		responseData := api.Response{
			Status:  api.Status("SUCCESS"),
			Message: "Token generated successfully",
			Data:    userData,
		}
		api.SuccessRespond(responseData, w)
		return
	}

	responseData := api.Response{
		Status:  api.Status("FAILED"),
		Message: "Something went wrong",
		Data:    "",
	}
	api.ErrorResponse(responseData, w)
}

func (c *UserController) AuthData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header.Get("auth_id"))
	//user := models.GetUserById(id)
	user := models.GetRolePageByUserId(id)
	responseData := api.Response{
		Status:  api.Status("SUCCESS"),
		Message: user.Message,
		Data:    user.Data,
	}
	api.SuccessRespond(responseData, w)
}
