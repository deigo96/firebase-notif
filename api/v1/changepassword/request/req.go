package request

type Req struct {
	OldPass string `json:"old_password" form:"old_password" binding:"required,min=6"`
	NewPass string `json:"new_password" form:"new_password"  binding:"required,min=6"`
}

type ReqEmail struct {
	Email string `json:"email"`
}

type ReqResetPassword struct {
	Password        string `json:"password" binding:"required, min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required, min=6"`
}
