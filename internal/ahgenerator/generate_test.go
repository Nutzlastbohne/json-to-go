package ahgenerator_test

import (
	"testing"
	"time"

	"github.com/nutzlastbohne/json-to-go/internal/ahgenerator"
)

const root = "testdata/"
const standaloneSchema = root + "standalone.json"
const selfRefSchema = root + "selfRef.json"
const otherRefSchema = root + "otherRef.json"
const deepRefSchema = root + "deepRef.json"
const deepNestedRefSchema = root + "deepNestedRef.json"
const nodeSelfRefSchema = root + "nodeSelfRef.json"

func TestStandaloneSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(standaloneSchema); err != nil {
		t.Error(err)
	}
}

func TestSelfReferencingSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(selfRefSchema); err != nil {
		t.Error(err)
	}
}

func TestReferenceToOtherSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(otherRefSchema); err != nil {
		t.Error(err)
	}
}

func TestDeepReferenceToOtherSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(deepRefSchema); err != nil {
		t.Error(err)
	}
}

func TestDeepReferenceToOtherNestedSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(deepNestedRefSchema); err != nil {
		t.Error(err)
	}
}

func TestNodeSelfRefSchema(t *testing.T) {
	var err error
	errChan := make(chan error, 1)

	go func() {
		_, err = ahgenerator.ToStruct(nodeSelfRefSchema)
		errChan <- err
	}()

	select {
	case <-time.After(100 * time.Millisecond):
		t.Errorf("schema with self-referencing node should throw an error to prevent endless-loop")
	case <-errChan:

	}

}
