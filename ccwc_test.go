package main

import (
	"fmt"
	"strings"
	"testing"
)

type test struct {
	name  string
	input string
	stats statistics
}

var tests []test = []test{
	{
		name:  "test1",
		input: "tests/test.txt",
		stats: statistics{
			bytes: 342190,
			lines: 7145,
			words: 58164,
			chars: 339292,
		},
	},
	{
		name:  "test2",
		input: "tests/test2.txt",
		stats: statistics{
			bytes: 25,
			lines: 6,
			words: 1,
			chars: 25,
		},
	},
	{
		name:  "test3",
		input: "tests/test3.txt",
		stats: statistics{
			bytes: 19,
			lines: 3,
			words: 1,
			chars: 19,
		},
	},
	{
		name:  "test4",
		input: "tests/test4.txt",
		stats: statistics{
			bytes: 5,
			lines: 0,
			words: 1,
			chars: 5,
		},
	},
}

func TestCountStats(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var sb strings.Builder
			result := count_stats(tc.input)
			if result != tc.stats {
				sb.WriteString("count_stats(\"")
				sb.WriteString(tc.input)
				sb.WriteString("\") FAIL: expected result:")

				sb.WriteString("\n\t\tbytes: ")
				fmt.Fprintf(&sb, "%d", result.bytes)
				sb.WriteString(", got ")
				fmt.Fprintf(&sb, "%d", tc.stats.bytes)

				sb.WriteString("\n\t\tlines: ")
				fmt.Fprintf(&sb, "%d", result.lines)
				sb.WriteString(", got ")
				fmt.Fprintf(&sb, "%d", tc.stats.lines)

				sb.WriteString("\n\t\twords: ")
				fmt.Fprintf(&sb, "%d", result.words)
				sb.WriteString(", got ")
				fmt.Fprintf(&sb, "%d", tc.stats.words)

				sb.WriteString("\n\t\tchars: ")
				fmt.Fprintf(&sb, "%d", result.chars)
				sb.WriteString(", got ")
				fmt.Fprintf(&sb, "%d\n", tc.stats.chars)
				t.Error(sb.String())
			}
		})
	}
}
