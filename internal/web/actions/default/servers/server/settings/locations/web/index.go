package web

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
}

func (this *IndexAction) RunGet(params struct {
	LocationId int64
}) {
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithLocationId(this.AdminContext(), params.LocationId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["rootConfig"] = webConfig.Root

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId    int64
	RootJSON []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo(codes.ServerRoot_LogUpdateRoot, params.WebId)

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWeb(this.AdminContext(), &pb.UpdateHTTPWebRequest{
		HttpWebId: params.WebId,
		RootJSON:  params.RootJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
