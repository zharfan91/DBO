package repository

import (
	app "dbo/app/customer"
	"dbo/model"
	"log"
	"math"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBORepository struct {
	db *gorm.DB
}

func NewDBORepository(db *gorm.DB) app.IDBORepository {
	return &DBORepository{db: db}
}
func (r *DBORepository) CreateCustomer(request model.DBOCustomer) error {
	err := r.db.Create(&request)
	return err.Error
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (r *DBORepository) GetCustomer(pagination model.Pagination) (*model.Pagination, error) {
	var Customer model.DBOCustomer
	var Customers []*model.Customer
	var totalRows int64
	tx := r.db.Scopes(Paginate(&pagination, r.db)).Model(Customer).Where(" deleted_at is null").Find(&Customers)
	res := r.db.Model(Customer).Where(" deleted_at is null").Count(&totalRows)
	log.Println("rep totalRows", totalRows)
	if err := tx.Error; err != nil {
		return nil, err
	}
	if err := res.Error; err != nil {
		return nil, err
	}
	pagination.Rows = Customers
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages
	return &pagination, nil
}
func (r *DBORepository) GetCustomerDetail(customerID string) (*model.Customer, error) {
	var customers model.DBOCustomer
	var customer model.Customer
	tx := r.db.Model(customers).Where("id = ? and deleted_at is null", customerID).Find(&customer)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func Paginate(pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
func (r *DBORepository) DeleteCustomer(customerID string) bool {
	var customer model.DBOCustomer
	err := r.db.Where("id = ?", customerID).Delete(&customer).Debug().Error
	if err != nil {
		return false
	}
	return true
}
func (r *DBORepository) UpdateCustomer(request model.DBOCustomer, customerID string) error {
	var customer model.DBOCustomer
	if request.Password != "" {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
		request.Password = string(bytes)
	}
	err := r.db.Model(&customer).Where("id=?", customerID).Updates(&request).Debug().Error
	if err != nil {
		return err
	}
	return nil
}
