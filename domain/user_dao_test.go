package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotUserFound(t *testing.T) {
	//Initialisatin

	//Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user, "We are not expecting with user id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Sharan", user.FirstName)
	assert.EqualValues(t, "Badni", user.LastName)
	assert.EqualValues(t, "sharan.badni@gmail.com", user.Email)
}
