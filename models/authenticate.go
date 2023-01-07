package models

type ApiCredential struct {
	Account string `url:"account"`
	Passwd  string `url:"passwd"`
	Session string `url:"session,omitempty"`
	Format  string `url:"format,omitempty"`
	OtpCode string `url:"otp_code,omitempty"`
	sid     string
}

type AuthenticateResponse struct {
	Sid string `json:"sid"`
	Did string `json:"did"`
}
