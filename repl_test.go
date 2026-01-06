package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input string
		expected []string
	} {
		"all capitals": {input: "HELLO WORLD", expected: []string{"hello", "world"}},
		"white space start": {input: "    hello there", expected: []string{"hello", "there"}},
		"white space trailing": {input: "how are you doing?   ", expected: []string{"how", "are", "you", "doing?"}},
		"empty string": {input: "", expected: []string{}},
		"only whitespace string": {input: "          ", expected: []string{}},
		"white space start/trailing": {input: "    hi friend   ", expected: []string{"hi", "friend"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}