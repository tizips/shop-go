package basic

type DoUploadOfFile struct {
	Dir string `json:"dir" form:"dir" valid:"required,max=100,dirs" label:"目录"`
}
