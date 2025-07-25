package dns

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("dns")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	dnsInfoResp, err := this.RPC().ServerRPC().FindEnabledServerDNS(this.AdminContext(), &pb.FindEnabledServerDNSRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["dnsName"] = dnsInfoResp.DnsName
	if dnsInfoResp.Domain != nil {
		this.Data["dnsDomain"] = dnsInfoResp.Domain.Name
	} else {
		this.Data["dnsDomain"] = ""
	}
	this.Data["supportCNAME"] = dnsInfoResp.SupportCNAME

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId     int64
	SupportCNAME bool

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.ServerDNS_LogUpdateDNSSettings, params.ServerId)

	_, err := this.RPC().ServerRPC().UpdateServerDNS(this.AdminContext(), &pb.UpdateServerDNSRequest{
		ServerId:     params.ServerId,
		SupportCNAME: params.SupportCNAME,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
