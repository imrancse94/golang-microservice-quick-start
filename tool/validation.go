package tool

import (
	"gopkg.in/asaskevich/govalidator.v4"
)

// ValidateRequest incoming request
func ValidateRequest(data interface{}) interface{} {
	if valid, err := govalidator.ValidateStruct(data); valid == false {
		return govalidator.ErrorsByField(err)
	}

	return nil
}
