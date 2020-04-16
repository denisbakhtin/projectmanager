package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsRepositorySpent(t *testing.T) {
	logs, err := ReportsDB.Spent(userID)
	assert.Nil(t, err)
	for _, l := range logs {
		assert.Equal(t, l.UserID, userID)
	}
}
