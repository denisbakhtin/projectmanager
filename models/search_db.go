package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

//SearchDB is a db search repository
var SearchDB SearchRepository

func init() {
	SearchDB = SearchRepository{}
}

//SearchRepository is a repository for searching entities
type SearchRepository struct{}

//Search returns all matching entities owned by the specified user
func (pr *SearchRepository) Search(userID uint64, query string) (SearchVM, error) {
	vm := SearchVM{}

	if len(strings.TrimSpace(query)) == 0 {
		return SearchVM{}, fmt.Errorf("Query string is empty")
	}
	err := db.Where("user_id = ?", userID).
		Where("to_tsvector(name) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Order("id asc").Find(&vm.Categories).Error
	if err != nil {
		return SearchVM{}, err
	}

	err = db.Preload("Tasks").Preload("Category").Where("user_id = ?", userID).
		Where("to_tsvector(name || ' ' || description) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Order("archived asc, created_at asc").Find(&vm.Projects).Error
	if err != nil {
		return SearchVM{}, err
	}

	q := db.Where("user_id = ?", userID).
		Where("to_tsvector(name || ' ' || description) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Preload("Project").Preload("Comments").Preload("Category")
	q = q.Preload("TaskLogs", func(db *gorm.DB) *gorm.DB {
		return db.Where("session_id = 0 and minutes > 0")
	})
	err = q.Order("tasks.completed asc, tasks.created_at asc").Find(&vm.Tasks).Error
	if err != nil {
		return SearchVM{}, err
	}

	err = db.Where("user_id = ?", userID).
		Where("to_tsvector(contents) @@ to_tsquery(?)", fmt.Sprintf("%s:*", query)).
		Order("id").Find(&vm.Comments).Error

	return vm, err
}
