package app

import "github.com/gin-gonic/gin"

type IDBOHandler interface {
	CreateOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	GetOrderDetail(c *gin.Context)
	DeleteOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
}
