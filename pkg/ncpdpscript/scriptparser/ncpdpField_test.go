package scriptparser

import (
	"log"
	"testing"
)

var fieldTestData = "CC123456"
var ncpdpTestField, initError = newField([]byte(fieldTestData))

func TestFieldId(t *testing.T) {
	if initError != nil {
		log.Panic(initError)
	}
	want := "CC"
	got := ncpdpTestField.fieldId
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestField_String(t *testing.T) {
	if initError != nil {
		log.Panic(initError)
	}
	want := FieldSeparator + "CC123456"
	got, err := ncpdpTestField.String()
	if err != nil {
		log.Panic(err)
	}
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
