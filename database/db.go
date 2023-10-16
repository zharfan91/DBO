package postgresql

import (
	"dbo/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

var (
	psqlConn *gorm.DB
	err      error
)

// initialize database
func init() {
	setupPostsqlConn()
}

func setupPostsqlConn() {
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("TIME_ZONE"))
	psqlConn, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to create a connection to database error:%s", err)
	}

	psqlConn.AutoMigrate(model.DBOrder{}, model.DBOCustomer{})
}

func PostsqlConn() *gorm.DB {
	return psqlConn
}
