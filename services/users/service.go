package users

import (
	"Book_Management_System/models"
	"Book_Management_System/services"
	"context"
	"time"
)

type service struct {
	userStore services.Users
}

func New(userService services.Users) service {
	return service{userStore: userService}
}

func (svc service) Create(ctx context.Context, user *models.User) (*models.User, error) {
	currentTime := time.Now().Unix()

	user.CreatedAt = int(currentTime)

	response, err := svc.userStore.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (svc service) Update(ctx context.Context, user *models.User) (*models.User, error) {
	currentTime := time.Now().Unix()

	user.UpdatedAt = int(currentTime)

	updateUser, err := svc.userStore.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (svc service) GetAll(ctx context.Context) ([]models.User, error) {
	users, err := svc.userStore.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc service) GetByID(ctx context.Context, id int) (*models.User, error) {
	user, err := svc.userStore.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc service) Delete(ctx context.Context, id int) error {
	err := svc.userStore.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
