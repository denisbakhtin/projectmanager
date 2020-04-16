package mocks

import (
	"fmt"
	"strconv"

	"github.com/denisbakhtin/projectmanager/models"
)

type PagesDBMock struct {
	Pages []models.Page
}

//GetAll returns all pages owned by specified user
func (pm *PagesDBMock) GetAll() ([]models.Page, error) {
	return pm.Pages, nil
}

//Get fetches a page by its id
func (pm *PagesDBMock) Get(id interface{}) (models.Page, error) {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range pm.Pages {
		if pm.Pages[i].ID == idi {
			return pm.Pages[i], nil
		}
	}
	return models.Page{}, fmt.Errorf("Page not found")
}

//GetPagesForMenu returns a list of pages for navbar menu
func (pm *PagesDBMock) GetPagesForMenu() ([]models.Page, error) {
	return pm.Pages, nil
}

//Create pmeates a new page in db
func (pm *PagesDBMock) Create(page models.Page) (models.Page, error) {
	pm.Pages = append(pm.Pages, page)
	return page, nil
}

//Update updates a page in db
func (pm *PagesDBMock) Update(page models.Page) (models.Page, error) {
	for i := range pm.Pages {
		if pm.Pages[i].ID == page.ID {
			pm.Pages[i] = page
			return pm.Pages[i], nil
		}
	}
	return models.Page{}, fmt.Errorf("Page not found")
}

//Delete removes a page from db
func (pm *PagesDBMock) Delete(id interface{}) error {
	idi, _ := strconv.ParseUint(id.(string), 10, 64)
	for i := range pm.Pages {
		if pm.Pages[i].ID == idi {
			pm.Pages[i] = pm.Pages[len(pm.Pages)-1] // Copy last element to index i.
			pm.Pages = pm.Pages[:len(pm.Pages)-1]   // Truncate slice.
			return nil
		}
	}
	return fmt.Errorf("Page not found")
}
