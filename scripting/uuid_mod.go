package scripting

import (
	"fmt"

	"github.com/d5/tengo/objects"
	"github.com/google/uuid"
)

var moduleMap objects.Object = &objects.ImmutableMap{
	Value: map[string]objects.Object{
		"new": &objects.UserFunction{
			Name:  "new_uuid",
			Value: newUUID,
		},
	},
}

type UUIDScriptModule struct{}

func newUUID(args ...objects.Object) (objects.Object, error) {
	if len(args) > 0 {
		return nil, newArgumentCountError(0, len(args))
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &UUIDScriptObject{UUID: id}, nil
}

func (UUIDScriptModule) Import(moduleName string) (interface{}, error) {
	switch moduleName {
	case "uuid":
		return moduleMap, nil
	default:
		return nil, fmt.Errorf("undefined module %q", moduleName)
	}
}
