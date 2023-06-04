package model

import "gorm.io/gorm"

// Test Model - Model for the testModel table
type TestModel struct {
	UserId   int    `json:"user_id" orm:"auto"`
	Email    string `json:"email" orm:"size(128)"`
	Password string `json:"password" orm:"size(64)"`
	UserName string `json:"user_name" orm:"size(32)"`
}

// create query methods for this model

// add create query method here
func CreateTestModelInstance(db *gorm.DB, model *TestModel) *gorm.DB {
	result := db.Create(&model)
	return result
}

// add get by last element query method here
func GetLastTestModelInstance(db *gorm.DB, model *TestModel) *TestModel {
	db.Last(&model)
	return model
}
