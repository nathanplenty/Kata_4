package main

import "fmt"

func longestCommonPrefix(stringIn []string) string {
	// Initialize longest common prefix as string to stringOut
	stringOut := ""
	// Get shortest length of Input Words
	shortesLength := len(stringIn[0])
	for i := 0; i < len(stringIn); i++ {
		if len(stringIn[i]) < shortesLength {
			shortesLength = len(stringIn[i])
		}
	}
	// Get each character as char
	for char := 0; char < shortesLength; char++ {
		// Get each word as word
		for word := 0; word < len(stringIn); word++ {
			// Check if a collumn of characters is the same rune
			if stringIn[word][char] != stringIn[0][char] {
				return stringOut
			}
		}
		// If collumn has same rune, add rune as character to stringOut
		stringOut += string(stringIn[0][char])
	}
	return stringOut
}

func main() {
	fmt.Printf("%q\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"apple", "apricot", "apex"}))
}
