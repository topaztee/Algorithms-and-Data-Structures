package chapterone

import (
	"testing"
)

func Test_permutationMethod1(t *testing.T) {
	expected := true
	actual := permutationMethod1("doG", "dog")

	if expected != actual {
		t.Error("not a permutation!")
	}
}

func Test_permutationMethod2(t *testing.T) {
	expected := true
	actual := permutationMethod2("dopG", "dog")
	if expected != actual {
		t.Error("not a permutation!")
	}
}
