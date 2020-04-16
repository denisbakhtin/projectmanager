package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRepositorySearch(t *testing.T) {
	repo := SearchRepository{}
	vm, err := repo.Search(userID, "a")
	assert.Nil(t, err)
	for _, el := range vm.Categories {
		assert.Equal(t, el.UserID, userID)
	}
	for _, el := range vm.Comments {
		assert.Equal(t, el.UserID, userID)
	}
	for _, el := range vm.Projects {
		assert.Equal(t, el.UserID, userID)
	}
	for _, el := range vm.Tasks {
		assert.Equal(t, el.UserID, userID)
	}
}
