package clean

import (
	"fmt"
	"testing"
)

func TestCleanPrefixWithAWord(t *testing.T) {
	testString := "The. The quick brown fox jumps over the lazy dog"
	result := Prefixes(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove prefix of %s", testString)
	}
}

func TestCleanPrefixWithAPeriod(t *testing.T) {
	testString := ". The quick brown fox jumps over the lazy dog"
	result := Prefixes(testString)
	if result != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove prefix of %s", testString)
	}
}

func TestCleanPostfixWithAWord(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. dog"
	result := PostFixes(testString)
	fmt.Println(testString)
	if result != "The quick brown fox jumps over the lazy" {
		t.Errorf("Failed to remove postfix of %s", testString)
	}
}

func TestCleanPostfixWithALetter(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. d"
	result := PostFixes(testString)
	if result != "The quick brown fox jumps over the lazy" {
		t.Errorf("Failed to remove postfix of %s", testString)
	}
}
