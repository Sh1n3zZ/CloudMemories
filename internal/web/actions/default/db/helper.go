package db

import (
	"net/http"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo/actions"
)

type Helper struct {
	helpers.LangHelper
}

func (this *Helper) BeforeAction(action *actions.ActionObject) {
	if action.Request.Method != http.MethodGet {
		return
	}

	action.Data["teaMenu"] = "db"

	var selectedTabbar = action.Data["mainTab"]

	var tabbar = actionutils.NewTabbar()
	tabbar.Add(this.Lang(action, codes.DBNode_TabNodes), "", "/db", "", selectedTabbar == "db")
	actionutils.SetTabbar(action, tabbar)
}
