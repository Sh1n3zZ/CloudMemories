// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package charts

import (
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/metrics/charts/chartutils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/default/servers/metrics/metricutils"
	"github.com/iwind/TeaGo/maps"
)

type ChartAction struct {
	actionutils.ParentAction
}

func (this *ChartAction) Init() {
	this.Nav("", "", "chart,chartIndex")
}

func (this *ChartAction) RunGet(params struct {
	ChartId int64
}) {
	chart, err := chartutils.InitChart(this.Parent(), params.ChartId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = metricutils.InitItem(this.Parent(), chart.MetricItem.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	var itemMap = this.Data["item"].(maps.Map)
	itemMap["valueTypeName"] = serverconfigs.FindMetricValueName(itemMap.GetString("category"), itemMap.GetString("value"))

	this.Show()
}
