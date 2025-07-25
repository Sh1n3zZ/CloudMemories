//go:build !plus

package settingutils

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	teaconst "github.com/Sh1n3zZ/CloudMemories/internal/const"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo/actions"
)

type AdvancedHelper struct {
	helpers.LangHelper

	tab string
}

func NewAdvancedHelper(tab string) *AdvancedHelper {
	return &AdvancedHelper{
		tab: tab,
	}
}

func (this *AdvancedHelper) BeforeAction(actionPtr actions.ActionWrapper) (goNext bool) {
	goNext = true

	var action = actionPtr.Object()

	// 左侧菜单
	action.Data["teaMenu"] = "settings"
	action.Data["teaSubMenu"] = "advanced"

	// 标签栏
	var tabbar = actionutils.NewTabbar()
	var session = action.Session()
	var adminId = session.GetInt64(teaconst.SessionAdminId)
	if configloaders.AllowModule(adminId, configloaders.AdminModuleCodeSetting) {
		tabbar.Add(this.Lang(actionPtr, codes.AdminSetting_TabTransfer), "", "/settings/database", "", this.tab == "database")
		tabbar.Add(this.Lang(actionPtr, codes.AdminSetting_TabAPINodes), "", "/settings/api", "", this.tab == "apiNodes")
		tabbar.Add(this.Lang(actionPtr, codes.AdminSetting_TabAccessLogDatabases), "", "/db", "", this.tab == "dbNodes")
		tabbar.Add(this.Lang(actionPtr, codes.AdminSetting_TabTransfer), "", "/settings/transfer", "", this.tab == "transfer")

		//tabbar.Add(codes.AdminSettingsTabBackup, "", "/settings/backup", "", this.tab == "backup")
	}
	actionutils.SetTabbar(actionPtr, tabbar)

	return
}
