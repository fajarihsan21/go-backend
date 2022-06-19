package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string      `json:"status_message"`
	Data    interface{} `json:"data"`
}

func (r *Response) Send(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func ResJSON(code int, data interface{}) *Response {
	return &Response{
		Message: getCode(code),
		Data:    data,
	}
}

func getCode(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 202:
		desc = "Accepted"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 403:
		desc = "Forbidden"
	case 404:
		desc = "Not Found"
	case 500:
		desc = "Internal Server Error"
	case 502:
		desc = "Bad Gateway"
	case 503:
		desc = "Service Unavailable"
	case 504:
		desc = "Gateway Timeout"
	}
	return desc
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

// Payload Validation
func Validate(data interface{}) error {
	v := validator.New()
	err := v.Struct(data)
	for _, e := range err.(validator.ValidationErrors) {
		return e
	}
	return nil
}
