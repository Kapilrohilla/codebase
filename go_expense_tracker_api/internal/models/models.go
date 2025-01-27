package models

import (
	"time"

	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
)

func MigrateModels(dbs storage.DBs) {
	dbs.Migrate(Accounts{})
	dbs.Migrate(Expenses{})
	dbs.Migrate(Split{})
}

type CommonModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:datetime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime" json:"-"`
}
