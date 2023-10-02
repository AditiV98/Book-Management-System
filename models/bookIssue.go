package models

type Issue struct {
	Id          int `json:"id"`
	BookID      int `json:"bookID"`
	UserID      int `json:"userID"`
	IssuedAt    int `json:"issuedAt"`
	SubmittedAt int `json:"submittedAt"`
}
