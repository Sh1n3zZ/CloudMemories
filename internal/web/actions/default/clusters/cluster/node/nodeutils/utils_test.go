// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeutils_test

import (
	"testing"

	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/clusters/cluster/node/nodeutils"
	_ "github.com/iwind/TeaGo/bootstrap"
)

func TestInstallLocalNode(t *testing.T) {
	err := nodeutils.InstallLocalNode()
	if err != nil {
		t.Fatal(err)
	}
}
