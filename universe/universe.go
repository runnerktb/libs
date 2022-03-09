package universe

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed universe.json
	data     []byte
	universe = make(map[string]string)
)

func init() {
	err := json.Unmarshal(data, &universe)
	if err != nil {
		panic(err)
	}
}

func Version() string {
	return universe["$id"]
}

func GetModule(module string) (string, bool) {
	id, ok := universe[module]
	return id, ok
}
