package reload

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	got := ReadFile("sample.txt")
	want := "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?\n"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestHexToDec(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"1A", "26"},
		{"17", "23"},
		{"2A", "42"},
		{"29A", "666"},
	}
	for _, tt := range tests {
		tt_batch := tt.got
		t.Run(tt_batch, func(t *testing.T) {
			ans := HexToDec(tt.got)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestBinaryToDec(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"101", "5"},
		{"10111", "23"},
		{"101010", "42"},
		{"1010011010", "666"},
		{"01111111111111111111111111111111", "2147483647"}, // note: biggest positive number that will fit in 32 bits when using the “two’s complement” notation
		{"11111111111111111111111111111111", "4294967295"},
	}
	for _, tt := range tests {
		tt_batch := tt.got
		t.Run(tt_batch, func(t *testing.T) {
			ans := BinaryToDec(tt.got)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	got := ToLower("Omae Wa Mou Shindeiru")
	want := "omae wa mou shindeiru"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestToUpper(t *testing.T) {
	got := ToUpper("Omae Wa Mou Shindeiru")
	want := "OMAE WA MOU SHINDEIRU"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCapitalize(t *testing.T) {
	got := Capitalize("ShinDeiru")
	want := "Shindeiru"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRemoveIndex(t *testing.T) {
	myWords := []string{"one", "two", "three", "four"}
	got := len(RemoveIndex(myWords, 2))
	want := 3
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
