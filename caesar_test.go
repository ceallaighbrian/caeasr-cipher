package caesar

import "testing"

func TestEncode(t *testing.T) {
	text := "Hello, World"
	expected := "Khoor, Zruog"
	if result := Encode(text, 3); result != expected {
		t.Errorf("Hello() = %q, want %q", result, expected)
	}
}

func TestDecode(t *testing.T) {
	text := "Khoor, Zruog"
	expected := "Hello, World"
	if result := Decode(text, 3); result != expected {
		t.Errorf("Hello() = %q, want %q", result, expected)
	}
}
