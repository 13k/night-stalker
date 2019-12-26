package ns

import (
	nssess "github.com/13k/night-stalker/internal/processors/session"
)

// SteamCredentials contains Steam authentication information.
type SteamCredentials struct {
	// Steam API key.
	APIKey string
	// Steam username.
	Username string
	// Steam password.
	Password string
	// Steam Guard email code.
	AuthCode string
	// Steam Guard mobile two-factor authentication code.
	TwoFactorCode string
	// If set to true, ns will save a login key and next logins won't need a password.
	// ns will never save the password.
	RememberPassword bool
}

func (c *SteamCredentials) sessionCredentials() *nssess.Credentials {
	return &nssess.Credentials{
		APIKey:           c.APIKey,
		Username:         c.Username,
		Password:         c.Password,
		AuthCode:         c.AuthCode,
		TwoFactorCode:    c.TwoFactorCode,
		RememberPassword: c.RememberPassword,
	}
}
