package main

import (
	postgresql "dbo/database"
	"dbo/engine"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Swagger for [DBO Technical Test]
// @version 1.0
// @description This is a document API DBO Technical Test Backend Position
// @host localhost
// @BasePath /
// @schemes http
// @accept json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	db := postgresql.PostsqlConn()
	r := gin.Default()
	engine.SetupRouter(r, db)
	// server
	if err := r.Run(":80"); err != nil {
		log.Println(err)
		return
	}
}
