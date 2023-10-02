package models

import "github.com/google/uuid"

// Admin defines the attributes explaining the personal information of a
// admin along with some configurations
type Admin struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"given_name"`
	LastName      string    `json:"family_name"`
	UserID        string    `json:"email"`
	Picture       string    `json:"picture"`
	CreditBalance float64   `json:"creditBalance"`
	Token         uuid.UUID `json:"token"`
	CreatedAt     int       `json:"createdAt"`
	LoggedAt      int       `json:"loggedAt"`
}
