package clean

import (
	"fmt"
	"testing"
)

func TestRemoveWhiteSpaces(t *testing.T) {
	testString := " The quick brown fox jumps over the lazy dog "
	testString = WhiteSpaces(testString)
	if testString != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove whitespaces of %s", testString)
	}
}

func TestCleanPrefixWithAWord(t *testing.T) {
	testString := "The. The quick brown fox jumps over the lazy dog"
	testString = Prefixes(testString)
	if testString != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove prefix of %s", testString)
	}
}

func TestCleanPrefixWithAPeriod(t *testing.T) {
	testString := ". The quick brown fox jumps over the lazy dog"
	testString = Prefixes(testString)
	if testString != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("Failed to remove prefix of %s", testString)
	}
}

func TestCleanPostfixWithAWord(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. dog"
	testString = PostFixes(testString)
	fmt.Println(testString)
	if testString != "The quick brown fox jumps over the lazy" {
		t.Errorf("Failed to remove postfix of %s", testString)
	}
}

func TestCleanPostfixWithALetter(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy. d"
	testString = PostFixes(testString)
	if testString != "The quick brown fox jumps over the lazy" {
		t.Errorf("Failed to remove postfix of %s", testString)
	}
}
