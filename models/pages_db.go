package models

//PagesDB is a pages db repository
var PagesDB PagesRepository

func init() {
	PagesDB = &pagesRepository{}
}

//PagesRepository is a repository of pages
type PagesRepository interface {
	GetAll() ([]Page, error)
	Get(id interface{}) (Page, error)
	GetPagesForMenu() ([]Page, error)
	Create(page Page) (Page, error)
	Update(page Page) (Page, error)
	Delete(id interface{}) error
}
type pagesRepository struct{}

//GetAll returns all pages owned by specified user
func (cr *pagesRepository) GetAll() ([]Page, error) {
	var pages []Page
	err := DB.Order("published desc, id asc").Find(&pages).Error
	return pages, err
}

//Get fetches a page by its id
func (cr *pagesRepository) Get(id interface{}) (Page, error) {
	page := Page{}
	err := DB.First(&page, id).Error
	return page, err
}

//GetPagesForMenu returns a list of pages for navbar menu
func (cr *pagesRepository) GetPagesForMenu() ([]Page, error) {
	var pages []Page
	err := DB.Where("published = true").Order("id asc").Select("id, name").Find(&pages).Error
	return pages, err
}

//Create creates a new page in db
func (cr *pagesRepository) Create(page Page) (Page, error) {
	err := DB.Create(&page).Error
	return page, err
}

//Update updates a page in db
func (cr *pagesRepository) Update(page Page) (Page, error) {
	err := DB.Save(&page).Error
	return page, err
}

//Delete removes a page from db
func (cr *pagesRepository) Delete(id interface{}) error {
	page := Page{}
	if err := DB.First(&page, id).Error; err != nil {
		return err
	}
	return DB.Delete(&page).Error
}
