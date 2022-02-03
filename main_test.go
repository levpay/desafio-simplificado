package main

import (
	"testing"
)

func Test_for_shuffle(t *testing.T) {
	var victory int
	var defeat int
	choose := "head"

	for i := 0; i < 11; i++ {
		shuffle := shuffle()
		t.Log(shuffle)
		if shuffle == choose {
			victory++
		} else {
			defeat++
		}
	}

	if defeat == 0 {
		t.Error("expect defeat, got 0")
	}
}
