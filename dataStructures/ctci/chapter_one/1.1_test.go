package chapterone

import (
	"testing"
)

func TestIsUniqueChars2(t *testing.T) {
	expected := false
	actual := isUniqueChars2("aa")
	if expected != actual {
		t.Error("not as expected")
	}
}
