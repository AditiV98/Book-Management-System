package models

import "github.com/google/uuid"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserID    string `json:"userID"`
	Address   string `json:"address"`
	PhoneNo   string `json:"phoneNo"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type Users struct {
	ID        int       `json:"-"`
	FirstName string    `json:"given_name"`
	LastName  string    `json:"family_name"`
	UserID    string    `json:"email"`
	Picture   string    `json:"picture"`
	CreatedAt int       `json:"iat"`
	LoggedAt  int       `json:"loggedAt"`
	Token     uuid.UUID `json:"token"`
	Verified  bool      `json:"email_verified,omitempty"`
	Error     ErrorMsg  `json:"error,omitempty"`
}

type ErrorMsg struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

// UserInput defines attributes which supports the login of a user
type UserInput struct {
	AccessToken string `json:"accessCode"`
}
