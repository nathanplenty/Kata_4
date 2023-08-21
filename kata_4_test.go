package main

import (
	"testing"
)

func TestLCP(t *testing.T) {
	b := false
	sw := "str"
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
	sg := longestCommonPrefix([]string{"strong", "string", "nope"})
	if sw == sg {
		b = true
	}
	if b != true {
		t.Errorf("b = %v; want 'str'", sw)
	}
}
