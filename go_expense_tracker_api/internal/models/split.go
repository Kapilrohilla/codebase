package models

type Split struct {
	CommonModel
	ExpenseId   uint `gorm:"not null" json:"expense_id"`
	AccountId   uint `gorm:"not null" json:"account_id"`
	SplitAmount uint `gorm:"not null" json:"split_amount"`
}
