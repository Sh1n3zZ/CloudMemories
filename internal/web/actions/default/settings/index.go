package settings

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	this.RedirectURL("/settings/server")
}
