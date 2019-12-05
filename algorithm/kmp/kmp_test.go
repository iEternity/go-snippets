package kmp

import (
	"testing"
)

func TestKmp(t *testing.T) {
	cases := map[string]map[string]int{
		"kmpmpmmkmpkmpmmkmpmkmmmpkmpmmkmpmppp": {"kmpmmkmpm": 10},
		"abcd":                                 {"ab": 0},
		"abcdef":                               {"def": 3},
		"fjdkjfs":                              {"zk": -1},
	}

	for text, m := range cases {
		for sub, ans := range m {
			if Index(text, sub) != ans {
				t.Error(text, sub)
			}
		}
	}
}
