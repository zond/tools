
package tools

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	r := RandomString(10)
	if len(r) != 10 {
		t.Error("Should produce string of length 10, produced", r)
	}
}

func TestUuid(t *testing.T) {
	Uuid()
}

func bigIntTest(t *testing.T, i, b int, s string) {
	r := NewBigIntInt(i).BaseString(b)
	if r != s {
		t.Error("Wrong representation of ", i, " in base ",b,": wanted ",s," but got ",r)
	}
}

func bigIntTestString(t *testing.T, s string, b int, s2 string, b2 int) {
	r := NewBigIntString(s, b).BaseString(b2)
	if r != s2 {
		t.Error("Wrong representation of ", s, " in base ",b," when converted to base", b2, " : wanted ",s2," but got ",r)
	}
}

func TestBigInt(t *testing.T) {
	bigIntTestString(t, "45", 10, "50", 9)
	bigIntTestString(t, "1712", 8, "68A", 12)
	bigIntTestString(t, "45", 14, "2021", 3)
	bigIntTest(t, 10, 10, "10")
	bigIntTest(t, 11, 10, "11")
	bigIntTest(t, 1, 10, "1")
	bigIntTest(t, 4, 10, "4")
	bigIntTest(t, 14, 10, "14")
	bigIntTest(t, 61, 10, "61")
	bigIntTest(t, 615, 10, "615")
	bigIntTest(t, 11261, 10, "11261")
	bigIntTest(t, 10, 8, "12")
	bigIntTest(t, 11, 8, "13")
	bigIntTest(t, 1, 8, "1")
	bigIntTest(t, 4, 8, "4")
	bigIntTest(t, 14, 8, "16")
	bigIntTest(t, 61, 8, "75")
	bigIntTest(t, 615, 8, "1147")
	bigIntTest(t, 11261, 8, "25775")
	bigIntTest(t, 10, 16, "A")
	bigIntTest(t, 11, 16, "B")
	bigIntTest(t, 1, 16, "1")
	bigIntTest(t, 4, 16, "4")
	bigIntTest(t, 14, 16, "E")
	bigIntTest(t, 61, 16, "3D")
	bigIntTest(t, 615, 16, "267")
	bigIntTest(t, 11261, 16, "2BFD")
}