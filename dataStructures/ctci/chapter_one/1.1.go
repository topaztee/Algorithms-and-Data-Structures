/*Package chapterone there are 256 ASCII characters. This includes standard ASCII characters(0-127) and Extended ASCII characters(128-255).
* Ascii has 8 bytes a char. 1 is a parity bit so 2^7 = 128.
* Extended Ascii has 2^8=256. Unicode is a superset of Ascii.
 */
package chapterone

import (
	"fmt"
	"strings"
)

func isUniqueChars2(str string) bool {
	str = strings.ToLower(str)
	if len(str) > 128 {
		return false
	}
	var charSet [128]bool // every bool represents if a character was seen
	for _, val := range str {
		fmt.Println(val)
		if charSet[val] {

			return false
		}
		charSet[val] = true
	}
	return true
}
