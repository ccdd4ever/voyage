//Given a string, find the length of the longest substring without repeating characters.
package problems

import "strings"

func lengthOfLongestSubstring(s string) int {
	length := 0
	pleft := 0
	pright := 0
	for i := 0; i < len(s); i++ {
		if pright >= len(s) {
			return length
		}
		if !strings.ContainsRune(s[pleft:pright], rune(s[pright+1])) {
			if length < pright-pleft {
				length = pright - pleft + 1
			}
			pright = pright + 1
		} else {
			pleft++
			pright++
		}
	}
	return length
}
