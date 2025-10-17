package main

import "fmt"

func checkPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(s)-1

	for left < right {
		if runes[left] != runes[right]{
			return false
		}
		left++
		right--
	}

	return true
}
func main(){
	fmt.Println(checkPalindrome("aaaaaaaaaa"))
	fmt.Println(checkPalindrome("abbbba"))
	fmt.Println(checkPalindrome("oekjddjks"))
}