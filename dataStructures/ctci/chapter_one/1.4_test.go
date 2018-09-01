package chapterone

import "testing"

func TestReplaceSpace(t *testing.T) {
	expected := "hello%20world"
	actual := replaceSpace("hello world")
	if expected != actual {
		t.Error("didnt replace spaces")
	}
}
