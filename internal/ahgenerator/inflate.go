package ahgenerator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func inflateJson(jsonPath string) ([]byte, error) {
	rawJson := loadRawJson(jsonPath)
	inflatedJson := seekDestroy(jsonPath, rawJson)
	return json.Marshal(inflatedJson)
}

func seekDestroy(jsonPath string, rawJson map[string]interface{}) map[string]interface{} {
	var refPath string
	//var err error
	var ok bool
	resolvedJson := make(map[string]interface{})

	for key, value := range rawJson {
		if key == "$ref" {
			if refPath, ok = value.(string); !ok {
				log.Panicf("invalid $ref file: %v", refPath)
			}

			if refPath[0] == '#' {
				log.Printf("file internal reference found - will be handled by a-h gen")
				resolvedJson[key] = value
				continue
			}

			if strings.Contains(refPath, "#") {
				log.Printf("WARN - sub-file references not supported yet.") // more likely a panic?
				continue
			}

			if !filepath.IsAbs(refPath) {
				refPath = filepath.Dir(jsonPath) + "/" + refPath
			}

			refJson := loadRawJson(refPath)
			resolvedRefJson := seekDestroy(jsonPath, refJson)

			// add to current node
			for refKey, refVal := range resolvedRefJson {
				resolvedJson[refKey] = refVal
			}
		} else if nestedMap, isMap := value.(map[string]interface{}); isMap {
			resolvedNestedMap := seekDestroy(jsonPath, nestedMap)
			resolvedJson[key] = resolvedNestedMap
		} else {
			resolvedJson[key] = value
		}
	}

	return resolvedJson
}

func loadRawJson(refPath string) map[string]interface{} {
	var fileBytes []byte
	var err error

	if fileBytes, err = ioutil.ReadFile(refPath); err != nil {
		log.Panicf("reading file failed: %v", err)
	}
	var refJson map[string]interface{}
	if err = json.Unmarshal(fileBytes, &refJson); err != nil {
		log.Panicf("unmarshalling referenced json failed: %v", err)
	}
	return refJson
}
