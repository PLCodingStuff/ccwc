package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type statistics struct {
	bytes uint64
	lines uint64
	words uint64
	chars uint64
}

// func parse_dots(input_string string) [2][]byte {
// 	var filename_parse [2][]byte
// 	var word uint8 = 0
// 	char := 0

// 	for input_char := range len(input_string) {
// 		filename_parse[word] = append(filename_parse[word], input_string[input_char])
// 		if input_string[input_char] == 46 {
// 			word++
// 			char = 0
// 		}
// 		char++
// 	}

// 	return filename_parse
// }

/*
func read_bytes(filename string) uint64 {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var bytes uint64

	for {
		_, r_size, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		bytes += uint64(r_size)
	}

	return bytes

}

func read_lines(filename string) uint64 {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var lines uint64

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		if r == '\n' {
			lines++
		}
	}

	// Case only one line
	if lines == 0 {
		lines++
	}

	return lines
}

func split_words_seq(s string, sep []byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		var start, s_length uint64 = 0, uint64(len(s))

		for c := range s_length + 1 {
			if (c == s_length || slices.Contains(sep, s[c])) && s[start:c] != "" {
				if !yield(s[start:c]) {
					return
				}
				start = uint64(c + 1)
			}
		}
	}
}

func read_words(filename string) uint64 {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var words uint64
	var prev_r, l_space rune

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			if err.Error() == "EOF" {
				if prev_r != rune(0) && !unicode.IsSpace(prev_r) {
					words++
				}
				break
			}
			panic(err)
		}

		// Check for trailing left whitespace characters
		if l_space == rune(0) || unicode.IsSpace(l_space) {
			l_space = r
			prev_r = r
			continue
		}

		if !unicode.IsSpace(prev_r) && unicode.IsSpace(r) {
			words++
		}

		prev_r = r
	}

	return words
}

func read_chars(filename string) uint64 {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var runes uint64
	for {
		_, _, err = reader.ReadRune()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		runes++
	}

	return runes
}
*/

func count_stats(filename string) statistics {
	var stats statistics
	var prev_r, l_space rune

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		r, r_size, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				if prev_r != rune(0) && !unicode.IsSpace(prev_r) {
					stats.words++
				}
				break
			}
			panic(err)
		}

		if r == '\n' {
			stats.lines++
		}

		stats.bytes += uint64(r_size)
		stats.chars++

		// Check for trailing left whitespace characters
		if l_space == rune(0) || unicode.IsSpace(l_space) {
			l_space = r
			prev_r = r
			continue
		}

		if !unicode.IsSpace(prev_r) && unicode.IsSpace(r) {
			stats.words++
		}

		prev_r = r
	}

	// // Case only one line
	// if stats.lines == 0 {
	// 	stats.lines++
	// }
	return stats
}

func parse_input() ([]string, [][]byte, error) {
	input_args := os.Args

	if len(input_args) < 2 {
		return nil, nil, errors.New("No arguments provided")
	}

	var files []string
	var flags [][]byte

	// Parse cli args
	for arg := 1; arg < len(input_args); arg++ {
		// Case input arg is file
		if input_args[arg][0] != 45 {
			files = append(files, input_args[arg])
		} else {
			flags = append(flags, []byte(input_args[arg]))
		}
	}

	if len(files) == 0 {
		return nil, nil, errors.New("No files provided")
	}

	return files, flags, nil
}

func format_result_string(stats statistics, file string, l bool, w bool, c bool, m bool) string {
	var sb strings.Builder

	if l {
		fmt.Fprintf(&sb, "%d ", stats.lines)
	}
	if w {
		fmt.Fprintf(&sb, "%d ", stats.words)
	}
	if c {
		fmt.Fprintf(&sb, "%d ", stats.bytes)
	}
	if m {
		fmt.Fprintf(&sb, "%d ", stats.chars)
	}
	if !(c || l || w || m) {
		fmt.Fprintf(&sb, "%d ", stats.lines)
		fmt.Fprintf(&sb, "%d ", stats.words)
		fmt.Fprintf(&sb, "%d ", stats.bytes)
	}
	sb.WriteString(" " + file + "\n")

	return sb.String()
}

func compute_stats(file string, flags [][]byte) string {
	var print_c, print_l, print_w, print_m bool = false, false, false, false
	stats := count_stats(file)

	for _, flag := range flags {
		for _, char := range flag {
			switch char {
			case 'c':
				print_c = true
			case 'l':
				print_l = true
			case 'w':
				print_w = true
			case 'm':
				print_m = true
			default:
			}
		}
	}

	return format_result_string(stats, file, print_l, print_w, print_c, print_m)
}

func main() {
	files, flags, err := parse_input()

	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Print(compute_stats(file, flags))
	}
}
