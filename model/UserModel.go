package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserModel - Model for the user table, username and email must be unique
type UserModel struct {
	Id       int    `json:"id" orm:"auto;unique"`
	UserName string `json:"user_name" orm:"size(32);unique"`
	Email    string `json:"email" orm:"size(32);unique"`
	PassWord string `json:"password" orm:"size(32)"`
}

// create query methods for this model

// add create user query method here
func CreateUserModelInstance(db *gorm.DB, model *UserModel) *gorm.DB {
	// create new user and return the result

	result := db.Create(&model)
	return result
}

// add get user by id query method here
func GetUserModelInstanceById(db *gorm.DB, model *UserModel, id int) *gorm.DB {
	// return single based on id
	result := db.First(&model, id)
	return result
}

// add get user by user name query method here
func GetUserModelInstanceByUserName(db *gorm.DB, model *UserModel, userName string) *gorm.DB {
	// return single based on user name
	result := db.Where("user_name = ?", userName).First(&model)
	return result
}

// add get user by email query method here
func GetUserModelInstanceByEmail(db *gorm.DB, model *UserModel, email string) *gorm.DB {
	// return single based on email
	result := db.Where("email = ?", email).First(&model)
	return result
}

// password encryption methods
// func EncryptPassword(password string) string {
// 	// encrypt the password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(hashedPassword)

// }

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PassWord = string(hashedPassword)
	return nil
}

func CompareAndValidateHashedPassword(hashedPassword string, password string) bool {
	// compare the password, return true if the password is correct, return false if the password is incorrect
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
