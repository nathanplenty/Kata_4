package main

import (
	"testing"
)

func TestLCP(t *testing.T) {
	boolean := false
	stringTarget := "str"
	stringIn := longestCommonPrefix([]string{"strong", "string", "stress"})
	if stringTarget == stringIn {
		boolean = true
	}
	if boolean != true {
		t.Errorf("boolean = %v; want 'str'", stringTarget)
	}
}

func TestLCPempty(t *testing.T) {
	boolean := false
	stringTarget := ""
	stringIn := longestCommonPrefix([]string{"strong", "string", "nope"})
	if stringTarget == stringIn {
		boolean = true
	}
	if boolean != true {
		t.Errorf("boolean = %v; want 'str'", stringTarget)
	}
}
