package main

import "fmt"

func longestCommonPrefix(str []string) string {
	// Initialize longest common prefix as string to LCP
	LCP := ""
	// Get shortest lenght of Input Words
	short := len(str[0])
	for i := 0; i < len(str); i++ {
		if len(str[i]) < short {
			short = len(str[i])
		}
	}
	// Get each character as i
	for i := 0; i < short; i++ {
		// Get each Word as j
		for j := 0; j < len(str); j++ {
			// Check if a collumn of characters is the same rune
			if str[j][i] != str[0][i] {
				return LCP
			}
		}
		// If collumn has same rune, add rune as character to LCP
		LCP += string(str[0][i])
	}
	return LCP
}

func main() {
	fmt.Printf("%q\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"apple", "apricot", "apex"}))
}
