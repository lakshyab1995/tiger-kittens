package db

import (
	"log"

	"github.com/lakshyab1995/tiger-kittens/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) (*jwt.TokenModel, error)
	AuthenticateUser(username string, password string) bool
	GetUsrIdByUsername(username string) (string, error)
}

type userrepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userrepository{db: db}
}

func (r *userrepository) Create(user *User) (*jwt.TokenModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}

	return token, r.db.Create(user).Error
}

func (r *userrepository) AuthenticateUser(username string, password string) bool {
	var hashedPwd string
	if err := r.db.Model(&User{}).Select("password").Where("username = ?", username).First(&hashedPwd).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			log.Println(err)
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
	return err == nil
}

func (r *userrepository) GetUsrIdByUsername(username string) (string, error) {
	var usrId string
	err := r.db.Model(&User{}).Select("id").Where("username = ?", username).First(&usrId).Error
	if err != nil {
		log.Println(err)
		return usrId, err
	}
	return usrId, nil
}
