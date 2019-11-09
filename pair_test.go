package pair

import (
	"testing"
)

func TestCreateAndSet(t *testing.T) {
	p := Create()

	p.Set(1, 2)

	if p[0] != 1 || p[1] != 2 {
		t.Errorf("Got %v, wanted (1, 2)", p)
	}

	p.Set("b", true)

	if p[0] != "b" || p[1] != true {
		t.Errorf("Got %v, wanted (\"b\", true)", p)
	}
}

func TestNewPair(t *testing.T) {
	p := NewPair(1, 2)

	if p[0] != 1 || p[1] != 2 {
		t.Errorf("%v is not (1, 2)", p)
	}
}

func TestEquals(t *testing.T) {
	p1 := NewPair(1, 2)
	p2 := NewPair(1, 2)

	got := p1.Equals(p2)

	if got == false {
		t.Errorf("%v.Equals(p2) is false, want true.", p1)
	}

	p1 = NewPair(1, "a")
	p2 = NewPair(1, "a")

	got = p1.Equals(p2)
	if got == false {
		t.Errorf("%v.Equals(p2) is false, want true", p1)
	}
}

func TestNotEquals(t *testing.T) {
	p1 := NewPair(1, 2)
	p2 := NewPair(2, 1)

	got := p1.Equals(p2)
	if got == true {
		t.Errorf("%v.Equals(%v) is true, want false", p1, p2)
	}
}
