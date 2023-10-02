package admin

import (
	"Book_Management_System/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	insertQuery = "insert into users (first_name,last_name,email,picture,token," +
		"created_at,logged_at) value (?,?,?,?,?,?,?)"
	selectQuery = "select id, first_name,last_name,email,picture,token,created_at," +
		"logged_at from users"
	updateQuery = "update users set first_name = ?, last_name =?, email=?, picture=?, token=?," +
		" logged_at = ?"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) GetByEmail(ctx context.Context, email string) (*models.Admin, error) {
	var adminDetails models.Admin

	query := selectQuery + " where email = ?"
	row := s.db.QueryRowContext(ctx, query, email)

	err := row.Scan(&adminDetails.ID, &adminDetails.FirstName, &adminDetails.LastName, &adminDetails.UserID,
		&adminDetails.Picture, &adminDetails.Token, &adminDetails.CreatedAt, &adminDetails.LoggedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("admin details not found")
	} else if err != nil {
		return nil, errors.New("scan error")
	}

	return &adminDetails, nil
}

func (s store) Create(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	var adminInfo []interface{}

	currentTime := time.Now().UTC()
	adminInfo = append(adminInfo, admin.FirstName, admin.LastName, admin.UserID,
		admin.Picture, admin.Token, currentTime.Unix(), currentTime.Unix())

	dbResult, err := s.db.ExecContext(ctx, insertQuery, adminInfo...)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	lastID, err := dbResult.LastInsertId()
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	admin.CreatedAt = int(currentTime.Unix())
	admin.LoggedAt = int(currentTime.Unix())
	admin.ID = int(lastID)

	return admin, nil
}

func (s store) Update(ctx context.Context, admin *models.Admin) error {
	var values []interface{}

	values = append(values, admin.FirstName, admin.LastName, admin.UserID,
		admin.Picture, admin.Token, admin.LoggedAt, admin.UserID)

	_, err := s.db.ExecContext(ctx, updateQuery+" where email =? ", values...)

	if err != nil {
		return errors.New("query execution failed")
	}

	return nil
}
