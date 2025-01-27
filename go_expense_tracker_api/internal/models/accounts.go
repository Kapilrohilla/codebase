package models

type Accounts struct {
	CommonModel
	Name            string `gorm:"type:varchar(50);not null" json:"name"`
	Email           string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Phone           string `gorm:"type:varchar(10);not null;unique" json:"phone"`
	IsEmailVerified bool   `gorm:"type:boolean;default:false" json:"is_email_verified"`
	IsPhoneVerified bool   `gorm:"type:boolean;default:false" json:"is_phone_verified"`
	Password        string `gorm:"type:text;not null" json:"-"`
	IsAdmin         bool   `gorm:"type:boolean;default:false" json:"is_admin"`
	// relationships
	Expenses []Expenses `gorm:"foreignKey:AccountId;references:ID"`
	Splits   []Split    `gorm:"foreignKey:AccountId;references:ID"`
}
