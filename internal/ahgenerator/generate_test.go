package ahgenerator_test

import (
	"json-to-struct/internal/ahgenerator"
	"testing"
)

const root = "testdata/"
const personSchemaFile = root + "person.json"
const arrRefSchemaFile = root + "arrayRef.json"
const arrPersonRefSchemaFile = root + "arrayPersonRef.json"

func TestStandaloneSchema(t *testing.T) {
	if err := ahgenerator.ToStruct(personSchemaFile); err != nil {
		t.Error(err)
	}
}

func TestSelfReferencingSchema(t *testing.T) {
	if err := ahgenerator.ToStruct(arrRefSchemaFile); err != nil {
		t.Error(err)
	}
}

func TestOtherReferencingSchema(t *testing.T) {
	if err := ahgenerator.ToStruct(arrPersonRefSchemaFile); err != nil {
		t.Error(err)
	}
}
