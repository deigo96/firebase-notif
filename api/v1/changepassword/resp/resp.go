package resp

// import (
// 	"change_password_service/business/changepassword"
// )

// type Resp struct {
// 	// TotalAllData int
// 	// CurrentPage  int
// 	// PerPage      int
// 	// TotalPage    int
// 	Result RespDomain `json:"result"`
// }

// type RespDomain struct {
// 	ID int
// }

// func FromService(v changepassword.Result) Resp {
// 	return Resp{
// 		Result: FromServiceDomain(v.Domain),
// 	}
// }
// func FromServiceDomain(u changepassword.Domain) RespDomain {
// 	return RespDomain{
// 		ID: int(u.ID),
// 	}
// }

type DataForgot struct {
	Id       int
	Email    string
	Username string
}

type DataReset struct {
	Id int
	Uid string
	Email string
}
