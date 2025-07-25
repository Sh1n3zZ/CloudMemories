package grants

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/clusters/grants/grantutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"golang.org/x/crypto/ssh"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "grant", "update")
}

func (this *UpdateAction) RunGet(params struct {
	GrantId int64
}) {
	this.Data["methods"] = grantutils.AllGrantMethods(this.LangCode())

	grantResp, err := this.RPC().NodeGrantRPC().FindEnabledNodeGrant(this.AdminContext(), &pb.FindEnabledNodeGrantRequest{NodeGrantId: params.GrantId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if grantResp.NodeGrant == nil {
		this.WriteString("can not find the grant")
		return
	}

	// TODO 处理节点专用的认证

	grant := grantResp.NodeGrant
	this.Data["grant"] = maps.Map{
		"id":          grant.Id,
		"name":        grant.Name,
		"method":      grant.Method,
		"methodName":  grantutils.FindGrantMethodName(grant.Method, this.LangCode()),
		"username":    grant.Username,
		"password":    grant.Password,
		"privateKey":  grant.PrivateKey,
		"passphrase":  grant.Passphrase,
		"description": grant.Description,
		"su":          grant.Su,
	}

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	GrantId     int64
	Name        string
	Method      string
	Username    string
	Password    string
	PrivateKey  string
	Passphrase  string
	Description string
	Su          bool

	Must *actions.Must
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.NodeGrant_LogUpdateSSHGrant, params.GrantId)

	params.Must.
		Field("name", params.Name).
		Require("请输入名称")

	switch params.Method {
	case "user":
		if len(params.Username) == 0 {
			this.FailField("username", "请输入SSH登录用户名")
		}
	case "privateKey":
		if len(params.Username) == 0 {
			this.FailField("username", "请输入SSH登录用户名")
		}
		if len(params.PrivateKey) == 0 {
			this.FailField("privateKey", "请输入RSA私钥")
		}

		// 验证私钥
		var err error
		if len(params.Passphrase) > 0 {
			_, err = ssh.ParsePrivateKeyWithPassphrase([]byte(params.PrivateKey), []byte(params.Passphrase))
		} else {
			_, err = ssh.ParsePrivateKey([]byte(params.PrivateKey))
		}
		if err != nil {
			this.Fail("私钥验证失败，请检查格式：" + err.Error())
			return
		}
	default:
		this.Fail("请选择正确的认证方式")
	}

	// TODO 检查grantId是否存在

	_, err := this.RPC().NodeGrantRPC().UpdateNodeGrant(this.AdminContext(), &pb.UpdateNodeGrantRequest{
		NodeGrantId: params.GrantId,
		Name:        params.Name,
		Method:      params.Method,
		Username:    params.Username,
		Password:    params.Password,
		PrivateKey:  params.PrivateKey,
		Passphrase:  params.Passphrase,
		Description: params.Description,
		Su:          params.Su,
		NodeId:      0,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
