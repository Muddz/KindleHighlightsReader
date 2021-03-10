package option

import (
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestSetQuotations(t *testing.T) {
	expected := "“Hello”"
	c1 := TestCase{input: "Hello", expected: expected}
	c2 := TestCase{input: "“Hello”", expected: expected}
	c3 := TestCase{input: "“Hello", expected: expected}
	c4 := TestCase{input: "Hello”", expected: expected}
	c5 := TestCase{input: "\"Hello\"", expected: expected}
	cases := []TestCase{c1, c2, c3, c4, c5}
	for i, v := range cases {
		actual := SetQuotations(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: %s expected: %s", i+1, actual, v.expected)
		}
	}
}

func TestRemoveQuotations(t *testing.T) {
	expected := "Hello"
	c1 := TestCase{input: "Hello", expected: expected}
	c2 := TestCase{input: "“Hello”", expected: expected}
	c3 := TestCase{input: "“Hello", expected: expected}
	c4 := TestCase{input: "Hello”", expected: expected}
	c5 := TestCase{input: "\"Hello\"", expected: expected}
	cases := []TestCase{c1, c2, c3, c4, c5}
	for i, v := range cases {
		actual := RemoveQuotations(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: %s expected: %s", i+1, actual, v.expected)
		}
	}
}

func TestSetPeriod(t *testing.T) {
	c1 := TestCase{input: "Hello", expected: "Hello."}
	c2 := TestCase{input: "“Hello”", expected: "“Hello.”"}
	c3 := TestCase{input: "Hello.", expected: "Hello."}
	c4 := TestCase{input: "“Hello.”", expected: "“Hello.”"}
	cases := []TestCase{c1, c2, c3, c4}
	for i, v := range cases {
		actual := SetPeriod(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: %s expected: %s", i+1, actual, v.expected)
		}
	}
}

func TestRemovePeriod(t *testing.T) {
	c1 := TestCase{input: "Hello", expected: "Hello"}
	c2 := TestCase{input: "“Hello”", expected: "“Hello”"}
	c3 := TestCase{input: "Hello.", expected: "Hello"}
	c4 := TestCase{input: "“Hello.”", expected: "“Hello”"}
	cases := []TestCase{c1, c2, c3, c4}
	for i, v := range cases {
		actual := RemovePeriod(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: %s expected: %s", i+1, actual, v.expected)
		}
	}
}

func TestTrimBefore(t *testing.T) {
	c1 := TestCase{input: "the. Quick brown fox jumps", expected: "Quick brown fox jumps"}
	c2 := TestCase{input: "the. “Quick brown fox jumps”", expected: "“Quick brown fox jumps”"}
	c3 := TestCase{input: ". Quick brown fox jumps", expected: "Quick brown fox jumps"}
	c4 := TestCase{input: "Quick brown fox jumps", expected: "Quick brown fox jumps"}
	cases := []TestCase{c1, c2, c3, c4}
	for i, v := range cases {
		actual := TrimBefore(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: '%s' expected: '%s'", i+1, actual, v.expected)
		}
	}
}

func TestTrimAfter(t *testing.T) {
	c1 := TestCase{input: "Quick brown fox", expected: "Quick brown fox"}
	c2 := TestCase{input: "Quick brown fox. Jumps", expected: "Quick brown fox."}
	c3 := TestCase{input: "“Quick brown fox.” Jumps", expected: "“Quick brown fox.”"}
	cases := []TestCase{c1, c2, c3}
	for i, v := range cases {
		actual := TrimAfter(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: '%s' expected: '%s'", i+1, actual, v.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	c1 := TestCase{input: "Hello", expected: "Hello"}
	c2 := TestCase{input: "hello", expected: "Hello"}
	c3 := TestCase{input: "“hello”", expected: "“Hello”"}
	c4 := TestCase{input: "“Hello”", expected: "“Hello”"}
	cases := []TestCase{c1, c2, c3, c4}
	for i, v := range cases {
		actual := Capitalize(v.input)
		if actual != v.expected {
			t.Errorf("test case %d failed with actual: '%s' expected: '%s'", i+1, actual, v.expected)
		}
	}
}
