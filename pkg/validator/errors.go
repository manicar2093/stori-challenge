package validator

import "net/http"

type (
	ValidationError struct {
		Errors interface{} `json:"errors,inline"`
	}
)

func (c *ValidationError) Error() string {
	return "request is not valid. Some field does not fulfill requirements"
}

func (c *ValidationError) StatusCode() int {
	return http.StatusBadRequest
}
