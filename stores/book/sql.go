package book

import (
	"Book_Management_System/models"
	"context"
	"database/sql"
	"errors"
)

const (
	insertQuery  = "insert into book(book_name,author,publication_year,genre,availability,created_at,updated_at) values (?,?,?,?,?,?,?)"
	updateQuery  = "update book set book_name=?,author=?,publication_year=?,updated_at=? where id=?"
	getAllQuery  = "select * from book"
	getByIDQuery = "select id,book_name,author,publication_year,updated_at,created_at from book where id=?"
	deleteQuery  = "delete from book where id=?"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) Create(ctx context.Context, book *models.Book) (*models.Book, error) {
	var values []interface{}

	values = append(values, book.BookName, book.Author, book.PublicationYear, book.Genre, book.Availability, book.CreatedAt, 0)

	_, err := s.db.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	return book, nil
}

func (s store) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	var values []interface{}

	values = append(values, book.BookName, book.Author, book.PublicationYear, book.UpdatedAt, book.ID)

	_, err := s.db.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	return book, nil
}

func (s store) GetAll(ctx context.Context) ([]models.Book, error) {
	response := make([]models.Book, 0)

	dbResult, err := s.db.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, errors.New("query execution failed")
	}

	defer dbResult.Close()

	for dbResult.Next() {
		var book models.Book

		err = dbResult.Scan(&book.ID, &book.BookName, &book.Author, &book.PublicationYear, &book.UpdatedAt, &book.CreatedAt, &book.Genre, &book.Availability)
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

func (s store) GetByID(ctx context.Context, id int) (*models.Book, error) {
	var book models.Book

	dbResult := s.db.QueryRowContext(ctx, getByIDQuery, id)

	err := dbResult.Scan(&book.ID, &book.BookName, &book.Author, &book.PublicationYear, &book.UpdatedAt, &book.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("no book")
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
