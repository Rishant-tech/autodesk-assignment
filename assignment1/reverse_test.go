package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "ascii", in: "hello", want: "olleh"},
		{name: "unicode", in: "Go语言", want: "言语oG"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); got != tt.want {
				t.Fatalf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
