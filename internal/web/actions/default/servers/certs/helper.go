package certs

import (
	"net/http"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/helpers"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
)

type Helper struct {
	helpers.LangHelper
}

func NewHelper() *Helper {
	return &Helper{}
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) {
	var action = actionPtr.Object()
	if action.Request.Method != http.MethodGet {
		return
	}

	action.Data["teaMenu"] = "servers"

	var countOCSP int64 = 0
	parentAction, ok := actionPtr.(actionutils.ActionInterface)
	if ok {
		countOCSPResp, err := parentAction.RPC().SSLCertRPC().CountAllSSLCertsWithOCSPError(parentAction.AdminContext(), &pb.CountAllSSLCertsWithOCSPErrorRequest{})
		if err == nil {
			countOCSP = countOCSPResp.Count
		}
	}

	var ocspMenuName = this.Lang(actionPtr, codes.SSLCert_MenuOCSP)
	if countOCSP > 0 {
		ocspMenuName += "(" + types.String(countOCSP) + ")"
	}

	var menu = []maps.Map{
		{
			"name":     this.Lang(actionPtr, codes.SSLCert_MenuCerts),
			"url":      "/servers/certs",
			"isActive": action.Data.GetString("leftMenuItem") == "cert",
		},
		{
			"name":     this.Lang(actionPtr, codes.SSLCert_MenuApply),
			"url":      "/servers/certs/acme",
			"isActive": action.Data.GetString("leftMenuItem") == "acme",
		},
		{
			"name":     ocspMenuName,
			"url":      "/servers/certs/ocsp",
			"isActive": action.Data.GetString("leftMenuItem") == "ocsp",
		},
	}
	action.Data["leftMenuItems"] = menu
}
