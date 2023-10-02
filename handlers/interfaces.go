package handlers

import (
	"Book_Management_System/models"
	"context"
)

type BookSvc interface {
	Create(ctx context.Context, book *models.Book) (*models.Book, string, error)
	Update(ctx context.Context, book *models.Book) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
	GetByID(ctx context.Context, id int) (*models.Book, error)
	Delete(ctx context.Context, id int) error
}

type UserSvc interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
}

type BookIssue interface {
	Create(ctx context.Context, book *models.Issue) (*models.Issue, error)
	Update(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Issue, error)
	GetByID(ctx context.Context, id int) (*models.Issue, error)
	Delete(ctx context.Context, id int) error
}

type AdminSvc interface {
	Login(c context.Context, admin *models.UserInput) (*models.Admin, error)
	Logout(c context.Context, email string) error
}
