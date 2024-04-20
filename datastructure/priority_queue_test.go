package datastructure

import (
	"testing"
)

func Test_hp(t *testing.T) {
	h := &hp{}
	h.push(&node{value: 13})
	h.push(&node{value: 1})
	h.push(&node{value: 5})
	h.push(&node{value: 2})
	for h.empty() == false {
		t.Log(h.pop().value)
	}
}
