package main

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"testing"
)

func TestCountStats(t *testing.T) {
	tests := []struct {
		name  string
		input string
		stats statistics
	}{
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
				sb.WriteString(fmt.Sprintf("%d", tc.stats.bytes))
				sb.WriteString(", got ")
				sb.WriteString(fmt.Sprintf("%d", result.bytes))

				sb.WriteString("\n\t\tlines: ")
				sb.WriteString(fmt.Sprintf("%d", tc.stats.lines))
				sb.WriteString(", got ")
				sb.WriteString(fmt.Sprintf("%d", result.lines))

				sb.WriteString("\n\t\twords: ")
				sb.WriteString(fmt.Sprintf("%d", tc.stats.words))
				sb.WriteString(", got ")
				sb.WriteString(fmt.Sprintf("%d", result.words))

				sb.WriteString("\n\t\tchars: ")
				sb.WriteString(fmt.Sprintf("%d\n", tc.stats.chars))
				sb.WriteString(", got ")
				sb.WriteString(fmt.Sprintf("%d", result.chars))
				t.Error(sb.String())
			}
		})
	}
}

func TestValidateFlag(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "-c",
			input:  "-c",
			result: true,
		},
		{
			name:   "-l",
			input:  "-l",
			result: true,
		},
		{
			name:   "-w",
			input:  "-w",
			result: true,
		},
		{
			name:   "-m",
			input:  "-m",
			result: true,
		},
		{
			name:   "False Flag",
			input:  "-f",
			result: false,
		},
		{
			name:   "Multiple Flags True",
			input:  "-cclmw",
			result: true,
		},
		{
			name:   "Multiple Flags False",
			input:  "-calmw",
			result: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := validate_flag(tc.input)
			if err != nil {
				t.Error(err)
			}
			if result != tc.result {
				t.Errorf("validate_flag(\"%s\")FAIL: expected %t, got %t", tc.input, tc.result, result)
			}
		})
	}
}

func TestParseFilesAndFlags(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		results struct {
			files []string
			flags [4]bool
		}
	}{
		{
			name:  "Single File",
			input: []string{"ccwc.go", "test.txt"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt"},
				flags: [4]bool{true, true, true, false},
			},
		},
		{
			name:  "Multiple Files",
			input: []string{"ccwc.go", "test.txt", "test2.txt"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt", "test2.txt"},
				flags: [4]bool{true, true, true, false},
			},
		},
		{
			name:  "Single File Single Flag",
			input: []string{"ccwc.go", "test.txt", "-c"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt"},
				flags: [4]bool{true, false, false, false},
			},
		},
		{
			name:  "Single File Multiple Flags",
			input: []string{"ccwc.go", "test.txt", "-c", "-l", "-m", "-w"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt"},
				flags: [4]bool{true, true, true, true},
			},
		},
		{
			name:  "Multiple Files Multiple Flags",
			input: []string{"ccwc.go", "test.txt", "test2.txt", "test3.txt", "-c", "-l", "-m", "-w"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt", "test2.txt", "test3.txt"},
				flags: [4]bool{true, true, true, true},
			},
		},
		{
			name:  "Multiple Files Multiple Flags Random Order",
			input: []string{"ccwc.go", "-c", "test.txt", "test2.txt", "-l", "test3.txt", "-m", "-w"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt", "test2.txt", "test3.txt"},
				flags: [4]bool{true, true, true, true},
			},
		},
		{
			name:  "Single File Single Flag with Multiple Options",
			input: []string{"ccwc.go", "test.txt", "-clmw"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt", "test2.txt", "test3.txt"},
				flags: [4]bool{true, true, true, true},
			},
		},
		{
			name:  "Single File Single Flag with Multiple Duplicate Options",
			input: []string{"ccwc.go", "test.txt", "-cclwmwlw"},
			results: struct {
				files []string
				flags [4]bool
			}{
				files: []string{"test.txt", "test2.txt", "test3.txt"},
				flags: [4]bool{true, true, true, true},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			files, flags, err := parse_files_and_flags(tc.input)

			if err != nil {
				log.Fatal(err)
			}

			if len(flags) != 0 {
				if tc.results.flags != flags {
					var sb strings.Builder
					sb.WriteString("invalid flags:\n\texpected: ")
					if tc.results.flags[0] {
						sb.WriteRune('c')
						sb.WriteString(" ")
					}
					if tc.results.flags[1] {
						sb.WriteRune('l')
						sb.WriteString(" ")
					}
					if tc.results.flags[2] {
						sb.WriteRune('w')
						sb.WriteString(" ")
					}
					if tc.results.flags[3] {
						sb.WriteRune('m')
					}

					sb.WriteString("\n\tgot:\t  ")
					if flags[0] {
						sb.WriteRune('c')
						sb.WriteString(" ")
					}
					if flags[1] {
						sb.WriteRune('l')
						sb.WriteString(" ")
					}
					if flags[2] {
						sb.WriteRune('w')
						sb.WriteString(" ")
					}
					if flags[3] {
						sb.WriteRune('m')
					}

					t.Error(sb.String())
				}
			}

			for _, file := range files[1:] {
				if !slices.Contains(tc.results.files, file) {
					t.Errorf("invalid file: %s", file)
				}
			}
		})
	}
}
