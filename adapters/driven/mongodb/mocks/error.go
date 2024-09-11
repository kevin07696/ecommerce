package mocks

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ mongo.ServerError = MockServerError{}

type MockServerError struct {
	Code    int
	Name    string
	Message string
	mongo.ServerError
}

func (ms MockServerError) Error() string {
	if ms.Name != "" {
		return fmt.Sprintf("(%v) %v", ms.Name, ms.Message)
	}
	return ms.Message
}

func (ms MockServerError) HasErrorCode(code int) bool {
	return ms.Code == code
}

func (ms MockServerError) HasErrorMessage(substr string) bool {
	return strings.Contains(ms.Message, substr)
}

func (ms MockServerError) HasErrorCodeWithMessage(code int, substr string) bool {
	return ms.HasErrorCode(code) && ms.HasErrorMessage(substr)
}
