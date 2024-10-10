package basic

type DoRegisterOfEMail struct {
	EMail     string `json:"email" form:"email" valid:"required,max=60,email" label:"E-Mail"`
	FirstName string `json:"first_name" form:"first_name" valid:"required,max=64" label:"First Name"`
	LastName  string `json:"last_name" form:"last_name" valid:"required,max=64" label:"Last Name"`
	Password  string `json:"password" form:"password" valid:"required,password" label:"Password"`
}
