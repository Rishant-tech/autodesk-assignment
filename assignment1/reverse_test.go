package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "ascii", in: "hello", want: "olleh"},
		{name: "prompt example", in: "String; 2be reversed...", want: "gnirtS; eb2 desrever..."},
		{name: "keeps punctuation", in: "abc, def!", want: "cba, fed!"},
		{name: "unicode", in: "Go语言 123", want: "言语oG 321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.in); got != tt.want {
				t.Fatalf("ReverseWords(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
