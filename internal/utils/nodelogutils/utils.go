// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus

package nodelogutils

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/langs"
	"github.com/Sh1n3zZ/CMCommon/pkg/langs/codes"
	"github.com/iwind/TeaGo/maps"
)

// FindCommonTags 查找常用的标签
func FindNodeCommonTags(langCode langs.LangCode) []maps.Map {
	return []maps.Map{
		{
			"name": langs.Message(langCode, codes.Log_TagListener),
			"code": "LISTENER",
		},
		{
			"name": langs.Message(langCode, codes.Log_TagWAF),
			"code": "WAF",
		},
		{
			"name": langs.Message(langCode, codes.Log_TagAccessLog),
			"code": "ACCESS_LOG",
		},
	}
}
