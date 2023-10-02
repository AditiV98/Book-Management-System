package admin

import (
	"Book_Management_System/models"
	"context"
)

// Repository is an abstraction of the methods from admin store that can be consumed
type Repository interface {
	GetByEmail(ctx context.Context, email string) (*models.Admin, error)
	Create(ctx context.Context, admin *models.Admin) (*models.Admin, error)
	Update(ctx context.Context, admin *models.Admin) error
}

// TokenService is an abstraction of the methods from token service that can be consumed
type TokenService interface {
	CheckUserAuthorized(authToken string) (*models.Users, error)
}
