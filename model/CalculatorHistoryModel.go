package model

import "gorm.io/gorm"

// CalculatorHistoryModel - Model for the calculator_history table
type CalculatorHistoryModel struct {
	Id             int    `json:"id" orm:"auto"`
	UserName       string `json:"user_name" orm:"size(32)"`
	CalculatorType string `json:"calculator_type" orm:"size(32)"`
	Params         string `json:"params" orm:"size(128)"`
	Results        string `json:"result" orm:"size(64)"`
}

// create query methods for this model

// add create query method here
func CreateCalculatorHistoryModelInstance(db *gorm.DB, model *CalculatorHistoryModel) *gorm.DB {
	result := db.Create(&model)
	return result
}

// add get list of all instances query method here
func GetAllCalculatorHistoryModelInstances(db *gorm.DB, model *[]CalculatorHistoryModel) *gorm.DB {
	result := db.Find(&model)
	return result
}

// add get by last element query method here
func GetLastCalculatorHistoryModelInstance(db *gorm.DB, model *CalculatorHistoryModel) *CalculatorHistoryModel {
	db.Last(&model)
	return model
}
