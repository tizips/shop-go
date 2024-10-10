package consoles

import (
	"errors"
	"github.com/gookit/color"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/console"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"project.io/shop/model"
	"regexp"
	"strconv"
	"strings"
)

type DeveloperProvider struct {
}

func (that *DeveloperProvider) Register() console.Console {

	return console.Console{
		Cmd:  "developer",
		Name: "开发账号",
		Run: func(cmd *cobra.Command, args []string) {

			var err error
			var username, nickname, password string

			createUser := false

			if username, err = that.username(); err != nil {
				color.Errorf("\n\n用户名读取失败：%v\n\n", err)
				return
			}

			var user model.SysUser

			fu := facades.Gorm.First(&user, "`username`=?", username)

			if fu.Error != nil && !errors.Is(fu.Error, gorm.ErrRecordNotFound) {
				color.Errorf("\n\n用户信息查询失败：%v\n\n", fu.Error)
				return
			}

			if errors.Is(fu.Error, gorm.ErrRecordNotFound) {

				if password, err = that.password(); err != nil {
					color.Errorf("\n\n密码读取失败：%v\n\n", err)
					return
				}

				if nickname, err = that.nickname(); err != nil {
					color.Errorf("\n\n昵称读取失败：%v\n\n", err)
					return
				}

				createUser = true
			}

			platforms, _ := facades.Cfg.Get("auth.platforms").([]uint16)

			platform := lo.Min(platforms)

			tx := facades.Gorm.Begin()

			if createUser {

				user = model.SysUser{
					ID:       facades.Snowflake.Generate().String(),
					Username: &username,
					Nickname: nickname,
					Password: auth.Password(password),
					IsEnable: util.Yes,
				}

				if cu := tx.Create(&user); cu.Error != nil {
					tx.Rollback()
					color.Errorf("\n\n用户写入失败：%v\n\n", cu.Error)
					return
				}
			}

			role := model.SysRole{
				ID:       authConstants.CodeOfDeveloper,
				Platform: platform,
				Name:     "开发者",
				Summary:  "系统开发者暂用角色，无法修改",
			}

			if cr := tx.FirstOrCreate(&role, "`id`=?", role.ID); cr.Error != nil {
				tx.Rollback()
				color.Errorf("\n\n关联写入失败：%v\n\n", cr.Error)
				return
			}

			bind := model.SysUserBindRole{
				Platform: role.Platform,
				UserID:   user.ID,
				RoleID:   role.ID,
			}

			if cb := tx.FirstOrCreate(&bind, "`user_id`=? and `role_id`=?", user.ID, role.ID); cb.Error != nil {
				tx.Rollback()
				color.Errorf("\n\n关联写入失败：%v\n\n", cb.Error)
				return
			}

			_, err = facades.Casbin.AddRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(role.ID), strconv.Itoa(int(role.Platform)))
			if err != nil {
				tx.Rollback()
				color.Errorf("\n\n权限写入失败：%v\n\n", err)
				return
			}

			tx.Commit()

			color.Successf("\n\n开发者账号写入成功\n\n")
		},
	}
}

func (that *DeveloperProvider) username() (username string, err error) {

	prompt := promptui.Prompt{
		Label: "用户名",
		Validate: func(str string) error {

			if ok, _ := regexp.MatchString(util.PatternOfUsername, str); !ok {
				return errors.New("格式错误")
			}

			return nil
		},
	}

	if username, err = prompt.Run(); err != nil {
		return
	}

	return username, nil
}

func (that *DeveloperProvider) password() (password string, err error) {

	prompt := promptui.Prompt{
		Label: "密码",
		Validate: func(str string) error {

			if ok, _ := regexp.MatchString(util.PatternOfPassword, str); !ok {
				return errors.New("格式错误")
			}

			return nil
		},
	}

	if password, err = prompt.Run(); err != nil {
		return
	}

	return password, nil
}

func (that *DeveloperProvider) nickname() (nickname string, err error) {

	prompt := promptui.Prompt{
		Label: "昵称",
		Validate: func(str string) error {

			s := strings.TrimSpace(str)

			if s == "" {
				return errors.New("昵称不能为空")
			} else if s != str {
				return errors.New("不能包含空字符")
			}

			return nil
		},
	}

	if nickname, err = prompt.Run(); err != nil {
		return
	}

	return nickname, nil
}
