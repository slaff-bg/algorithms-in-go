package stringreversing

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"Hello!", "!olleH"},
		{"Hello World!", "!dlroW olleH"},
		// {"こんにちは世界!", "!界世はちにんこ"}, // Japanese
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc.word), func(t *testing.T) {
			got := Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse: got %v want %v", got, tc.want)
			}
		})
	}
}

func TestUsingRunesReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"Hello World!", "!dlroW olleH"},
		{"こんにちは世界!", "!界世はちにんこ"}, // Japanese
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc.word), func(t *testing.T) {
			got := ReverseUsingRunes(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse: got %v want %v", got, tc.want)
			}
		})
	}
}

func Test_AlternativeReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"Good morning!", "!gninrom dooG"},
		{"Superstar", "ratsrepuS"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc.word), func(t *testing.T) {
			got := ReverseAlternative(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse: got %v want %v", got, tc.want)
			}
		})
	}
}

func Test_ReverseUsingStringsBilder(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"Helloween!", "!neewolleH"},
		{"Test word", "drow tseT"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc.word), func(t *testing.T) {
			got := ReverseUsingStringsBilder(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse: got %v want %v", got, tc.want)
			}
		})
	}
}
