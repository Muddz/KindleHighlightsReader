package option

import (
	"testing"
)

func TestSetDoubleQuotationsIfNotPresent(t *testing.T) {
	testString := "Hello"
	result := SetDoubleQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestSetDoubleQuotationsIfPresent(t *testing.T) {
	testString := "“Hello”"
	result := SetDoubleQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestSetDoubleQuotationsIfPrimePresent(t *testing.T) {
	testString := "\"Hello\""
	result := SetDoubleQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on %s", testString)
	}
}

func TestSetDoubleQuotationsIfSinglePresent(t *testing.T) {
	testString := "'Hello'"
	result := SetDoubleQuotations(testString)
	if result != "“Hello”" {
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

func TestRemoveDoublePrimeQuotations(t *testing.T) {
	testString := "\"Hello\""
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove double quotations on %s", testString)
	}
}

func TestRemoveDoubleQuotations(t *testing.T) {
	testString := "“Hello”"
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove double quotations on %s", testString)
	}
}

func TestSetPeriod(t *testing.T) {
	testString := "Hello"
	result := SetPeriod(testString)
	if result != "Hello." {
		t.Errorf("Failed to set period on '%s'", testString)
	}
}

func TestSetPeriodIfPresent(t *testing.T) {
	testString := "Hello."
	result := SetPeriod(testString)
	if result != "Hello." {
		t.Errorf("Failed to set period on '%s'", testString)
	}
}

func TestRemovePeriod(t *testing.T) {
	testString := "Hello."
	result := RemovePeriod(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove period on '%s'", testString)
	}
}

func TestTrimBeforeWithOneWord(t *testing.T) {
	testString := "The. The quick brown fox jumps over the lazy dog"
	result := TrimBefore(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to trim-before of '%s'", testString)
	}
}

func TestTrimAfterWithAPeriod(t *testing.T) {
	testString := ". The quick brown fox jumps over the lazy dog"
	result := TrimBefore(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to trim-before of '%s'", testString)
	}
}

func TestTrimAfterWithOneWord(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. dog"
	result := TrimAfter(testString)
	if result != "The quick brown fox jumps over the lazy." {
		t.Errorf("Failed to trim-after of '%s'", testString)
	}
}
