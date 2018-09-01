package chapterone

import "testing"

func TestCompressRepeatedCharacters(t *testing.T) {
	expected := "a3b2c1"
	actual := compressRepeatedCharacters("aaabbc")
	if expected != actual {
		t.Error("didnt compress string to: ", expected, "instead this was returned: ", actual)

	}
}
