package response

import (
	"errors"
	"fmt"
	"net/http"
)

type body map[string]interface{}

type SuccessResponse struct {
	Message string
	Data    interface{}
}

// Map for errors with http code
var ResponseCode = make(map[string]int, 0)

var (

	// common
	ErrInvalidRequestParams  = NewError("request params is not valid", http.StatusBadRequest)
	ErrDecodeRequestBody     = NewError("failed to decode request body", http.StatusInternalServerError)
	ErrMissingParams         = NewError("params not found in the request", http.StatusBadRequest)
	ErrParsingRequestParams  = NewError("failed to parse request params", http.StatusBadRequest)
	ErrParsingRequestBody    = NewError("failed to parse request body", http.StatusBadRequest)
	ErrParsingRequestHeaders = NewError("failed to parse request headers", http.StatusBadRequest)
	AccessForbidden          = NewError("access forbidden", http.StatusForbidden)
	ErrParsingEndpoint       = NewError("failed to parse endpoint", http.StatusInternalServerError)
	ErrJSONMarshal           = NewError("failed to marshal data", http.StatusInternalServerError)
	ErrJSONUnmarshal         = NewError("failed to unmarshal data", http.StatusInternalServerError)
)

func responseMap() map[string]int {
	return ResponseCode
}

func NewError(message string, httpCode int) error {
	_, available := ResponseCode[message]
	if !available {
		ResponseCode[message] = httpCode
	}
	return errors.New(message)
}

func ValidationErrors(err error, entity string) (int, body) {
	message := fmt.Sprintf("failed to validate the fields of the %v", entity)
	return validationResponse(message, err)
}

func GenerateErrorResponseBody(err error) (int, body) {
	message := err.Error()
	return readFromMap(message)
}

func readFromMap(message string) (int, body) {
	httpStatus, available := responseMap()[message]
	if available {
		return httpStatus, generateResponseBody(message)
	}
	return http.StatusInternalServerError, generateResponseBody(message)
}

func generateResponseBody(message string) body {
	return body{
		"message": message,
	}
}

func validationResponse(message string, err error) (int, body) {
	return http.StatusBadRequest, body{
		"message":          message,
		"validation_error": err,
	}
}

func GenerateSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Data:    data,
	}

}
