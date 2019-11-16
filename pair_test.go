package pair

import (
	"testing"
)

func TestMake(t *testing.T) {
	p := Make(1, 2)

	if p[0] != 1 || p[1] != 2 {
		t.Errorf("%v is not (1, 2)", p)
	}
}

func TestSet(t *testing.T) {
	var p Pair

	p.Set(1, 2)

	if p[0] != 1 || p[1] != 2 {
		t.Errorf("Got %v, wanted (1, 2)", p)
	}

	p.Set("b", true)

	if p[0] != "b" || p[1] != true {
		t.Errorf("Got %v, wanted (\"b\", true)", p)
	}
}

func TestEquals(t *testing.T) {
	p1 := Make(1, 2)
	p2 := Make(1, 2)

	got := p1.Equals(p2)
	if got == false {
		t.Errorf("%v.Equals(p2) is false, want true.", p1)
	}

	p1.Set(1, "a")
	p2.Set(1, "a")

	got = p1.Equals(p2)
	if got == false {
		t.Errorf("%v.Equals(p2) is false, want true", p1)
	}

	got = p1.Equals(p1)
	if got == false {
		t.Errorf("%v.Equals(%v) is false, want true.", p1, p1)
	}
}

func TestNotEquals(t *testing.T) {
	var p1, p2 Pair

	p1.Set(1, 2)
	p2.Set(2, 1)

	got := p1.Equals(p2)
	if got == true {
		t.Errorf("%v.Equals(%v) is true, want false", p1, p2)
	}
}

func TestCompositeType(t *testing.T) {
	data := struct {
		X int
		Y int
	}{
		1,
		2,
	}

	p := Make(data, 2)

	if p[0] != data || p[1] != 2 {
		t.Errorf("Got %v, wanted ({1, 2}, 2)", p)
	}
}

func TestPairOfPairs(t *testing.T) {
	p1 := Make("a", "b")

	p2 := Make(p1, 2)

	p3 := Make(p1, 2)

	got := p2.Equals(p3)
	if got == false {
		t.Errorf("Got %v, wanted (`p2`, 2)", p1)
	}
}
