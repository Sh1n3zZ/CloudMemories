package web

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

// 创建首页文件
type CreateIndexAction struct {
	actionutils.ParentAction
}

func (this *CreateIndexAction) Init() {
	this.Nav("", "", "")
}

func (this *CreateIndexAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreateIndexAction) RunPost(params struct {
	Index string

	Must *actions.Must
}) {
	params.Must.
		Field("index", params.Index).
		Require("首页文件不能为空")

	this.Data["index"] = params.Index
	this.Success()
}
