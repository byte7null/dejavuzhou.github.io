package codebase

import "testing"

func TestLeftRotateString(t *testing.T) {
	s := "1234567890"
	r := LeftRotateString(s,5)
	testEq(t,r,"6789012345")
}

func testEq(t *testing.T,result,target interface{}){
	if result != target{
		t.Error("fail")
	}
}