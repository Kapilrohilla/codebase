package split

type CreateSplit struct {
	Amount    uint `json:"amount" validate:"required,min=1"`
	AccountId uint `json:"account_id" validate:"required,min=1"`
	ExpenseId uint `json:"expense_id" validate:"required,min=1"`
}

type GetQuery struct {
	Page  uint
	Limit uint
}
