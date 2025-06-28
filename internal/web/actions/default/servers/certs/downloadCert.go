package certs

import (
	"encoding/json"
	"strconv"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/sslconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type DownloadCertAction struct {
	actionutils.ParentAction
}

func (this *DownloadCertAction) Init() {
	this.Nav("", "", "")
}

func (this *DownloadCertAction) RunGet(params struct {
	CertId int64
}) {
	defer this.CreateLogInfo(codes.SSLCert_LogDownloadSSLCert, params.CertId)

	certResp, err := this.RPC().SSLCertRPC().FindEnabledSSLCertConfig(this.AdminContext(), &pb.FindEnabledSSLCertConfigRequest{SslCertId: params.CertId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	certConfig := &sslconfigs.SSLCertConfig{}
	err = json.Unmarshal(certResp.SslCertJSON, certConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.AddHeader("Content-Disposition", "attachment; filename=\"cert-"+strconv.FormatInt(params.CertId, 10)+".pem\";")
	_, _ = this.Write(certConfig.CertData)
}
