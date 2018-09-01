package chapterone

import "testing"

func TestReverseString(t *testing.T) {
	actual := reverse("hello")
	expected := "olleh"

	if actual != expected {
		t.Error("didnt receive reversed string")
	}

}
