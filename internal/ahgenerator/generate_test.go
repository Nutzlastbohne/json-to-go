package ahgenerator_test

import (
	"testing"

	"github.com/nutzlastbohne/json-to-go/internal/ahgenerator"
)

const root = "testdata/"
const standaloneSchema = root + "standalone.json"
const selfRefSchema = root + "selfRef.json"
const otherRefSchema = root + "otherRef.json"
const deepRefSchemaFile = root + "deepRef.json"

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
	if _, err := ahgenerator.ToStruct(deepRefSchemaFile); err != nil {
		t.Error(err)
	}
}