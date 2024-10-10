package basic

type DoLoginOfAccount struct {
	Username string `json:"username" form:"username" valid:"required,username" label:"用户名"`
	Password string `json:"password" form:"password" valid:"required,password" label:"密码"`
}
