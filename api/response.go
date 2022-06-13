package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessRespond(fields Response, writer http.ResponseWriter) {
	message, err := json.MarshalIndent(fields, "", " ")
	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occurred internally"))
		return
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func ErrorResponse(fields Response, writer http.ResponseWriter) {
	//Create a new map and fill it
	data, err := json.MarshalIndent(fields, "", " ")

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
	return
}
