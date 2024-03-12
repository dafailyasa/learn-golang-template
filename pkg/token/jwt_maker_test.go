package token

import (
	"testing"
	"time"

	"github.com/dafailyasa/learn-golang-template/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockInvalidSecretKey = "less than 32 characters"
const mockEmail = "mock@mail.com"
const expectedError = "invalid key size: must be at least 32 characters"

func TestNewJwtMaker(t *testing.T) {
	maker, err := NewJwtMaker(tools.RandomString(32))

	require.NoError(t, err)
	require.NotNil(t, maker)
}

func TestNewJwtMakerIfSecretKeyInvalidCount(t *testing.T) {
	maker, err := NewJwtMaker(mockInvalidSecretKey)

	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err.Error())
	}

	require.Nil(t, maker)
}

func TestShouldBeGenerateAccessTokenAndVerifyToken(t *testing.T) {
	maker, err := NewJwtMaker(tools.RandomString(32))

	require.NoError(t, err)
	require.NotNil(t, maker)

	duration := time.Minute
	expiredAt := time.Now().Add(duration)
	token, payload, err := maker.CreateToken(mockEmail, duration)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotNil(t, payload)

	assert.Equal(t, payload.Email, mockEmail)
	require.WithinDuration(t, payload.ExpiresAt.Time, expiredAt, time.Second)
	require.WithinDuration(t, payload.IssuedAt.Time, time.Now(), time.Second)
	require.WithinDuration(t, payload.NotBefore.Time, time.Now(), time.Second)

	payload, err = maker.VerifyToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, payload)

	assert.Equal(t, payload.Email, mockEmail)
	require.WithinDuration(t, payload.ExpiresAt.Time, expiredAt, time.Second)
	require.WithinDuration(t, payload.IssuedAt.Time, time.Now(), time.Second)
	require.WithinDuration(t, payload.NotBefore.Time, time.Now(), time.Second)
}

func TestExpiredToken(t *testing.T) {
	maker, err := NewJwtMaker(tools.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker)

	token, payload, err := maker.CreateToken(mockEmail, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Nil(t, payload)
	require.Error(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "token has invalid claims: token is expired")
	}
}
