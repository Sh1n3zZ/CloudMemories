// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package ipadmin

import (
	"strings"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CloudMemories/internal/utils"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type SelectCountriesPopupAction struct {
	actionutils.ParentAction
}

func (this *SelectCountriesPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *SelectCountriesPopupAction) RunGet(params struct {
	Type               string
	SelectedCountryIds string
}) {
	this.Data["type"] = params.Type

	var selectedCountryIds = utils.SplitNumbers(params.SelectedCountryIds)

	countriesResp, err := this.RPC().RegionCountryRPC().FindAllRegionCountries(this.AdminContext(), &pb.FindAllRegionCountriesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// special regions

	var countryMaps = []maps.Map{}
	for _, country := range countriesResp.RegionCountries {
		countryMaps = append(countryMaps, maps.Map{
			"id":        country.Id,
			"name":      country.DisplayName,
			"letter":    strings.ToUpper(string(country.Pinyin[0][0])),
			"pinyin":    country.Pinyin,
			"codes":     country.Codes,
			"isCommon":  country.IsCommon,
			"isChecked": lists.ContainsInt64(selectedCountryIds, country.Id),
		})
	}
	this.Data["countries"] = countryMaps

	this.Show()
}
