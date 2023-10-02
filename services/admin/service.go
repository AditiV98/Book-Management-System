package admin

import (
	"Book_Management_System/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"regexp"
	"time"
)

const (
	TokenExpiryTime = 86400
	emailFormat     = `^[a-zA-Z0-9]+[a-zA-Z0-9._+!#$%&*/=?^\-]*[a-zA-Z0-9]*@[a-zA-Z0-9\-][a-zA-Z0-9.\-]*[.][a-z]+$`
)

type service struct {
	adminStore   Repository
	tokenService TokenService
}

// nolint - unexported return type for exported function
func New(adminStore Repository, tokenService TokenService) service {
	return service{
		adminStore:   adminStore,
		tokenService: tokenService,
	}
}

func (svc service) Login(c context.Context, admin *models.UserInput) (*models.Admin, error) {
	adminInfo, err := svc.tokenService.CheckUserAuthorized(admin.AccessToken)
	if err != nil {
		return nil, err
	}

	if !ValidateEmail(adminInfo.UserID) {
		return nil, errors.New("please check your access code as we are not able to fetch corresponding user details")
	}

	bearerToken, _ := uuid.NewUUID()
	adminProfile := &models.Admin{
		FirstName: adminInfo.FirstName,
		LastName:  adminInfo.LastName,
		UserID:    adminInfo.UserID,
		Picture:   adminInfo.Picture,
	}

	adminDetails, err := svc.adminStore.GetByEmail(c, adminProfile.UserID)
	if err == errors.New("admin details not found") {
		adminProfile.Token = bearerToken
		return svc.adminStore.Create(c, adminProfile)
	} else if err != nil {
		return nil, err
	}

	if !isPreviousTokenValid(adminDetails.LoggedAt) || adminDetails.Token == uuid.Nil {
		adminDetails.Token = bearerToken
	}

	adminDetails.LoggedAt = int(time.Now().UTC().Unix())

	err = svc.Update(c, adminDetails)
	if err != nil {
		return nil, err
	}

	return adminDetails, nil
}

func (svc service) Logout(c context.Context, email string) error {
	adminDetails, err := svc.adminStore.GetByEmail(c, email)
	if err != nil {
		return err
	}

	adminDetails.Token = uuid.Nil

	return svc.adminStore.Update(c, adminDetails)
}

func (svc service) Update(c context.Context, admin *models.Admin) error {
	return svc.adminStore.Update(c, admin)
}

// ValidateEmail function validates the email
func ValidateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(emailFormat)

	return emailRegex.MatchString(email)
}

// isPreviousTokenValid checks if the existing bearer token of an admin is valid and
// returns true if not expired else returns false
func isPreviousTokenValid(lastLoggedAt int) bool {
	lastLoggedAtTime := time.Unix(int64(lastLoggedAt), 0)
	difference := time.Since(lastLoggedAtTime)

	return difference <= TokenExpiryTime
}
