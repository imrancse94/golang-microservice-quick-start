package api

func Status(key string) string {
	var status = make(map[string]string)

	status["SUCCESS"] = "200"
	status["ERROR"] = "502"
	status["FAILED"] = "400"
	status["NOTFOUND"] = "404"
	status["UNAUTHORIZED"] = "401"
	status["FORBIDDEN"] = "403"
	status["INTERNAL_SERVER_ERROR"] = "500"
	status["VALIDATION_ERROR"] = "422"

	return status[key]
}

func StatusMessage(key string) string {
	var status = make(map[string]string)

	status["VALIDATION_ERROR"] = "Input Parameters are not valid"

	return status[key]
}
