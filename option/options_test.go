package option

import (
	"fmt"
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

func TestSetPeriod(t *testing.T) {
	testString := "Hello"
	result := SetPeriod(testString)
	if result != "Hello." {
		t.Errorf("Failed to set fullstop on %s", testString)
	}
}

func TestSetPeriodIfPresent(t *testing.T) {
	testString := "Hello."
	result := SetPeriod(testString)
	if result != "Hello." {
		t.Errorf("Failed to set fullstop on %s", testString)
	}
}

func TestRemovePeriod(t *testing.T) {
	testString := "Hello."
	result := RemovePeriod(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove fullstop on %s", testString)
	}
}

//TODO make these tests simpler

func TestTrimBeforeWithOneWord(t *testing.T) {
	testString := "The. The quick brown fox jumps over the lazy dog"
	result := TrimBefore(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove prefix of %s", testString)
	}
}

func TestTrimAfterWithAPeriod(t *testing.T) {
	testString := ". The quick brown fox jumps over the lazy dog"
	result := TrimBefore(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove prefix of %s", testString)
	}
}

func TestTrimAfterWithOneWord(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. dog"
	result := TrimAfter(testString)
	fmt.Println(testString)
	if result != "The quick brown fox jumps over the lazy" {
		t.Errorf("Failed to remove postfix of %s", testString)
	}
}

func TestTrimAfterWithALetter(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. d"
	result := TrimAfter(testString)
	if result != "The quick brown fox jumps over the lazy" {
		t.Errorf("Failed to remove postfix of %s", testString)
	}
}
