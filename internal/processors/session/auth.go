package session

type Credentials struct {
	APIKey           string
	Username         string
	Password         string
	AuthCode         string
	TwoFactorCode    string
	RememberPassword bool
}
