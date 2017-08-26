package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	var dict [256]int
	var maxLen int = 0
	var start int = -1
	for i := 0; i != len(s); i++ {
		if dict[s[i]] > start {
			start = dict[s[i]]
		}
		dict[s[i]] = i
		if i-start >= maxLen {
			maxLen = i - start
		}
	}
	fmt.Println(dict)
	return maxLen
}
