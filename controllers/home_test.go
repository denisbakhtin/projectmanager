package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeGet(t *testing.T) {
	resp, err := http.Get(server.URL)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
