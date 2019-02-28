package test

import (
	"testing"

	"github.com/test0/ext"
)

func TestCheck(t *testing.T) {
	s := "dialo"
	m, h := ext.Check(s)

	if m != 3 && h != 2 {
		t.Errorf("Resutl should be 'm=3 and h=2'")
	}
}
