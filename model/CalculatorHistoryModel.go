package model

// CalculatorHistoryModel - Model for the calculator_history table
type CalculatorHistoryModel struct {
	Id             int    `json:"id" orm:"auto"`
	UserName       string `json:"user_name" orm:"size(32)"`
	CalculatorType string `json:"calculator_type" orm:"size(32)"`
	Params         string `json:"params" orm:"size(128)"`
	Results        string `json:"result" orm:"size(64)"`
}
