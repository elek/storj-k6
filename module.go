package k6

import (
	"go.k6.io/k6/js/modules"
	_ "storj.io/storj/shared/dbutil/cockroachutil"
)

func init() {
	modules.Register("k6/x/stbb", New())
}

type (
	RootModule struct{}

	ModuleInstance struct {
		vu modules.VU
	}
)

var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu: vu,
	}
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: MetainfoTest,
		Named: map[string]interface{}{
			"MetainfoTest": MetainfoTest,
			"UplinkTest":   UplinkTest,
		},
	}
}
