package RestErrors

import "net/http"

const (
	badRequest = "BAD_REQUEST"
	notFound = "NOT_FOUND"
	internalServerError = "INTERNAL_SERVER_ERROR"
	validationError = "VALIDATION_ERROR"
)

type RestErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   badRequest,
		Status:  http.StatusBadRequest,
	}
}

func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   notFound,
		Status:  http.StatusNotFound,
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   internalServerError,
		Status:  http.StatusInternalServerError,
	}
}

func ValidationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   validationError,
		Status:  http.StatusBadRequest,
	}
}