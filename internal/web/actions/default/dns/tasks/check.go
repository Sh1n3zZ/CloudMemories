package tasks

import (
	"time"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type CheckAction struct {
	actionutils.ParentAction
}

func (this *CheckAction) RunPost(params struct {
	IsDoing   bool
	HasError  bool
	IsUpdated bool
}) {
	var isStream = this.Request.ProtoMajor >= 2
	this.Data["isStream"] = isStream

	var maxTries = 10
	for i := 0; i < maxTries; i++ {
		resp, err := this.RPC().DNSTaskRPC().ExistsDNSTasks(this.AdminContext(), &pb.ExistsDNSTasksRequest{})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		// 如果没有数据变化，继续查询
		if i < maxTries-1 && params.IsUpdated && resp.ExistTasks == params.IsDoing && resp.ExistError == params.HasError && isStream {
			time.Sleep(3 * time.Second)
			continue
		}

		this.Data["isDoing"] = resp.ExistTasks
		this.Data["hasError"] = resp.ExistError
		break
	}

	this.Success()
}
