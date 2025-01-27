package models

import (
	"errors"

	"gorm.io/gorm"
)

type Expenses struct {
	CommonModel
	Title       string  `gorm:"type:varchar(100);not null" json:"title"`
	Amount      float64 `gorm:"not null" json:"amount"`
	Description string  `gorm:"type:text;not null" json:"description"`
	AccountId   uint    `gorm:"not null" json:"account_id"`
	OwnerSplit  uint    `gorm:"not null" json:"owner_split"`
	// relationships
	Splits []Split `gorm:"foreignKey:ExpenseId;references:ID"`
}

func (u *Expenses) BeforeCreate(tx *gorm.DB) (err error) {

	if u.OwnerSplit > uint(u.Amount) {
		return errors.New("owner_split cann't be greater than Amount")
	}

	return nil
}

func (u *Expenses) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.OwnerSplit > uint(u.Amount) {
		return errors.New("owner_split cann't be greater than amount")
	}
	return nil
}
