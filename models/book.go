package models

type Book struct {
	ID              int    `json:"id"`
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	PublicationYear string `json:"publication_year"`
	Genre           string `json:"genre"`
	Availability    string `json:"availability"`
	CreatedAt       int    `json:"createdAt"`
	UpdatedAt       int    `json:"updatedAt"`
	//Admin           Admin  `json:"admin"`
}

//type Admin struct {
//	Number  string `json:"number"`
//	Message string `json:"message"`
//}
