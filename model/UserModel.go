package model

import "gorm.io/gorm"

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
