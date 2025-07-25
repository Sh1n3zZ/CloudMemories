// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package ui

import (
	"strings"

	"github.com/Sh1n3zZ/CMCommon/pkg/rpc/pb"
	"github.com/Sh1n3zZ/CMCommon/pkg/serverconfigs/regionconfigs"
	"github.com/Sh1n3zZ/CloudMemories/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type CountryOptionsAction struct {
	actionutils.ParentAction
}

func (this *CountryOptionsAction) RunPost(params struct{}) {
	countriesResp, err := this.RPC().RegionCountryRPC().FindAllRegionCountries(this.AdminContext(), &pb.FindAllRegionCountriesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var countryMaps = []maps.Map{}
	for _, country := range countriesResp.RegionCountries {
		if lists.ContainsInt64(regionconfigs.FindAllGreaterChinaSubRegionIds(), country.Id) {
			continue
		}

		if country.Codes == nil {
			country.Codes = []string{}
		}

		var letter = ""
		if len(country.Pinyin) > 0 && len(country.Pinyin[0]) > 0 {
			letter = strings.ToUpper(country.Pinyin[0][:1])
		}

		countryMaps = append(countryMaps, maps.Map{
			"id":       country.Id,
			"name":     country.Name,
			"fullname": letter + " " + country.Name,
			"codes":    country.Codes,
		})
	}
	this.Data["countries"] = countryMaps

	this.Success()
}
