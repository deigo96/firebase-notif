package changepassword

import (
	"change_password_service/api/v1/changepassword/resp"
	commonservice "change_password_service/business/common-service"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ChangepasswordRepo interface {
	FindByUserID(userID int) (*Domain, error)
	ChangePassword(userID int, newPass string) error
	CheckEmail(string) (*resp.DataForgot, error)
	ResetPassword(int, string) error
	CheckVerif(header string) (*resp.DataForgot, error)
	SaveVerif(resp.DataReset) error
}

type ChangepasswordService interface {
	ChangePassword(userID int, oldPass string, newPass string) (*Domain, error)
	VerifyCredential(userID int, password string) error
	FindByUserID(userID int) (*Domain, error)
	CheckEmailService(string) (bool, error)
	ResetPasswordService(int, string) error
	CheckVerification(header string) (*resp.DataForgot, error)
}

type changepasswordService struct {
	changepasswordRepo ChangepasswordRepo
}

func NewChangepasswordService(
	changepasswordRepo ChangepasswordRepo,
) ChangepasswordService {
	return &changepasswordService{
		changepasswordRepo: changepasswordRepo,
	}
}

func (c *changepasswordService) ChangePassword(userID int, oldPass string, newPass string) (*Domain, error) {

	u, err := c.changepasswordRepo.FindByUserID(userID)
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}
	err = c.VerifyCredential(u.ID, oldPass)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("wrong old password")
	}
	hs := hashAndSalt([]byte(newPass))

	err = c.changepasswordRepo.ChangePassword(userID, hs)
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}
	return nil, nil
}

func (c *changepasswordService) VerifyCredential(userID int, password string) error {
	user, err := c.changepasswordRepo.FindByUserID(userID)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("wrong old password")
	}
	return nil
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (c *changepasswordService) FindByUserID(userID int) (*Domain, error) {
	user, err := c.changepasswordRepo.FindByUserID(userID)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return user, nil
}

func (c *changepasswordService) CheckEmailService(email string) (bool, error) {
	status, err := c.changepasswordRepo.CheckEmail(email)
	if err != nil {
		return false, err
	}

	data := resp.DataReset{
		Id:    status.Id,
		Uid:   generateId(),
		Email: status.Email,
	}

	if err := commonservice.SendEmailNotification(data); err != nil {
		return false, errors.New("oops something went wrong! please try again later")
	}

	if err := c.changepasswordRepo.SaveVerif(data); err != nil {
		return false, errors.New("something went wrong")
	}

	return true, nil
}

func (c *changepasswordService) CheckVerification(header string) (t *resp.DataForgot, err error) {
	res, err := c.changepasswordRepo.CheckVerif(header)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *changepasswordService) ResetPasswordService(id int, pass string) error {
	passHash := hashAndSalt([]byte(pass))

	err := c.changepasswordRepo.ResetPassword(id, passHash)
	if err != nil {
		return err
	}

	return nil
}

func generateId() string {
	id := uuid.New()
	return id.String()
}
