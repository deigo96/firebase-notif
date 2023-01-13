package changepassword

import (
	"change_password_service/business/changepassword"
)

type Data struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (u *Data) toService() changepassword.Domain {
	return changepassword.Domain{
		ID:       int(u.ID),
		Password: u.Password,
	}
}

func toServiceDetails(data []Data) []changepassword.Domain {
	a := []changepassword.Domain{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
