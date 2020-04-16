package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	resp, err := jsonGet(server.URL + "/api/settings")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
