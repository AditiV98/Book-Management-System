package bookIssue

import (
	"Book_Management_System/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	insertQuery  = "insert into book_issue(book_id,user_id,issued_at,submitted_at) values (?,?,?,?)"
	updateQuery  = "update book_issue set submitted_at=? where id=?"
	getAllQuery  = "select * from book_issue"
	getByIDQuery = "select id,book_id,user_id,issued_at,submitted_at from book_issue where id=?"
	deleteQuery  = "delete from book_issue where id=?"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{
		db: db,
	}
}

func (s store) Create(ctx context.Context, book *models.Issue) (*models.Issue, error) {
	var values []interface{}

	values = append(values, book.BookID, book.UserID, book.IssuedAt, 0)

	_, err := s.db.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	return book, nil
}

func (s store) Update(ctx context.Context, id int) error {
	var values []interface{}

	values = append(values, int(time.Now().Unix()), id)

	_, err := s.db.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return errors.New("query execution failed")
	}

	return nil
}

func (s store) GetAll(ctx context.Context) ([]models.Issue, error) {
	response := make([]models.Issue, 0)

	dbResult, err := s.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	defer dbResult.Close()

	for dbResult.Next() {
		var book models.Issue

		err = dbResult.Scan(&book.Id, &book.BookID, &book.UserID, &book.IssuedAt, &book.SubmittedAt)
		if err != nil {
			return nil, errors.New("scan error")
		}

		response = append(response, book)
	}

	if len(response) == 0 {
		return nil, errors.New("no books")
	}

	return response, nil
}

func (s store) GetByID(ctx context.Context, id int) (*models.Issue, error) {
	var book models.Issue

	dbResult := s.db.QueryRowContext(ctx, getByIDQuery, id)

	err := dbResult.Scan(&book.Id, &book.BookID, &book.UserID, &book.IssuedAt, &book.SubmittedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("no book issued")
	} else if err != nil {
		return nil, errors.New("scan error")
	}

	return &book, nil
}

func (s store) Delete(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return errors.New("query execution failed")
	}

	return nil
}
