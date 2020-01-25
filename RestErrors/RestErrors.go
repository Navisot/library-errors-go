package RestErrors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "BAD_REQUEST",
		Status:  http.StatusBadRequest,
	}
}

func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "NOT_FOUND",
		Status:  http.StatusNotFound,
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "INTERNAL_SERVER_ERROR",
		Status:  http.StatusInternalServerError,
	}
}

func ValidationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "VALIDATION_ERROR",
		Status:  http.StatusBadRequest,
	}
}