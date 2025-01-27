package types

type IError struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
}
