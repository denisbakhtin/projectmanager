package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesRepositoryGetAll(t *testing.T) {
	list, err := CategoriesDB.GetAll(userID)
	assert.Nil(t, err, "Error should be nil")
	assert.GreaterOrEqual(t, len(list), 0)
	for _, cat := range list {
		assert.Equal(t, cat.UserID, userID, "Category should belong to the specified user")
	}
}

func TestCategoriesRepositoryGet(t *testing.T) {
	const id = uint64(1)
	cat, err := CategoriesDB.Get(userID, id)
	assert.True(t, err == nil || gorm.IsRecordNotFoundError(err), "Error should be nil")
	if cat.ID > 0 {
		assert.Equal(t, cat.UserID, userID, "Category should belong to the specified user")
	}
}

func TestCategoriesRepositoryCreate(t *testing.T) {
	//testing Get & GetAll here aswell
	name := fmt.Sprintf("Category-%d", time.Now().Nanosecond())
	category := Category{
		Name: name,
	}
	all, _ := CategoriesDB.GetAll(userID)
	count := len(all)
	cat, err := CategoriesDB.Create(userID, category)
	assert.Nil(t, err)
	assert.NotZero(t, cat.ID)
	assert.Equal(t, cat.Name, name)

	cat, err = CategoriesDB.Get(userID, cat.ID)
	assert.Nil(t, err)
	assert.NotZero(t, cat.ID)
	assert.Equal(t, cat.Name, name)

	all, _ = CategoriesDB.GetAll(userID)
	assert.Equal(t, count+1, len(all), "The number of categories should have been increased by 1")
}

func TestCategoriesRepositoryUpdate(t *testing.T) {
	name := fmt.Sprintf("Category-%d", time.Now().Nanosecond())
	category := Category{}
	DB.Where("user_id = ?", userID).First(&category)
	assert.NotZero(t, category.ID)
	assert.NotEqual(t, category.Name, name)

	category.Name = name
	cat, err := CategoriesDB.Update(userID, category)
	assert.Nil(t, err)
	assert.Equal(t, cat.UserID, userID)
	assert.Equal(t, cat.Name, name)
}

func TestCategoriesRepositoryDelete(t *testing.T) {
	err := CategoriesDB.Delete(userID, 11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This category should not exist")

	category, _ := getUnrelatedCategory(userID)
	assert.NotZero(t, category.ID)

	err = CategoriesDB.Delete(userID+1, category.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Should check category owner")

	err = CategoriesDB.Delete(userID, category.ID)
	assert.Nil(t, err)
	_, err = CategoriesDB.Get(userID, category.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Category should have been removed")
}

func TestCategoriesRepositorySummary(t *testing.T) {
	//ensure atleast one category exists
	getOrCreateCategory()

	vm, err := CategoriesDB.Summary(userID)
	assert.Nil(t, err)
	assert.NotEmpty(t, vm)
	assert.Greater(t, vm.Count, 0)
}
