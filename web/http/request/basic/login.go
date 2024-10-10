package basic

type DoLoginOfAccount struct {
	Username string `json:"username" form:"username" valid:"required,username" label:"Username"`
	Password string `json:"password" form:"password" valid:"required,password" label:"Password"`
}

type DoLoginOfEMail struct {
	EMail    string `json:"email" form:"email" valid:"required,max=60,email" label:"E-Mail"`
	Password string `json:"password" form:"password" valid:"required,password" label:"Password"`
}
