package ahgenerator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const refPathNodeSeparator = "#"

func InflateJson(jsonPath string) (map[string]interface{}, error) {
	rawJson, err := loadRawJson(jsonPath)

	if err != nil {
		return nil, err
	}

	filePath, _ := splitRefPath(jsonPath)
	inflatedJson, err := inflate(filePath, rawJson)

	if err != nil {
		return nil, err
	}

	return inflatedJson, nil
}

func inflate(jsonPath string, rawJson map[string]interface{}) (resolvedJson map[string]interface{}, err error) {
	resolvedJson = make(map[string]interface{})
	var refValue string
	var ok bool

	for key, value := range rawJson {
		if key == "$ref" {
			if refValue, ok = value.(string); !ok {
				return nil, fmt.Errorf("invalid $ref value must be string, but is: '%v' (type=%T)", refValue, refValue)
			}

			absJsonPath, err := filepath.Abs(jsonPath)

			if err != nil {
				return nil, fmt.Errorf("getting absolute path of '%v' failed: %v", jsonPath, absJsonPath)
			}

			refPath, nodePath := splitRefPath(refValue)

			if refPath == "" {
				// build path for self-reference
				refValue = jsonPath + "#" + nodePath
			} else  {
				// build path for reference to other file
				rootDir, _ := filepath.Split(absJsonPath)
				refValue = filepath.Join(rootDir, refPath) + "#" + nodePath
			}

			refJson, err := InflateJson(refValue)

			if err != nil {
				return nil, fmt.Errorf("loading referenced json failed: %v", err)
			}

			// add all found nodes to current node
			for refKey, refVal := range refJson {
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

func traverse(rawJson map[string]interface{}, nodePath string) (subNode map[string]interface{}, err error) {
	subNode = make(map[string]interface{})
	nodePath = strings.TrimLeft(nodePath, "/")
	nodes := strings.Split(nodePath, "/")
	currNode := rawJson

	for _, node := range nodes {
		if value, ok := currNode[node].(map[string]interface{}); ok {
			currNode = value
		} else {
			return nil, fmt.Errorf("traversing nodePath='%v' failed. Node '%v' not found", nodePath, node)
		}
	}

	return currNode, nil
}

// loadRawJson loads the specified json structure. This function supports limiting the output to inner nodes (as '$ref' might do).
// E.g. a filePath like "myschema.json#/definitions/person" will only return the json structure of /definitions/person
func loadRawJson(filePath string) (rawJson map[string]interface{}, err error) {
	var fileBytes []byte
	filePath, nodePath := splitRefPath(filePath)

	if fileBytes, err = ioutil.ReadFile(filePath); err != nil {
		return nil, fmt.Errorf("reading file '%v' failed: %v", filePath, err)
	}

	if err = json.Unmarshal(fileBytes, &rawJson); err != nil {
		return nil, fmt.Errorf("unmarshalling json failed: %v", err)
	}

	if nodePath != "" { // limit result to node specified by nodePath
		rawJson, err = traverse(rawJson, nodePath)
		if err != nil {
			return nil, err
		}
	}

	return rawJson, nil
}
