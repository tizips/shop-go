package basic

type ToAccountOfPermissions struct {
	Module string `json:"module" query:"module" form:"module" valid:"required" label:"模块"`
}

type DoAccount struct {
	Mobile   string `json:"mobile" form:"mobile" valid:"omitempty,mobile" label:"手机号"`
	Email    string `json:"email" form:"email" valid:"omitempty,email" label:"邮箱"`
	Password string `json:"password" form:"password" valid:"omitempty,password" label:"密码"`
}
