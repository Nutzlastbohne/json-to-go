package ahgenerator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const nodePath = "/1st/2nd"

var rawJson = map[string]interface{}{
	"1st": map[string]interface{}{
		"2nd": nestedJson,
	},
}

var nestedJson = map[string]interface{}{
	"name": "lil' jimmy",
	"age":  10,
}

func TestTraversal(t *testing.T) {
	found, err := traverse(rawJson, nodePath)

	if err != nil {
		t.Errorf("traversing path=%v failed: %v", nodePath, err)
	}

	if !cmp.Equal(found, nestedJson) {
		t.Errorf("didn't find expected elements. expected=%+v, got=%+v", nestedJson, found)
	}
}
