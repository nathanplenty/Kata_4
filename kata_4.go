package main

import "fmt"

func longestCommonPrefix(stringIn []string) string {
	var stringOut string
	if stringIn == nil {
		return stringOut
	}
	shortesLength := len(stringIn[0])
	for i := 0; i < len(stringIn); i++ {
		if len(stringIn[i]) < shortesLength {
			shortesLength = len(stringIn[i])
		}
	}
	for char := 0; char < shortesLength; char++ {
		for word := 0; word < len(stringIn); word++ {
			if stringIn[word][char] != stringIn[0][char] {
				return stringOut
			}
		}
		stringOut += string(stringIn[0][char])
	}
	return stringOut
}

func main() {
	fmt.Printf("%q\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"apple", "apricot", "apex"}))
}
