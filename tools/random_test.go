package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	count := 33
	data := RandomString(count)

	assert.NotNil(t, data)
	assert.Equal(t, len(data), count)
}
