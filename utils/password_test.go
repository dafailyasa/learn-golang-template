package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockPassword = "12345Test$"

func TestCheckPassword(t *testing.T) {
	hashed, err := HashPassword(mockPassword)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = CheckPassword(mockPassword, hashed)
	assert.NoError(t, err)
}

func TestCheckPasswordIfNotMatch(t *testing.T) {
	hashed, err := HashPassword(mockPassword)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	invalidPassword := "this-invalid-password-1234"
	err = CheckPassword(mockPassword, invalidPassword)
	assert.Error(t, err)
}
