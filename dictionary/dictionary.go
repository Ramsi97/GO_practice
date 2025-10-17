package main

import "fmt"

func count(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char] += 1
	}

	return freq
}
func main() {

	c := count("abcdefghygfiwbown")

	for key, val := range c {
		fmt.Printf("%c : %d\n", key, val)
	}
}