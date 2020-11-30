package punctuations

import (
	"testing"
)

func TestSetSingleQuotations(t *testing.T) {
	testString := "Hello"
	result := SetSingleQuotations(testString)
	if result != "'Hello'" {
		t.Errorf("Failed to set single quotations on %s", testString)
	}
}

func TestSetSingleQuotationsIfDoubleIsPresent(t *testing.T) {
	testString := "\"Hello\""
	result := SetSingleQuotations(testString)
	if result != "'Hello'" {
		t.Errorf("Failed to set single quotations on %s", testString)
	}
}

func TestSetDoubleQuotations(t *testing.T) {
	testString := `Hello`
	result := SetDoubleQuotations(testString)
	if result != "\"Hello\"" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestSetDoubleQuotationsIfSingleIsPresent(t *testing.T) {
	testString := `'Hello'`
	result := SetDoubleQuotations(testString)
	if result != "\"Hello\"" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestRemoveSingleQuotations(t *testing.T) {
	testString := "'Hello'"
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove single quotations on %s", testString)
	}
}

func TestRemoveDoubleQuotations(t *testing.T) {
	testString := "\"Hello\""
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove double quotations on %s", testString)
	}
}

func TestSetFullStop(t *testing.T) {
	testString := "Hello"
	result := SetFullStop(testString)
	if result != "Hello." {
		t.Errorf("Failed to set fullstop on %s", testString)
	}
}

func TestSetFullStopWitIfPresent(t *testing.T) {
	testString := "Hello."
	result := SetFullStop(testString)
	if result != "Hello." {
		t.Errorf("Failed to set fullstop on %s", testString)
	}
}

func TestRemoveFullStop(t *testing.T) {
	testString := "Hello."
	result := RemoveFullStop(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove fullstop on %s", testString)
	}
}
