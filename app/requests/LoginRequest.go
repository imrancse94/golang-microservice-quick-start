package requests

import (
	"encoding/json"
	"net/http"
)

// Secret key to uniquely sign the token
var key []byte

// Credential User's login information
type Credential struct {
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email"`
	Password string `json:"password" valid:"required~Password is required"`
}

/*func RequestValidate(writer http.ResponseWriter, request *http.Request, input models.RegisterUserInput) (bool, error) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)

	//Could not decode json
	if err != nil {
		errors := make(map[string]string)
		api.ErrorResponse(http.StatusUnprocessableEntity, "E101", errors, "Invalid JSON", writer)
		return false, err
	}

	if ok, errors := validation.ValidateInputs(input); !ok {
		api.ErrorResponse(http.StatusUnprocessableEntity, "E102", errors, "Validation Error", writer)
		return false, err
	}

	return true, nil
}
*/

func LoginRequestToStruct(r *http.Request) Credential {
	var input Credential
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&input)
	return input
}
