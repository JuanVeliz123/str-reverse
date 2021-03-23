package internal

import (
	"fmt"
	"testing"
)

func TestReverseParagraph_WithNilPointer_ReturnsEmptyString(t *testing.T) {
	result := ReverseParagraph(nil)
	if *result != "" {
		t.Error(fmt.Sprintf("Expected empty string, got %s", *result))
	}
}

func TestReverseParagraph_WithEmptyString_ReturnsEmptyString(t *testing.T) {
	paragraph := ""
	result := ReverseParagraph(&paragraph)
	if *result != "" {
		t.Error(fmt.Sprintf("Expected empty string, got %s", *result))
	}
}

func TestReverseParagraph_WithOneWord_ReturnsTheWord(t *testing.T) {
	paragraph := "testword"
	result := ReverseParagraph(&paragraph)
	if *result != paragraph {
		t.Error(fmt.Sprintf("Expected %s, got %s", paragraph, *result))
	}
}

func TestReverseParagraph_WithShortParagraph_ReturnsReversedParagraph(t *testing.T) {
	paragraph := "a couple of words"
	expected := "words of couple a"
	result := ReverseParagraph(&paragraph)
	if *result != expected {
		t.Error(fmt.Sprintf("Expected %s, got %s", expected, *result))
	}
}
