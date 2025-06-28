// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package updates

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/systemconfigs"
	teaconst "github.com/Sh1n3zZ/CloudMemories/internal/const"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) RunPost(params struct {
	AutoCheck bool
}) {
	defer this.CreateLogInfo(codes.AdminUpdate_LogUpdateCheckSettings)

	// 读取当前设置
	valueResp, err := this.RPC().SysSettingRPC().ReadSysSetting(this.AdminContext(), &pb.ReadSysSettingRequest{Code: systemconfigs.SettingCodeCheckUpdates})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var valueJSON = valueResp.ValueJSON
	var config = systemconfigs.NewCheckUpdatesConfig()
	if len(valueJSON) > 0 {
		err = json.Unmarshal(valueJSON, config)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	config.AutoCheck = params.AutoCheck

	configJSON, err := json.Marshal(config)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 修改设置
	_, err = this.RPC().SysSettingRPC().UpdateSysSetting(this.AdminContext(), &pb.UpdateSysSettingRequest{
		Code:      systemconfigs.SettingCodeCheckUpdates,
		ValueJSON: configJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 重置状态
	if !config.AutoCheck {
		teaconst.NewVersionCode = ""
		teaconst.NewVersionDownloadURL = ""
	}

	this.Success()
}
