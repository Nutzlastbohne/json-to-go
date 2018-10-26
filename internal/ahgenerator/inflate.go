package ahgenerator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const refPathNodeSeparator = "#"

func InflateJson(jsonPath string) ([]byte, error) {
	rawJson, err := loadRawJson(jsonPath)

	if err != nil {
		return nil, err
	}

	inflatedJson, err := inflate(jsonPath, rawJson)

	if err != nil {
		return nil, err
	}

	return json.Marshal(inflatedJson)
}

func inflate(jsonPath string, rawJson map[string]interface{}) (resolvedJson map[string]interface{}, err error) {
	resolvedJson = make(map[string]interface{})
	var refPath string
	var ok bool

	for key, value := range rawJson {
		if key == "$ref" {
			if refPath, ok = value.(string); !ok {
				return nil, fmt.Errorf("invalid $ref value must be string, but is: '%v' (type=%T)", refPath, refPath)
			}

			filePath, nodePath := splitRefPath(refPath)

			if filePath == "" && nodePath != "" { // refPath is self reference
				// leave value as is. file internal reference will be handled by a-h/generator
				resolvedJson[key] = value
				continue
			}

			if nodePath != "" {
				// TODO: load json, go to specified node
				return nil, fmt.Errorf("deep references ('%v') not supported yet", refPath)
			}

			if !filepath.IsAbs(filePath) {
				filePath = filepath.Dir(jsonPath) + "/" + filePath
			}

			refJson, err := loadRawJson(filePath)

			if err != nil {
				return nil, fmt.Errorf("loading referenced json failed: %v", err)
			}

			inflatedRefJson, err := inflate(jsonPath, refJson)

			if err != nil {
				return nil, fmt.Errorf("inflating referenced json '%v' failed: %v", filePath, err)
			}

			// add all found nodes to current node
			for refKey, refVal := range inflatedRefJson {
				resolvedJson[refKey] = refVal
			}
		} else if nestedMap, isMap := value.(map[string]interface{}); isMap {
			inflatedMap, err := inflate(jsonPath, nestedMap)
			if err != nil {
				return nil, err
			}

			resolvedJson[key] = inflatedMap
		} else {
			resolvedJson[key] = value
		}
	}

	return resolvedJson, nil
}

func splitRefPath(refPath string) (filePath, nodePath string) {
	if len(refPath) < 1 {
		return filePath, nodePath
	}

	refParts := strings.Split(refPath, refPathNodeSeparator)
	filePath = refParts[0]

	if len(refParts) > 1 {
		nodePath = refParts[1]
	}

	return filePath, nodePath
}

func isSelfReference(refPath string) bool {
	return refPath[0] == '#'
}

func loadRawJson(filePath string) (rawJson map[string]interface{}, err error) {
	var fileBytes []byte

	if fileBytes, err = ioutil.ReadFile(filePath); err != nil {
		return nil, fmt.Errorf("reading file '%v' failed: %v", filePath, err)
	}

	if err = json.Unmarshal(fileBytes, &rawJson); err != nil {
		return nil, fmt.Errorf("unmarshalling json failed: %v", err)
	}
	return rawJson, nil
}
