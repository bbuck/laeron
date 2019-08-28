package main

import (
	"laeron/scripting"

	"github.com/d5/tengo/objects"
	"github.com/d5/tengo/script"
	"github.com/d5/tengo/stdlib"
)

func main() {
	s := script.New([]byte(`
		fmt := import("fmt")
		uuid := import("uuid")

		fmt.println(uuid.new())
	`))
	mm := objects.NewModuleMap()
	mm.Add("uuid", scripting.UUIDScriptModule{})
	mm.AddMap(stdlib.GetModuleMap("fmt"))
	s.SetImports(mm)
	if _, err := s.Run(); err != nil {
		panic(err)
	}
}
