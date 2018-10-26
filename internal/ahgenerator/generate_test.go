package ahgenerator_test

import (
	"testing"

	"github.com/nutzlastbohne/json-to-go/internal/ahgenerator"
)

const root = "testdata/"
const personSchemaFile = root + "person.json"
const arrRefSchemaFile = root + "arrayRef.json"
const arrPersonRefSchemaFile = root + "arrayPersonRef.json"
const deepRefSchemaFile = root + "deepReference.json"

func TestStandaloneSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(personSchemaFile); err != nil {
		t.Error(err)
	}
}

func TestSelfReferencingSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(arrRefSchemaFile); err != nil {
		t.Error(err)
	}
}

func TestReferenceToOtherSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(arrPersonRefSchemaFile); err != nil {
		t.Error(err)
	}
}

func TestDeepReferenceToOtherSchema(t *testing.T) {
	if _, err := ahgenerator.ToStruct(deepRefSchemaFile); err != nil {
		t.Error(err)
	}
}