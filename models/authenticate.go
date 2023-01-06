package models

type AuthenticateRequest struct {
	Account string `url:"account"`
	Passwd  string `url:"passwd"`
	Session string `url:"session"`
	Format  string `url:"format"`
	OtpCode string `url:"otp_code,omitempty"`
}

type AuthenticateResponse struct {
	Sid string `json:"sid"`
}
