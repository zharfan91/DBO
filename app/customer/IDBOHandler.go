package app

import "github.com/gin-gonic/gin"

type IDBOHandler interface {
	CreateCustomer(c *gin.Context)
	GetCustomer(c *gin.Context)
	GetCustomerDetail(c *gin.Context)
	DeleteCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
}
