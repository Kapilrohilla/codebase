package expense

type CreateExpense struct {
	Title       string  `json:"title" validate:"required,min=3,max=100"`        // String: min and max length
	Amount      float64 `json:"amount" validate:"required,gt=0,lt=1000000"`     // Float: greater than and less than
	Description string  `json:"description" validate:"required,min=3,max=1000"` // String: min and max length
	AccountID   uint    `json:"account_id"`                                     // Uint: required
}

type UpdateExpense struct {
	Title       string  `json:"title" validate:"min=3,max=100"`     // String: min and max length
	Amount      float64 `json:"amount" validate:"gt=-1,lt=1000000"` // Float: greater than and less than
	Description string  `json:"description" validate:"max=1000"`    // String: min and max length
}
