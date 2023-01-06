package models

import (
	"github.com/sirupsen/logrus"
)

type ApiCredential struct {
	Account string `url:"account"`
	Passwd  string `url:"passwd"`
	Session string `url:"session,omitempty"`
	Format  string `url:"format,omitempty"`
	OtpCode string `url:"otp_code,omitempty"`
	sid     string
}

func (a *ApiCredential) SetSID(sid string) {
	logrus.Info("SID is set.")
	a.sid = sid
}

func (a *ApiCredential) GetSID() string {
	return a.sid
}

type AuthenticateResponse struct {
	Sid string `json:"sid"`
	Did string `json:"did"`
}
