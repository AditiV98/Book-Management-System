package book

import (
	"Book_Management_System/models"
	"Book_Management_System/services"
	"context"
	"time"
)

type service struct {
	storeBook services.Repository
}

func New(storeBook services.Repository) service {
	return service{storeBook: storeBook}
}

func (s service) Create(ctx context.Context, book *models.Book) (*models.Book, string, error) {
	currentTime := time.Now().Unix()

	book.CreatedAt = int(currentTime)

	createBook, err := s.storeBook.Create(ctx, book)
	if err != nil {
		return nil, "", err
	}

	//msg, err := s.SendSMS(book.Admin)
	//if err != nil {
	//	return nil, "", errors.New("error from sms service")
	//}

	return createBook, "", nil
}

//
//func (s service) SendSMS( admin models.Admin) (string, error) {
//	body, err := json.Marshal(admin)
//	if err != nil {
//		return "", err
//	}
//
//	resp, err := s.httpService.PostWithHeaders("/sms", nil, body, nil)
//	if err != nil {
//		return "", err
//	}
//
//	// checking status code from the DataService response.
//	if resp.StatusCode != http.StatusCreated {
//		return "", err
//	}
//
//	return "message sent successfully", nil
//}

func (s service) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	currentTime := time.Now().Unix()

	book.UpdatedAt = int(currentTime)

	updateBook, err := s.storeBook.Update(ctx, book)
	if err != nil {
		return nil, err
	}

	return updateBook, nil
}

func (s service) GetAll(ctx context.Context) ([]models.Book, error) {
	books, err := s.storeBook.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s service) GetByID(ctx context.Context, id int) (*models.Book, error) {
	book, err := s.storeBook.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.storeBook.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
