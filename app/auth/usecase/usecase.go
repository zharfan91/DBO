package usecase

import (
	app "dbo/app/auth"
	"dbo/model"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type DBOUsecase struct {
	repo app.IDBORepository
}

func NewAuthUsecase(repo app.IDBORepository) app.IDBOUsecase {
	return &DBOUsecase{repo: repo}
}
func (u *DBOUsecase) Login(request model.Auth) (string, error) {
	if request.Username != "admin" && request.Password != "admin" {
		res, err := u.repo.Login(request)
		if err != nil {
			return "", err
		}
		claims := jwt.MapClaims{}
		claims["name"] = res.Name
		claims["role"] = "user"
		claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	} else {
		claims := jwt.MapClaims{}
		claims["name"] = request.Username
		claims["role"] = "admin"
		claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	}
}
