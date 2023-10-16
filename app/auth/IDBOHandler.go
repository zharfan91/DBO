package app

import "github.com/gin-gonic/gin"

type IDBOHandler interface {
	Login(c *gin.Context)
}
