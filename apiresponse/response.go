package apiresponse

import (
	"encoding/json"
	"net/http"
)

func SuccessRespond(fields map[string]interface{}, writer http.ResponseWriter) {
	fields["status"] = "success"
	message, err := json.Marshal(fields)
	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func ErrorResponse(statusCode int, errorCode string, errors map[string]string, message string, writer http.ResponseWriter) {
	//Create a new map and fill it
	fields := make(map[string]interface{})
	fields["status"] = false
	fields["error_code"] = errorCode
	fields["message"] = message
	fields["errors"] = errors

	data, err := json.Marshal(fields)

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(data)
	return
}
