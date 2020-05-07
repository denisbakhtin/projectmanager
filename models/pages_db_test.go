package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestPagesRepositoryGetAll(t *testing.T) {
	list, err := PagesDB.GetAll()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 0)
}

func TestPagesRepositoryGet(t *testing.T) {
	p := getOrCreatePage()
	page, err := PagesDB.Get(p.ID)
	assert.Nil(t, err)
	assert.NotEmpty(t, page)
	assert.NotZero(t, page.ID)
}

func TestPagesRepositoryCreate(t *testing.T) {
	name := fmt.Sprintf("Page-%d", time.Now().Nanosecond())
	page := Page{
		Name: name,
	}
	all, _ := PagesDB.GetAll()
	count := len(all)
	page, err := PagesDB.Create(page)
	assert.Nil(t, err)
	assert.NotZero(t, page.ID)
	assert.Equal(t, page.Name, name)

	page, err = PagesDB.Get(page.ID)
	assert.Nil(t, err)
	assert.NotZero(t, page.ID)
	assert.Equal(t, page.Name, name)

	all, _ = PagesDB.GetAll()
	assert.Equal(t, count+1, len(all), "The number of pages should have been increased by 1")
}

func TestPagesRepositoryUpdate(t *testing.T) {
	name := fmt.Sprintf("Page-%d", time.Now().Nanosecond())
	page := Page{}
	db.First(&page)
	assert.NotZero(t, page.ID)
	assert.NotEqual(t, page.Name, name)

	page.Name = name
	p, err := PagesDB.Update(page)
	assert.Nil(t, err)
	assert.Equal(t, p.Name, name)
}

func TestPagesRepositoryDelete(t *testing.T) {
	err := PagesDB.Delete(11111111111)
	assert.NotNil(t, err)
	assert.True(t, gorm.IsRecordNotFoundError(err), "This page should not exist")

	page := getOrCreatePage()
	assert.NotZero(t, page.ID)

	err = PagesDB.Delete(page.ID)
	assert.Nil(t, err)
	_, err = PagesDB.Get(page.ID)
	assert.True(t, err != nil && gorm.IsRecordNotFoundError(err), "Page should have been removed")
}
