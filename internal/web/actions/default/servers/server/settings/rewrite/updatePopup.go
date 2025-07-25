package rewrite

import (
	"encoding/json"
	"regexp"

	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/dao"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/shared"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/types"
)

type UpdatePopupAction struct {
	actionutils.ParentAction
}

func (this *UpdatePopupAction) Init() {
}

func (this *UpdatePopupAction) RunGet(params struct {
	WebId         int64
	RewriteRuleId int64
}) {
	this.Data["webId"] = params.WebId

	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithId(this.AdminContext(), params.WebId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	isFound := false
	for _, rewriteRule := range webConfig.RewriteRules {
		if rewriteRule.Id == params.RewriteRuleId {
			this.Data["rewriteRule"] = rewriteRule
			isFound = true
			break
		}
	}

	if !isFound {
		this.WriteString("找不到要修改的重写规则")
		return
	}

	this.Show()
}

func (this *UpdatePopupAction) RunPost(params struct {
	WebId          int64
	RewriteRuleId  int64
	Pattern        string
	Replace        string
	Mode           string
	RedirectStatus int
	ProxyHost      string
	WithQuery      bool
	IsBreak        bool
	IsOn           bool
	CondsJSON      []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo(codes.HTTPRewriteRule_LogUpdateRewriteRule, params.WebId, params.RewriteRuleId)

	params.Must.
		Field("pattern", params.Pattern).
		Require("请输入匹配规则").
		Expect(func() (message string, success bool) {
			_, err := regexp.Compile(params.Pattern)
			if err != nil {
				return "匹配规则错误：" + err.Error(), false
			}
			return "", true
		})

	params.Must.
		Field("replace", params.Replace).
		Require("请输入目标URL")

	// 校验匹配条件
	if len(params.CondsJSON) > 0 {
		conds := &shared.HTTPRequestCondsConfig{}
		err := json.Unmarshal(params.CondsJSON, conds)
		if err != nil {
			this.Fail("匹配条件校验失败：" + err.Error())
		}

		err = conds.Init()
		if err != nil {
			this.Fail("匹配条件校验失败：" + err.Error())
		}
	}

	// 修改
	_, err := this.RPC().HTTPRewriteRuleRPC().UpdateHTTPRewriteRule(this.AdminContext(), &pb.UpdateHTTPRewriteRuleRequest{
		RewriteRuleId:  params.RewriteRuleId,
		Pattern:        params.Pattern,
		Replace:        params.Replace,
		Mode:           params.Mode,
		RedirectStatus: types.Int32(params.RedirectStatus),
		ProxyHost:      params.ProxyHost,
		WithQuery:      params.WithQuery,
		IsBreak:        params.IsBreak,
		IsOn:           params.IsOn,
		CondsJSON:      params.CondsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
