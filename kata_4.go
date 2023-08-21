package main

import "fmt"

func longestCommonPrefix(str []string) string {
	// 1. Get Slice of Strings (Array / Stack / like these "[]string" )
	// 1.2. Check if enough Input is given
	if len(str) < 2 {
		return ""
	}
	c := ""
	cn := ""
	l := len(str)
	// 2. Take Slice[0]
	sl0 := str[0]
	// 3. Take Slice[n] (n = +1)
	n := 1
	sln := str[n]
	for {
		cn = ""
		sl := 0
		sln = str[n]
		// 3.1. Get shorter len of compared values
		if len(str[0]) <= len(str[n]) {
			sl = len(str[0])
		} else {
			sl = len(str[n])
		}
		// 4. Check Slice[0] element[m] with Slice[n] element[m] (m = + 1)
		for i := 0; i < sl; i++ {
			// 4.1. If True -> append char
			if sl0[i] == sln[i] {
				c = c + string(sl0[i])
				cn = cn + string(sl0[i])
			}
		}
		n = n + 1
		if len(c) > len(cn) {
			c = cn
		}
		// 5. Check if n >= n+1 ->
		if n >= l {
			break
		}
	}
	return c
}

func main() {
	fmt.Printf("%q\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"apple", "apricot", "apex"}))
}
