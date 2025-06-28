package configloaders

import "github.com/Sh1n3zZ/CMCommon/pkg/systemconfigs"

type AdminModuleList struct {
	IsSuper  bool
	Modules  []*systemconfigs.AdminModule
	Fullname string
	Theme    string
	Lang     string
}

func (this *AdminModuleList) Allow(module string) bool {
	if this.IsSuper {
		return true
	}
	for _, m := range this.Modules {
		if m.Code == module {
			return true
		}
	}
	return false
}
