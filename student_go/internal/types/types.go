package types

import "time"

type Student struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" `
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAT time.Time `json:"createdAt"`
}
