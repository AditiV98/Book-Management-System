package bookIssue

import (
	"Book_Management_System/models"
	"Book_Management_System/services"
	"context"
	"time"
)

type service struct {
	storeIssue services.BookIssue
}

func New(storeIssue services.BookIssue) service {
	return service{storeIssue: storeIssue}
}

func (s service) Create(ctx context.Context, book *models.Issue) (*models.Issue, error) {
	currentTime := time.Now().Unix()

	book.IssuedAt = int(currentTime)

	createBook, err := s.storeIssue.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	return createBook, nil
}

//func (s service) SendSMS( admin models.Admin) (string, error) {
//	body, err := json.Marshal(admin)
//	if err != nil {
//		return "", err
//	}
//
//	resp, err := s.httpService.PostWithHeaders(c, "/sms", nil, body, nil)
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

func (s service) Update(ctx context.Context, id int) error {
	err := s.storeIssue.Update(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s service) GetAll(ctx context.Context) ([]models.Issue, error) {
	books, err := s.storeIssue.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s service) GetByID(ctx context.Context, id int) (*models.Issue, error) {
	book, err := s.storeIssue.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.storeIssue.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
