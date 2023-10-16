package repository

import (
	app "dbo/app/order"
	"dbo/model"
	"log"
	"math"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBORepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) app.IDBORepository {
	return &DBORepository{db: db}
}
func (r *DBORepository) CreateOrder(request model.DBOrder) error {
	err := r.db.Create(&request)
	return err.Error
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (r *DBORepository) GetOrder(pagination model.Pagination) (*model.Pagination, error) {
	var Order model.DBOrder
	var Orders []*model.DBOrder
	var totalRows int64
	tx := r.db.Scopes(Paginate(&pagination, r.db)).Model(Order).Where(" deleted_at is null").Find(&Orders)
	res := r.db.Model(Order).Where(" deleted_at is null").Count(&totalRows)
	log.Println("rep totalRows", totalRows)
	if err := tx.Error; err != nil {
		return nil, err
	}
	if err := res.Error; err != nil {
		return nil, err
	}
	pagination.Rows = Orders
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages
	return &pagination, nil
}
func (r *DBORepository) GetOrderDetail(orderID string) (*model.DBOrder, error) {
	var order model.DBOrder
	tx := r.db.Model(order).Where("id = ? and deleted_at is null", orderID).Find(&order)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func Paginate(pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
func (r *DBORepository) DeleteOrder(orderID string) bool {
	var order model.DBOrder
	err := r.db.Where("id = ?", orderID).Delete(&order).Debug().Error
	if err != nil {
		log.Println("DeleteOrder", err)
		return false
	}
	return true
}
func (r *DBORepository) UpdateOrder(request model.DBOrder, orderID string) error {
	var order model.DBOrder
	err := r.db.Model(&order).Where("id=?", orderID).Updates(&request).Debug().Error
	if err != nil {
		return err
	}
	return nil
}
