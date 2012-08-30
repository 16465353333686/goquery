package goquery

import (
	"testing"
)

func TestHasClass(t *testing.T) {
	sel := doc.Find("div")
	if !sel.HasClass("span12") {
		t.Error("Expected at least one div to have class span12.")
	}
}
