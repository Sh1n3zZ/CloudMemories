// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package mysql

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/setup/mysql/mysqlinstallers/utils"
)

type InstallLogsAction struct {
	actionutils.ParentAction
}

func (this *InstallLogsAction) RunPost(params struct{}) {
	this.Data["logs"] = utils.SharedLogger.ReadAll()
	this.Success()
}
