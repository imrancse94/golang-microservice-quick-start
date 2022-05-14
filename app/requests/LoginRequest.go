package requests

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go.quick.start/apiresponse"
	"go.quick.start/models"
	"go.quick.start/validation"
	"net/http"
)

// Secret key to uniquely sign the token
var key []byte

// Credential User's login information
type Credential struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// Token jwt Standard Claim Object
type Token struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func RequestValidate(writer http.ResponseWriter, request *http.Request, input models.RegisterUserInput) (bool, error) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)

	//Could not decode json
	if err != nil {
		errors := make(map[string]string)
		apiresponse.ErrorResponse(http.StatusUnprocessableEntity, "E101", errors, "Invalid JSON", writer)
		return false, err
	}

	if ok, errors := validation.ValidateInputs(input); !ok {
		apiresponse.ErrorResponse(http.StatusUnprocessableEntity, "E102", errors, "Validation Error", writer)
		return false, err
	}

	return true, nil
}

func LoginValidate(writer http.ResponseWriter, request *http.Request) (bool bool, error error, data Credential) {
	var input Credential
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)

	if ok, errors := validation.ValidateInputs(input); !ok {
		apiresponse.ErrorResponse(http.StatusUnprocessableEntity, "E102", errors, "Validation Error", writer)
		return false, err, input
	}

	return true, nil, input
}
