package tempconv

import "testing"

func TestKToC(t *testing.T) {
	k := Kelvin(0)
	got := KToC(k)
	exp := Celsius(-273.15)
	if got != exp {
		t.Errorf("KToC(%s) expected: %s, got: %s", k.String(), exp.String(), got.String())
	}

	k = 273.15
	got = KToC(k)
	exp = Celsius(0)
	if got != exp {
		t.Errorf("KToC(%s) expected: %s, got: %s", k.String(), exp.String(), got.String())
	}
}

func TestCToK(t *testing.T) {
	c := Celsius(0)
	got := CToK(c)
	exp := Kelvin(273.15)
	if got != exp {
		t.Errorf("CToK(%s) expected: %s, got: %s", c.String(), exp.String(), got.String())
	}

	c = 100
	got = CToK(c)
	exp = Kelvin(273.15 + 100)
	if got != exp {
		t.Errorf("CToK(%s) expected: %s, got: %s", c.String(), exp.String(), got.String())
	}
}

func TestKToF(t *testing.T) {
	k := Kelvin(323.15)
	got := KToF(k)
	exp := Fahrenheit(122)
	if got != exp {
		t.Errorf("KToF(%s) expected: %s, got: %s", k.String(), exp.String(), got.String())
	}

	k = 373.15
	got = KToF(k)
	exp = Fahrenheit(212)
	if got != exp {
		t.Errorf("KToF(%s) expected: %s, got: %s", k.String(), exp.String(), got.String())
	}
}

func TestFToK(t *testing.T) {
	f := Fahrenheit(122)
	got := FToK(f)
	exp := Kelvin(323.15)
	if got != exp {
		t.Errorf("FToK(%s) expected: %s, got: %s", f.String(), exp.String(), got.String())
	}

	f = 212
	got = FToK(f)
	exp = Kelvin(373.15)
	if got != exp {
		t.Errorf("FToK(%s) expected: %s, got: %s", f.String(), exp.String(), got.String())
	}
}
