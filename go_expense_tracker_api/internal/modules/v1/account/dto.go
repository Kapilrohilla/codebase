package account

type Create struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateAdmin struct {
	Create
}

type GetAccounts struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetAccountQuery struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type UpdateAccount struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"omitempty,email"`
	Phone string `json:"phone"`
}
