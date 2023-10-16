package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBOrder struct {
	Id              uuid.UUID      `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CustomerID      string         `json:"customer_id" gorm:"type:varchar(255)" validate:"required"`
	ProductID       string         `json:"product_id" gorm:"type:varchar(255)" validate:"required"`
	ShipperID       string         `json:"shipper_id" validate:"required"`
	ProductName     string         `json:"Product_name" gorm:"type:varchar(255)" validate:"required"`
	Quantity        int            `json:"quantity" validate:"required"`
	UnitPrice       int            `json:"unit_price" validate:"required"`
	OrderDate       time.Time      `json:"order_date" validate:"required"`
	ShippingAddress string         `json:"shipping_address" gorm:"type:varchar(255)" validate:"required"`
	ShippedDate     time.Time      `json:"shipped_date"`
	ReceivedDate    time.Time      `json:"received_date"`
	Status          string         `json:"status" gorm:"type:varchar(255)" validate:"required"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
type Order struct {
	Id              uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CustomerID      string    `json:"customer_id" gorm:"type:varchar(255)" validate:"required"`
	ProductID       string    `json:"product_id" gorm:"type:varchar(255)" validate:"required"`
	ShipperID       string    `json:"shipper_id" validate:"required"`
	ProductName     string    `json:"Product_name" gorm:"type:varchar(255)" validate:"required"`
	Quantity        int       `json:"quantity" validate:"required"`
	UnitPrice       int       `json:"unit_price" validate:"required"`
	OrderDate       time.Time `json:"order_date" validate:"required"`
	ShippingAddress string    `json:"shipping_address" gorm:"type:varchar(255)" validate:"required"`
	ShippedDate     time.Time `json:"shipped_date"`
	ReceivedDate    time.Time `json:"received_date"`
	Status          string    `json:"status" gorm:"type:varchar(255)" validate:"required"`
}

func (DBOrder) TableName() string { return "order" }

func (DBOrder *DBOrder) BeforeCreate(tx *gorm.DB) (err error) {
	DBOrder.CreatedAt = time.Now()
	DBOrder.UpdatedAt = time.Now()
	return
}

func (DBOrder *DBOrder) BeforeUpdate(tx *gorm.DB) (err error) {
	DBOrder.UpdatedAt = time.Now()
	return
}
