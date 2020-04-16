package mocks

import (
	"fmt"
	"strconv"

	"github.com/denisbakhtin/projectmanager/models"
)

type CommentsDBMock struct {
	Comments []models.Comment
}

//GetAll returns all comments owned by specified user
func (cm *CommentsDBMock) GetAll(userID uint64, taskID interface{}) ([]models.Comment, error) {
	return cm.Comments, nil
}

//Get fetches a comment by its id
func (cm *CommentsDBMock) Get(userID uint64, id interface{}) (models.Comment, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range cm.Comments {
		if cm.Comments[i].ID == idi {
			return cm.Comments[i], nil
		}
	}
	return models.Comment{}, fmt.Errorf("Comment not found")
}

//Create creates a new comment in db
func (cm *CommentsDBMock) Create(userID uint64, comment models.Comment) (models.Comment, error) {
	comment.UserID = userID
	cm.Comments = append(cm.Comments, comment)
	return comment, nil
}

//Update updates a comment in db
func (cm *CommentsDBMock) Update(userID uint64, comment models.Comment) (models.Comment, error) {
	for i := range cm.Comments {
		if cm.Comments[i].ID == comment.ID {
			comment.UserID = userID
			cm.Comments[i] = comment
			return comment, nil
		}
	}
	return models.Comment{}, fmt.Errorf("Comment not found")
}

//Delete removes a comment from db
func (cm *CommentsDBMock) Delete(userID uint64, id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range cm.Comments {
		if cm.Comments[i].ID == idi && cm.Comments[i].UserID == userID {
			cm.Comments[i] = cm.Comments[len(cm.Comments)-1] // Copy last element to index i.
			cm.Comments = cm.Comments[:len(cm.Comments)-1]   // Truncate slice.
			return nil
		}
	}
	return fmt.Errorf("Comment not found")
}
