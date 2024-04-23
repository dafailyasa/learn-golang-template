package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockInvalidCurrency string = "JPN"
var mockError string = fmt.Sprintf("currently currency just support (%s,%s,%s)", USD, EUR, IDR)

func TestReturnNoErrorIfSupportCurrency(t *testing.T) {
	err := IsSupportCurrency(IDR)

	assert.NoError(t, err)
}

func TestReturnErrorIfNotSupportCurrency(t *testing.T) {
	err := IsSupportCurrency(mockInvalidCurrency)

	assert.NotEmpty(t, err)
	assert.Error(t, err)
	assert.EqualError(t, err, mockError)
}
