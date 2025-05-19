package tengomodules

import (
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/skateboard/tengomodules/http"
)

var Modules = map[string]map[string]tengo.Object{
	"http": http.Module,
}

func GetModule(names ...string) *tengo.ModuleMap {
	modules := tengo.NewModuleMap()
	for _, name := range names {
		if mod := Modules[name]; mod != nil {
			modules.AddBuiltinModule(name, mod)
		}
	}

	return modules
}

func LoadAllModules(includeStdlib bool) *tengo.ModuleMap {
	modules := tengo.NewModuleMap()

	if includeStdlib {
		for name, mod := range stdlib.BuiltinModules {
			modules.AddBuiltinModule(name, mod)
		}

		for name, mod := range stdlib.SourceModules {
			modules.AddSourceModule(name, []byte(mod))
		}
	}

	for name, mod := range Modules {
		modules.AddBuiltinModule(name, mod)
	}

	return modules
}
