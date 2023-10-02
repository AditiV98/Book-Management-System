package users

import (
	"Book_Management_System/models"
	"context"
	"database/sql"
	"errors"
)

const (
	insertQuery  = "insert into users (first_name,last_name,user_id,address,phone_no,created_at,updated_at) values(?,?,?,?,?,?,?)"
	updateQuery  = "update users set first_name=?,last_name=?,user_id=?,address=?,phone_no=?,updated_at=? where id=?"
	getAllQuery  = "select * from users"
	getByIDQuery = "select id,first_name,last_name,user_id,address,phone_no,created_at from users where id=?"
	deleteQuery  = "delete from users where id=?"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{
		db: db,
	}
}

func (s store) Create(ctx context.Context, user *models.User) (*models.User, error) {
	var values []interface{}

	values = append(values, user.FirstName, user.LastName, user.UserID, user.Address, user.PhoneNo, user.CreatedAt, 0)

	_, err := s.db.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	return user, nil
}

func (s store) Update(ctx context.Context, user *models.User) (*models.User, error) {
	var values []interface{}

	values = append(values, user.FirstName, user.LastName, user.UserID, user.Address, user.PhoneNo, user.UpdatedAt, user.ID)

	_, err := s.db.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	return user, nil
}

func (s store) GetAll(ctx context.Context) ([]models.User, error) {
	response := make([]models.User, 0)

	dbResult, err := s.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	defer dbResult.Close()

	for dbResult.Next() {
		var user models.User

		err = dbResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserID, &user.Address, &user.PhoneNo, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, errors.New("scan error")
		}

		response = append(response, user)
	}

	if len(response) == 0 {
		return nil, errors.New("no users")
	}

	return response, nil
}

func (s store) GetByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	dbResult := s.db.QueryRowContext(ctx, getByIDQuery, id)

	err := dbResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserID, &user.Address, &user.PhoneNo, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user")
	} else if err != nil {
		return nil, errors.New("scan error")
	}

	return &user, nil
}

func (s store) Delete(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return errors.New("query execution failed")
	}

	return nil
}
