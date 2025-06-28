package log

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/shared"
	"github.com/Sh1n3zZ/CloudMemories/internal/configloaders"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type SettingsAction struct {
	actionutils.ParentAction
}

func (this *SettingsAction) Init() {
	this.Nav("", "", "setting")
}

func (this *SettingsAction) RunGet(params struct{}) {
	config, err := configloaders.LoadLogConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["logConfig"] = config

	this.Show()
}

func (this *SettingsAction) RunPost(params struct {
	CanDelete    bool
	CanClean     bool
	CapacityJSON []byte
	Days         int
	CanChange    bool

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.Log_LogUpdateSettings)

	capacity := &shared.SizeCapacity{}
	err := json.Unmarshal(params.CapacityJSON, capacity)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	config, err := configloaders.LoadLogConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if config.CanChange {
		config.CanDelete = params.CanDelete
		config.CanClean = params.CanClean
		config.Days = params.Days
	}
	config.Capacity = capacity
	config.CanChange = params.CanChange
	err = configloaders.UpdateLogConfig(config)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
