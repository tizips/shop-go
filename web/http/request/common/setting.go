package common

type ToSetting struct {
	Module string `query:"module" valid:"required" label:"Module"`
}
