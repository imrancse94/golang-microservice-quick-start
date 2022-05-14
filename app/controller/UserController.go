package controller

import (
	"fmt"
	"go.quick.start/Services"
	"go.quick.start/apiresponse"
	"go.quick.start/app/requests"
	"go.quick.start/application"
	"go.quick.start/models"
	"net/http"
	"os"
	"strconv"
)

type UserController struct {
	application.BaseController
}

func (c *UserController) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	var input models.RegisterUserInput
	requests.RequestValidate(writer, request, input)

	response := make(map[string]interface{})
	//apiresponse["message"] = "Email is " + input["email"].(string)
	apiresponse.SuccessRespond(response, writer)

	//TODO validate json and use successresponse
	//writer.Write([]byte("Email is " + input["email"].(string)))
}

func (c *UserController) Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Index")
	response := make(map[string]interface{})
	data := make(map[string]interface{})
	data["name"] = "test"
	data["email"] = "abc@gmail.com"
	response["data"] = data
	response["status"] = true
	apiresponse.SuccessRespond(response, writer)
}

func test() {
	fmt.Println("This is for test")
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	isValidate, err, input := requests.LoginValidate(w, r)
	fmt.Println("Reflected: " + input.Email)
	if isValidate && err == nil {
		token, userData, message := Services.Login(input)
		errdata := make(map[string]string)

		if token != "" && message == "" {
			expireTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
			data := map[string]interface{}{
				"token":  token,
				"expire": strconv.Itoa(expireTime * 60),
				"data":   userData,
				"status": true,
			}

			apiresponse.SuccessRespond(data, w)
		} else {
			apiresponse.ErrorResponse(http.StatusNotFound, "E105", errdata, message, w)
		}
	}
	fmt.Println(isValidate, err)

}
