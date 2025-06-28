package unix

import (
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("unix")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	// TODO

	this.Show()
}
