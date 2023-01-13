package changepassword

import (
	"change_password_service/api/v1/changepassword/resp"
	"change_password_service/business/changepassword"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type changePassword struct {
	db *gorm.DB
}

func NewchangePassword(db *gorm.DB) changepassword.ChangepasswordRepo {
	return &changePassword{
		db: db,
	}
}

func (r *changePassword) FindByUserID(userID int) (*changepassword.Domain, error) {
	var record Data
	res := r.db.Table("users").Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}
	dataResult := record.toService()
	return &dataResult, nil

}
func (r *changePassword) ChangePassword(userID int, newPass string) error {
	res := r.db.Table("users").Where("id = ?", userID).Update("password", newPass)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}

	return nil

}

func (r *changePassword) CheckEmail(email string) (*resp.DataForgot, error) {
	var record resp.DataForgot

	res := r.db.Table("users").Where("email = ?", email).Scan(&record)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("email not found")
	}

	return &record, nil
}

func (r *changePassword) CheckVerif(header string) (res *resp.DataForgot, err error) {
	q := r.db.Table("public.users u ").Where("u.email_verification = ?", header).Scan(&res)
	if q.Error != nil {
		return nil, errors.New("something went wrong")
	}

	if q.RowsAffected == 0 {
		return nil, errors.New("you are not authorized")
	}

	return res, err
}

func (r *changePassword) ResetPassword(id int, pass string) error {
	data := map[string]interface{}{
		"password":           pass,
		"email_verification": nil,
	}
	if err := r.db.Table("users").Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *changePassword) SaveVerif(d resp.DataReset) error {
	if err := r.db.Table("users").Where("id = ?", d.Id).Update("email_verification", d.Uid).Error; err != nil {
		return err
	}
	return nil
}
