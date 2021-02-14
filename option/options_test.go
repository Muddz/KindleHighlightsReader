package option

import (
	"testing"
)

//------ Test for setting quotations

func TestSetQuotationsIfNotPresent(t *testing.T) {
	testString := "Hello"
	result := SetQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on: %s", testString)
	}
}

func TestSetQuotationsIfPresent(t *testing.T) {
	testString := "“Hello”"
	result := SetQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on: %s", testString)
	}
}

func TestSetQuotationsIfSinglesPresent(t *testing.T) {
	testString := "‘Hello’"
	result := SetQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on: %s", testString)
	}
}

func TestSetQuotationsIfPrimesPresent(t *testing.T) {
	testString := "\"Hello\""
	result := SetQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on: %s", testString)
	}
}

func TestSetQuotationsIfSinglePrimesPresent(t *testing.T) {
	testString := "'Hello'"
	result := SetQuotations(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to set double quotations on: %s", testString)
	}
}

//------ Test for removing quotations

func TestRemoveQuotations(t *testing.T) {
	testString := "“Hello”"
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove double quotations on: %s", testString)
	}
}

func TestRemoveSingleQuotations(t *testing.T) {
	testString := "‘Hello’"
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove single quotations on: %s", testString)
	}
}

func TestRemovePrimeQuotations(t *testing.T) {
	testString := "\"Hello\""
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove double quotations on: %s", testString)
	}
}

func TestRemoveSinglePrimeQuotations(t *testing.T) {
	testString := "'Hello'"
	result := RemoveQuotations(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove single quotations on: %s", testString)
	}
}

//------ Test for setting periods

func TestSetPeriodWithNoQuotations(t *testing.T) {
	testString := "Hello"
	result := SetPeriod(testString)
	if result != "Hello." {
		t.Errorf("Failed to set period on: %s", testString)
	}
}

func TestSetPeriodWithQuotations(t *testing.T) {
	testString := "“Hello”"
	result := SetPeriod(testString)
	if result != "“Hello.”" {
		t.Errorf("Failed to set period on: %s", testString)
	}
}

func TestSetPeriodWithSingleQuotations(t *testing.T) {
	testString := "‘Hello’"
	result := SetPeriod(testString)
	if result != "‘Hello.’" {
		t.Errorf("Failed to set period on: %s", testString)
	}
}

func TestSetPeriodWithPrimeQuotations(t *testing.T) {
	testString := "\"Hello\""
	result := SetPeriod(testString)
	if result != "\"Hello.\"" {
		t.Errorf("Failed to set period on: %s", testString)
	}
}

func TestSetPeriodWithSinglePrimeQuotations(t *testing.T) {
	testString := "'Hello'"
	result := SetPeriod(testString)
	if result != "'Hello.'" {
		t.Errorf("Failed to set period on: %s", testString)
	}
}

//------ Test for removing periods

func TestRemovePeriodWithNoQuotations(t *testing.T) {
	testString := "Hello."
	result := RemovePeriod(testString)
	if result != "Hello" {
		t.Errorf("Failed to remove period on: %s", testString)
	}
}

func TestRemovePeriodWithQuotations(t *testing.T) {
	testString := "“Hello.”"
	result := RemovePeriod(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to remove period on: %s", testString)
	}
}

func TestRemovePeriodWithSingleQuotations(t *testing.T) {
	testString := "‘Hello.’"
	result := RemovePeriod(testString)
	if result != "‘Hello’" {
		t.Errorf("Failed to remove period on: %s", testString)
	}
}

func TestRemovePeriodWithPrimeQuotations(t *testing.T) {
	testString := "\"Hello.\""
	result := RemovePeriod(testString)
	if result != "\"Hello\"" {
		t.Errorf("Failed to remove period on: %s", testString)
	}
}

func TestRemovePeriodWithSinglePrimeQuotations(t *testing.T) {
	testString := "'Hello.'"
	result := RemovePeriod(testString)
	if result != "'Hello'" {
		t.Errorf("Failed to remove period on: %s", testString)
	}
}

//------ Test for trimming for redundant chars/words

func TestTrimBefore(t *testing.T) {
	testString := "The. The quick brown fox jumps over the lazy dog"
	result := TrimBefore(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to trim-before of: %s", testString)
	}
}

func TestTrimAfterWithAPeriod(t *testing.T) {
	testString := ". The quick brown fox jumps over the lazy dog"
	result := TrimBefore(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to trim-before of: %s", testString)
	}
}

func TestTrimAfterWithOneWord(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. dog"
	result := TrimAfter(testString)
	if result != "The quick brown fox jumps over the lazy." {
		t.Errorf("Failed to trim-after of: %s", testString)
	}
}

//------ Test for capitalizing first letter in texts

func TestCapitalize(t *testing.T) {
	testString := "hello"
	result := Capitalize(testString)
	if result != "Hello" {
		t.Errorf("Failed to capitalize: %s", testString)
	}
}

func TestCapitalizeWithQuotations(t *testing.T) {
	testString := "“hello”"
	result := Capitalize(testString)
	if result != "“Hello”" {
		t.Errorf("Failed to capitalize: %s", testString)
	}
}

func TestCapitalizeWithSingleQuotations(t *testing.T) {
	testString := "‘hello’"
	result := Capitalize(testString)
	if result != "‘Hello’" {
		t.Errorf("Failed to capitalize: %s", testString)
	}
}

func TestCapitalizeWithPrimeQuotations(t *testing.T) {
	testString := "\"hello\""
	result := Capitalize(testString)
	if result != "\"Hello\"" {
		t.Errorf("Failed to capitalize: %s", testString)
	}
}

func TestCapitalizeWithSinglePrimeQuotations(t *testing.T) {
	testString := "'hello'"
	result := Capitalize(testString)
	if result != "'Hello'" {
		t.Errorf("Failed to capitalize: %s", testString)
	}
}
