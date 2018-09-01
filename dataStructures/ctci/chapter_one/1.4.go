package chapterone

import "strings"

//Write a method to replace all spaces in a string with '%20'.
//You may assume that the string has sufficient space at the end of the string
//to hold the additional characters, and that you are given the "true" length of the string.
//(Note: if implementing in Java, please usea character array so that you can perform this operation inplace.)

func replaceSpace(str string) string {
	return strings.Replace(str, " ", "%20", len(str))
}
