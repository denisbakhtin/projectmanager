package mocks

import (
	"fmt"
	"strconv"

	"github.com/denisbakhtin/projectmanager/models"
)

//CategoriesDBMock is a CategoriesDB repository mock
type CategoriesDBMock struct {
	Categories []models.Category
}

func (cr *CategoriesDBMock) GetAll(userID uint64) ([]models.Category, error) {
	return cr.Categories, nil
}

func (cr *CategoriesDBMock) Get(userID uint64, id interface{}) (models.Category, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for _, c := range cr.Categories {
		if c.ID == idi {
			return c, nil
		}
	}
	return models.Category{}, fmt.Errorf("Category not found")
}

//Create creates a new category in db
func (cr *CategoriesDBMock) Create(userID uint64, category models.Category) (models.Category, error) {
	category.UserID = userID
	cr.Categories = append(cr.Categories, category)
	return category, nil
}

//Update updates a category in db
func (cr *CategoriesDBMock) Update(userID uint64, category models.Category) (models.Category, error) {
	category.UserID = userID
	for i := range cr.Categories {
		if cr.Categories[i].ID == category.ID {
			cr.Categories[i] = category
			return category, nil
		}
	}
	return models.Category{}, fmt.Errorf("Category not found")
}

//Delete removes a category from db
func (cr *CategoriesDBMock) Delete(userID uint64, id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range cr.Categories {
		if cr.Categories[i].ID == idi && cr.Categories[i].UserID == userID {
			cr.Categories[i] = cr.Categories[len(cr.Categories)-1] // Copy last element to index i.
			cr.Categories = cr.Categories[:len(cr.Categories)-1]   // Truncate slice.
			return nil
		}
	}
	return nil
}

func (cr *CategoriesDBMock) Summary(userID uint64) (models.CategoriesSummaryVM, error) {
	return models.CategoriesSummaryVM{Count: len(cr.Categories)}, nil
}
