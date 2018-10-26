package ahgenerator

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/a-h/generate"
	"github.com/a-h/generate/jsonschema"
)

func ToStruct(schemaPath string) (gocode string, err error) {
	var fileSchema *jsonschema.Schema
	inflatedJson, err := InflateJson(schemaPath)

	if err != nil {
		return "", fmt.Errorf("inflating '%v' failed: %v", schemaPath, err)
	}

	inflatedBytes, err := json.Marshal(inflatedJson)

	if err != nil {
		return "", fmt.Errorf("marshalling inflated json failed: %v", err)
	}

	if fileSchema, err = jsonschema.Parse(string(inflatedBytes)); err != nil {
		return "", fmt.Errorf("parsing schema '%v' failed: %v", schemaPath, err)
	}

	gen := generate.New(fileSchema)
	structs, aliases, err := gen.CreateTypes()

	if err != nil {
		return "", fmt.Errorf("creating types failed: %v", err)
	}

	sBuilder := &strings.Builder{}
	output(sBuilder, structs, aliases)

	return sBuilder.String(), nil
}
