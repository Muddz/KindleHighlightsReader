package punctuations

import (
	"testing"
)

func TestSetSingleQuotations(t *testing.T) {
	testString := "Hello"
	testString = SetSingleQuotations(testString)
	firstChar := string(testString[0])
	lastChar := string(testString[len(testString)-1])
	if firstChar != "'" && lastChar != "'" {
		t.Errorf("Failed to set single quotations on %s", testString)
	}
}

func TestSetSingleQuotationsIfDoubleIsPresent(t *testing.T) {
	testString := "\"Hello\""
	testString = SetSingleQuotations(testString)
	firstChar := string(testString[0])
	lastChar := string(testString[len(testString)-1])
	if firstChar != "'" && lastChar != "'" {
		t.Errorf("Failed to set single quotations on %s", testString)
	}
}

func TestSetDoubleQuotations(t *testing.T) {
	testString := `Hello`
	testString = SetDoubleQuotations(testString)
	firstChar := string(testString[0])
	lastChar := string(testString[len(testString)-1])
	if firstChar != "\"" && lastChar != "\"" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestSetDoubleQuotationsIfSingleIsPresent(t *testing.T) {
	testString := `'Hello'`
	testString = SetDoubleQuotations(testString)
	firstChar := string(testString[0])
	lastChar := string(testString[len(testString)-1])
	if firstChar != "\"" && lastChar != "\"" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestRemoveSingleQuotations(t *testing.T) {
	testString := "'Hello'"
	testString = RemoveQuotations(testString)
	firstChar := string(testString[0])
	lastChar := string(testString[len(testString)-1])
	if firstChar == "'" && lastChar == "'" {
		t.Errorf("Failed to remove single quatations from %s", testString)
	}
}

func TestRemoveDoubleQuotations(t *testing.T) {
	testString := "\"Hello\""
	testString = RemoveQuotations(testString)
	firstChar := string(testString[0])
	lastChar := string(testString[len(testString)-1])
	if firstChar == "\"" && lastChar == "\"" {
		t.Errorf("Failed to remove single quatations from %s", testString)
	}
}

func TestSetFullStop(t *testing.T) {
	testString := "Hello"
	testString = SetFullStop(testString)
	lastChar := string(testString[len(testString)-1])
	if lastChar != "." {
		t.Errorf("Failed to set fullstop/period on %s", testString)
	}
}

func TestSetFullStopWitIfPresent(t *testing.T) {
	testString := "Hello."
	testString = SetFullStop(testString)
	lastChar := string(testString[len(testString)-1])
	secondLastChar := string(testString[len(testString)-2])
	if secondLastChar == "." && lastChar == "." {
		t.Errorf("Error! Inserted a fullstop/period when %s already had one", testString)
	}
}

func TestRemoveFullStop(t *testing.T) {
	testString := "Hello."
	testString = RemoveFullStop(testString)
	lastChar := string(testString[len(testString)-1])
	if lastChar == "." {
		t.Errorf("Failed to remove fullstop/periode from %s", testString)
	}
}
