package rewrite

import (
	"encoding/json"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
)

type SortAction struct {
	actionutils.ParentAction
}

func (this *SortAction) RunPost(params struct {
	WebId          int64
	RewriteRuleIds []int64
}) {
	defer this.CreateLogInfo(codes.HTTPRewriteRule_LogSortRewriteRules, params.WebId)

	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithId(this.AdminContext(), params.WebId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	refsMap := map[int64]*serverconfigs.HTTPRewriteRef{}
	for _, ref := range webConfig.RewriteRefs {
		refsMap[ref.RewriteRuleId] = ref
	}
	newRefs := []*serverconfigs.HTTPRewriteRef{}
	for _, rewriteRuleId := range params.RewriteRuleIds {
		ref, ok := refsMap[rewriteRuleId]
		if ok {
			newRefs = append(newRefs, ref)
		}
	}
	refsJSON, err := json.Marshal(newRefs)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebRewriteRules(this.AdminContext(), &pb.UpdateHTTPWebRewriteRulesRequest{
		HttpWebId:        params.WebId,
		RewriteRulesJSON: refsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
