// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ui

import (
	"encoding/json"
	"os"

	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/Tea"
)

type HideTipAction struct {
	actionutils.ParentAction
}

func (this *HideTipAction) RunPost(params struct {
	Code string
}) {
	tipKeyLocker.Lock()
	tipKeyMap[params.Code] = true
	tipKeyLocker.Unlock()

	// 保存到文件
	tipJSON, err := json.Marshal(tipKeyMap)
	if err == nil {
		_ = os.WriteFile(Tea.ConfigFile(tipConfigFile), tipJSON, 0666)
	}

	this.Success()
}
