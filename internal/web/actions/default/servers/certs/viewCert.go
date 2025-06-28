package certs

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/sslconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type ViewCertAction struct {
	actionutils.ParentAction
}

func (this *ViewCertAction) Init() {
	this.Nav("", "", "")
}

func (this *ViewCertAction) RunGet(params struct {
	CertId int64
}) {
	certResp, err := this.RPC().SSLCertRPC().FindEnabledSSLCertConfig(this.AdminContext(), &pb.FindEnabledSSLCertConfigRequest{SslCertId: params.CertId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if len(certResp.SslCertJSON) == 0 {
		this.NotFound("sslCert", params.CertId)
		return
	}

	certConfig := &sslconfigs.SSLCertConfig{}
	err = json.Unmarshal(certResp.SslCertJSON, certConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	_, _ = this.Write(certConfig.CertData)
}
