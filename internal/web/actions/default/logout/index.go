package logout

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction actions.Action

// 退出登录
func (this *IndexAction) Run(params struct {
	Auth *helpers.UserShouldAuth
}) {
	params.Auth.Logout()
	this.RedirectURL("/")
}
