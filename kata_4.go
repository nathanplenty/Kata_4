package main

import "fmt"

func longestCommonPrefix(str []string) string {
	// Put Code here
	c := ""
	cn := ""
	fmt.Println("Input:", str)
	l := len(str)
	// fmt.Println(l)
	sl0 := str[0]
	n := 1
	sln := str[n]
	// fmt.Println(string(sl0))
	// fmt.Println(string(sln))
	for {
		// fmt.Println("Round:", n)
		f1 := func() {
			cn = ""
			sl := 0
			sln = str[n]
			if len(str[0]) <= len(str[n]) {
				sl = len(str[0])
			} else {
				sl = len(str[n])
			}
			// fmt.Println(sl)
			for i := 0; i < sl; i++ {
				// fmt.Println(sl0[i], sln[i])
				if sl0[i] == sln[i] {
					c = c + string(sl0[i])
					cn = cn + string(sl0[i])
				}
			}
		}
		f1()
		n = n + 1
		if len(c) > len(cn) {
			// wenn c größer als cn ist
			c = cn
		}
		if n >= l {
			break
		}
	}
	fmt.Print("Output: ")
	fmt.Printf("%q", c)
	fmt.Println()
	return c
}

func main() {
	//Pseudo Code
	// 1. Get Slice of Strings (Array / Stack / like these "[]string" )
	// 2. Take Slice[0]
	// 3. Take Slice[n] (n = +1)
	// 4. Check Slice[0] element[m] with Slice[n] element[m] (m = + 1)
	// 4.1. If True -> GoTo 4.
	// 4.2. If False -> GoTo 3
	// 5. Check if n >= n+1 ->
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})    // Output: ""
	longestCommonPrefix([]string{"apple", "apricot", "apex"}) // Output: "ap"
}
