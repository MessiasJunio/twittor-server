package models

// LoginResponse token that returns with login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
