// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package ui

import "github.com/Sh1n3zZ/CMCommon/pkg/systemconfigs"

func (this *IndexAction) filterConfig(config *systemconfigs.AdminUIConfig) {
	this.Data["supportModuleCDN"] = true
	this.Data["supportModuleNS"] = true
	this.Data["nsIsVisible"] = false
}
