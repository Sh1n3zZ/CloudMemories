// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package utils_test

import (
	"testing"

	"github.com/Sh1n3zZ/CloudMemories/internal/utils"
)

func TestLookupCNAME(t *testing.T) {
	for _, domain := range []string{"www.yun4s.cn", "example.com"} {
		result, err := utils.LookupCNAME(domain)
		t.Log(domain, "=>", result, err)
	}
}
