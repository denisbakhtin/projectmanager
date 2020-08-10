package models

//CommentsDB is a comments db repository
var CommentsDB CommentsRepository

func init() {
	CommentsDB = &commentsRepository{}
}

//CommentsRepository is a repository of comments
type CommentsRepository interface {
	GetAll(userID uint64, taskID interface{}) ([]Comment, error)
	Get(userID uint64, id interface{}) (Comment, error)
	Create(userID uint64, comment Comment) (Comment, error)
	Update(userID uint64, comment Comment) (Comment, error)
	Delete(userID uint64, id interface{}) error
}

type commentsRepository struct{}

//GetAll returns all comments owned by specified user
func (cr *commentsRepository) GetAll(userID uint64, taskID interface{}) ([]Comment, error) {
	var comments []Comment
	err := db.Where("user_id = ? and task_id = ?", userID, taskID).Order("id").Find(&comments).Error
	return comments, err
}

//Get fetches a comment by its id
func (cr *commentsRepository) Get(userID uint64, id interface{}) (Comment, error) {
	comment := Comment{}
	err := db.Where("user_id = ?", userID).Preload("AttachedFiles").First(&comment, id).Error
	return comment, err
}

//Create creates a new comment in db
func (cr *commentsRepository) Create(userID uint64, comment Comment) (Comment, error) {
	comment.UserID = userID
	err := db.Create(&comment).Error
	return comment, err
}

//Update updates a comment in db
func (cr *commentsRepository) Update(userID uint64, comment Comment) (Comment, error) {
	comment.UserID = userID
	err := db.Save(&comment).Error
	return comment, err
}

//Delete removes a comment from db
func (cr *commentsRepository) Delete(userID uint64, id interface{}) error {
	cat := Comment{}
	if err := db.Where("user_id = ?", userID).First(&cat, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&cat).Error; err != nil {
		return err
	}
	return nil
}
