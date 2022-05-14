package validation

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)

var validate *validator.Validate

func ValidateInputs(dataSet interface{}) (bool, map[string]string) {
	//Make it global for caching
	validate = validator.New()
	err := validate.Struct(dataSet)

	if err != nil {
		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		//Validation errors occurred
		errors := make(map[string]string)
		//Use reflector to reverse engineer struct
		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = "This " + name + " is required"
				break
			case "email":
				errors[name] = "The " + name + " should be a valid email"
				break
			case "eqfield":
				errors[name] = "The " + name + " should be equal to the " + err.Param()
				break
			default:
				errors[name] = "The " + name + " is invalid"
				break
			}
		}

		return false, errors
	}
	return true, nil
}
