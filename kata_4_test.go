package main

import (
	"fmt"
	"testing"
)

func TestLCP(t *testing.T) {
	b := false
	sw := "str"
	fmt.Println("Test for Pass 'str': ")
	sg := longestCommonPrefix([]string{"strong", "string", "stress"})
	if sw == sg {
		b = true
	}
	if b != true {
		t.Errorf("b = %v; want 'str'", sw)
	}
}

func TestLCPempty(t *testing.T) {
	b := false
	sw := ""
	fmt.Println("Test for Pass '': ")
	sg := longestCommonPrefix([]string{"strong", "string", "nope"})
	if sw == sg {
		b = true
	}
	if b != true {
		t.Errorf("b = %v; want 'str'", sw)
	}
}
