package models

//CategoriesDB is a categories db repository
var CategoriesDB CategoriesRepository

func init() {
	CategoriesDB = categoriesRepository{}
}

//CategoriesRepository is a categories repository interface
type CategoriesRepository interface {
	GetAll(userID uint64) ([]Category, error)
	Get(userID uint64, id interface{}) (Category, error)
	Create(userID uint64, category Category) (Category, error)
	Update(userID uint64, category Category) (Category, error)
	Delete(userID uint64, id interface{}) error
	Summary(userID uint64) (CategoriesSummaryVM, error)
}

//categoriesRepository is a repository of categories
type categoriesRepository struct{}

//GetAll returns all categories owned by specified user
func (cr categoriesRepository) GetAll(userID uint64) ([]Category, error) {
	var categories []Category
	err := DB.Where("user_id = ?", userID).Order("id asc").Find(&categories).Error
	return categories, err
}

//Get fetches a category by its id
func (cr categoriesRepository) Get(userID uint64, id interface{}) (Category, error) {
	category := Category{}
	err := DB.Where("user_id = ?", userID).Preload("Tasks").
		Preload("Tasks.Comments").Preload("Projects").Preload("Projects.Tasks").First(&category, id).Error
	return category, err
}

//Create creates a new category in db
func (cr categoriesRepository) Create(userID uint64, category Category) (Category, error) {
	category.UserID = userID
	err := DB.Create(&category).Error
	return category, err
}

//Update updates a category in db
func (cr categoriesRepository) Update(userID uint64, category Category) (Category, error) {
	category.UserID = userID
	err := DB.Save(&category).Error
	return category, err
}

//Delete removes a category from db
func (cr categoriesRepository) Delete(userID uint64, id interface{}) error {
	cat := Category{}
	if err := DB.Where("user_id = ?", userID).First(&cat, id).Error; err != nil {
		return err
	}
	if err := DB.Delete(&cat).Error; err != nil {
		return err
	}
	return nil
}

//Summary returns summary info for a dashboard
func (cr categoriesRepository) Summary(userID uint64) (CategoriesSummaryVM, error) {
	vm := CategoriesSummaryVM{}
	err := DB.Model(Category{}).Where("user_id = ?", userID).Count(&vm.Count).Error
	return vm, err
}
