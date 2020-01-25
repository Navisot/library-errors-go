package RestErrors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBadRequestError(t *testing.T) {
	err := BadRequestError("Bad Request Error Message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "Bad Request Error Message", err.Message)
	assert.EqualValues(t, "BAD_REQUEST", err.Error)
}

func TestNotFoundError(t *testing.T) {
	err := NotFoundError("Not Found Error Message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Not Found Error Message", err.Message)
	assert.EqualValues(t, "NOT_FOUND", err.Error)
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError("Internal Server Error Message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Internal Server Error Message", err.Message)
	assert.EqualValues(t, "INTERNAL_SERVER_ERROR", err.Error)
}

func TestValidationError(t *testing.T) {
	err := ValidationError("Validation Error Message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "Validation Error Message", err.Message)
	assert.EqualValues(t, "VALIDATION_ERROR", err.Error)
}