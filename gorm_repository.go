package gobe

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Initiate a GORM repository
type GormRepository struct {
	Db *gorm.DB
}

// Create/insert a new record to the table by defining the model explicitly
func (g *GormRepository) Create(model interface{}) error {
	return g.Db.Create(model).Error
}

// Update a value in a record.
//
//	Example:
//	UpdateBy(User, map[string]interface{}{"id":1}, map[string]interface{}{"name":"YYY"}) // will update a User record name with ID = 1 to "YYY"
func (g *GormRepository) UpdateBy(model interface{}, by map[string]interface{}, value map[string]interface{}) error {
	return g.Db.Model(model).Where(by).Updates(value).Error
}

// Delete a record.
//
//	Example:
//	DeleteBy(User, map[string]interface{}{"id":1}) // will delete a User record name with ID = 1
func (g *GormRepository) DeleteBy(model interface{}, by map[string]interface{}) error {
	return g.Db.Model(model).Delete(by).Error
}

// Find a record by using the row name and the data.
//
//	Example:
//	FindBy(&User, map[string]interface{}{"id":1}) // will get result User with ID = 1
func (g *GormRepository) FindBy(model interface{}, by map[string]interface{}) (interface{}, error) {
	err := g.Db.Where(by).First(&model).Error
	return model, err
}

// Find a record by using the row name and the data. Preload will get all associations in the model
//
//	Example:
//	FindByWithPreload(&User, map[string]interface{}{"id":1}) // will get result User with ID = 1
func (g *GormRepository) FindByWithPreload(model interface{}, by map[string]interface{}) (interface{}, error) {
	err := g.Db.Preload(clause.Associations).Where(by).First(&model).Error
	return model, err
}

// Find a record by using the row name and the data. Preload will get all associations in the model
//
//	Example:
//	FindByWithNestedPreload(&User, map[string]interface{}{"id":1}, "User.Role") // will get result User with ID = 1
func (g *GormRepository) FindByWithNestedPreload(model interface{}, by map[string]interface{}, nestedPreload string) (interface{}, error) {
	err := g.Db.Preload(nestedPreload).Preload(clause.Associations).Where(by).First(&model).Error
	return model, err
}

// Find any records by using the row name and the data.
//
//	Example:
//	FindAllBy(&User, map[string]interface{}{}, "created_at desc") // will get all User
//	FindAllBy(&User, map[string]interface{}{"name":"XXX"}, "created_at desc") // will get all User with name = "XXX"
func (g *GormRepository) FindAllBy(model interface{}, by map[string]interface{}, orderBy string) (interface{}, error) {
	err := g.Db.Where(by).Order(orderBy).Unscoped().Find(&model).Error
	return model, err
}

// Find any records by using the row name and the data. Preload will get all associations in the model
//
//	Example:
//	FindAllByWithPreload(&User, map[string]interface{}{}, "created_at desc") // will get all User
//	FindAllByWithPreload(&User, map[string]interface{}{"name":"XXX"}, "created_at desc") // will get all User with name = "XXX"
func (g *GormRepository) FindAllByWithPreload(model interface{}, by map[string]interface{}, orderBy string) (interface{}, error) {
	err := g.Db.Preload(clause.Associations).Where(by).Order(orderBy).Unscoped().Find(&model).Error
	return model, err
}

// Find any records by using the row name and the data. Preload will get all associations in the model
//
//	Example:
//	FindAllByWithPreloadNestedPreload(&User, map[string]interface{}{}, "created_at desc", "User.Role") // will get all User
//	FindAllByWithPreloadNestedPreload(&User, map[string]interface{}{"name":"XXX"}, "created_at desc", "User.Role") // will get all User with name = "XXX"
func (g *GormRepository) FindAllByWithNestedPreload(model interface{}, by map[string]interface{}, orderBy, nestedPreload string) (interface{}, error) {
	err := g.Db.Preload(nestedPreload).Preload(clause.Associations).Where(by).Order(orderBy).Unscoped().Find(&model).Error
	return model, err
}

// Find any records by using the row name and the data. This will limit the result to specific number.
//
//	Example:
//	FindAllByWithPagination(&User, map[string]interface{}{}, 1, 10, "created_at desc") // will get the first 10 User in the column
//	FindAllByWithPagination(&User, map[string]interface{}{}, 2, 10, "created_at desc") // will get the next 10 User in the column
//	FindAllByWithPagination(&User, map[string]interface{}{"name":"XXX"}, 2, 10, "created_at desc") // will get the first 10 User in the column with name = "XXX"
func (g *GormRepository) FindAllByWithPagination(model interface{}, by map[string]interface{}, page, itemPerPage int, orderBy string) (interface{}, error) {
	err := g.Db.Where(by).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&model).Error
	return model, err
}

// Find any records by using the row name and the data. This will limit the result to specific number and will get all associations in the model
//
//	Example:
//	FindAllByWithPreloadAndPagination(&User, map[string]interface{}{}, 1, 10, "created_at desc") // will get the first 10 User in the column
//	FindAllByWithPreloadAndPagination(&User, map[string]interface{}{}, 2, 10, "created_at desc") // will get the next 10 User in the column
//	FindAllByWithPreloadAndPagination(&User, map[string]interface{}{"name":"XXX"}, 2, 10, "created_at desc") // will get the first 10 User in the column with name = "XXX"
func (g *GormRepository) FindAllByWithPreloadAndPagination(model interface{}, by map[string]interface{}, page, itemPerPage int, orderBy string) (interface{}, error) {
	err := g.Db.Preload(clause.Associations).Where(by).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&model).Error
	return model, err
}

// Find any records by using the row name and the data. This will limit the result to specific number and will get all associations in the model
//
//	Example:
//	FindAllByWithNestedPreloadAndPagination(&User, map[string]interface{}{}, 1, 10, "created_at desc", "User.Role") // will get the first 10 User in the column
//	FindAllByWithNestedPreloadAndPagination(&User, map[string]interface{}{}, 2, 10, "created_at desc", "User.Role") // will get the next 10 User in the column
//	FindAllByWithNestedPreloadAndPagination(&User, map[string]interface{}{"name":"XXX"}, 2, 10, "created_at desc", "User.Role") // will get the first 10 User in the column with name = "XXX"
func (g *GormRepository) FindAllByWithNestedPreloadAndPagination(model interface{}, by map[string]interface{}, page, itemPerPage int, orderBy, nestedPreload string) (interface{}, error) {
	err := g.Db.Preload(nestedPreload).Preload(clause.Associations).Where(by).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&model).Error
	return model, err
}

// Find any records by using custom SQL Query.
//
//	Example:
//	FindAllUsingCustomQuery(&User, "name = 'XXX' AND email == 'YYY'", "created_at desc") // will get all User in the column with name = "XXX" and email = "YYY"
func (g *GormRepository) FindAllUsingCustomQuery(model interface{}, query, orderBy string) (interface{}, error) {
	err := g.Db.Where(query).Order(orderBy).Unscoped().Find(&model).Error
	return model, err
}

// Find any records by using custom SQL Query. Preload will get all associations in the model
//
//	Example:
//	FindAllUsingCustomQueryWithPreload(&User, "name = 'XXX' AND email == 'YYY'", "created_at desc") // will get all User in the column with name = "XXX" and email = "YYY"
func (g *GormRepository) FindAllUsingCustomQueryWithPreload(model interface{}, query, orderBy string) (interface{}, error) {
	err := g.Db.Preload(clause.Associations).Where(query).Order(orderBy).Unscoped().Find(&model).Error
	return model, err
}

// Find any records by using custom SQL Query. Preload will get all associations in the model
//
//	Example:
//	FindAllUsingCustomQueryWithNestedPreload(&User, "name = 'XXX' AND email == 'YYY'", "created_at desc", "User.Role") // will get all User in the column with name = "XXX" and email = "YYY"
func (g *GormRepository) FindAllUsingCustomQueryWithNestedPreload(model interface{}, query, orderBy, nestedPreload string) (interface{}, error) {
	err := g.Db.Preload(nestedPreload).Preload(clause.Associations).Where(query).Order(orderBy).Unscoped().Find(&model).Error
	return model, err
}

// Find any records by using any SQL Query. This will limit the result to specific number.
//
//	Example:
//	FindAllUsingCustomQueryWithPagination(&User, "name = 'XXX' AND email = 'YYY'", "created_at desc", 1, 10) // will get the first 10 User in the column with name = "XXX" and email = "YYY"
//	FindAllUsingCustomQueryWithPagination(&User, "name = 'XXX' AND email = 'YYY'", "created_at desc", 2, 10) // will get the next 10 User in the column with name = "XXX" and email = "YYY"
func (g *GormRepository) FindAllUsingCustomQueryWithPagination(model interface{}, query, orderBy string, page, itemPerPage int) (interface{}, error) {
	err := g.Db.Where(query).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&model).Error
	return model, err
}

// Find any records by using any SQL Query. This will limit the result to specific number and will get all associations in the model
//
//	Example:
//	FindAllUsingCustomQueryWithPreloadAndPagination(&User, "name = 'XXX' AND email = 'YYY'", "created_at desc", 1, 10) // will get the first 10 User in the column with name = "XXX" and email = "YYY"
//	FindAllUsingCustomQueryWithPreloadAndPagination(&User, "name = 'XXX' AND email = 'YYY'", "created_at desc", 2, 10) // will get the next 10 User in the column with name = "XXX" and email = "YYY"
func (g *GormRepository) FindAllUsingCustomQueryWithPreloadAndPagination(model interface{}, query, orderBy string, page, itemPerPage int) (interface{}, error) {
	err := g.Db.Preload(clause.Associations).Where(query).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&model).Error
	return model, err
}

// Find any records by using any SQL Query. This will limit the result to specific number and will get all associations in the model
//
//	Example:
//	FindAllUsingCustomQueryWithPreloadAndPagination(&User, "name = 'XXX' AND email = 'YYY'", "created_at desc", 1, 10) // will get the first 10 User in the column with name = "XXX" and email = "YYY"
//	FindAllUsingCustomQueryWithPreloadAndPagination(&User, "name = 'XXX' AND email = 'YYY'", "created_at desc", 2, 10) // will get the next 10 User in the column with name = "XXX" and email = "YYY"
func (g *GormRepository) FindAllUsingCustomQueryWithNestedPreloadAndPagination(model interface{}, query, orderBy, nestedPreload string, page, itemPerPage int) (interface{}, error) {
	err := g.Db.Preload(nestedPreload).Preload(clause.Associations).Where(query).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&model).Error
	return model, err
}
