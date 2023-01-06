package models

type Response[T any] struct {
	Data    T             `json:"data"`
	Error   ResponseErrpr `json:"error"`
	Success bool          `json:"success"`
}

type ResponseErrpr struct {
	Code int `json:"code"`
}
