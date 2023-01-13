package changepassword

import "time"

type Result struct {
	// TotalAllData int
	// CurrentPage  int
	// PerPage      int
	// TotalPage    int
	Domain Domain
}
type Domain struct {
	ID                 int
	Branch_id          int
	Privilage          int
	Fullname           string
	Company_name       string
	Email              string
	Phone              string
	Address            string
	Username           string
	Password           string
	Status             int
	Last_login         time.Time
	Forgot_key         string
	Created_at         time.Time
	Created_by         int
	Updated_at         time.Time
	Updated_by         int
	Deleted_at         time.Time
	Deleted_by         int
	Type               int
	Remark             string
	Email_verification string
}
