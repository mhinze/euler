package main

import (
	"strconv"
)

func isPalindromeInt(number int64) bool {
	text := strconv.FormatInt(number, 10)
	return isPalindromeStr(text)
}

func isPalindromeStr(text string) bool {
	for a, x := 0, len(text)-1; a <= x; a, x = a+1, x-1 {
		if text[a] != text[x] {
			return false
		}
	}
	return true
}
