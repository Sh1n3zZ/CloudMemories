// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package cluster

import (
	"strconv"

	teaconst "github.com/Sh1n3zZ/CloudMemories/internal/const"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	ClusterId int64
}) {
	if teaconst.IsPlus {
		this.RedirectURL("/clusters/cluster/boards?clusterId=" + strconv.FormatInt(params.ClusterId, 10))
	} else {
		this.RedirectURL("/clusters/cluster/nodes?clusterId=" + strconv.FormatInt(params.ClusterId, 10))
	}
}
