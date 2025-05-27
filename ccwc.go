package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

type statistics struct {
	bytes uint64
	lines uint64
	words uint64
	chars uint64
}

func input_reader(filename string) *os.File {
	if filename != "stdin" {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		return file
	} else {
		return os.Stdin
	}
}

func count_stats(filename string) statistics {
	var stats statistics
	var prev_r, l_space rune

	file := input_reader(filename)
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
			log.Fatal(err)
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

	return stats
}

func validate_flag(flag string) (bool, error) {
	if len(flag) < 2 {
		return false, errors.New("invalid argument \"" + flag + "\"")
	}

	var inner_validate func(string) bool
	inner_validate = func(flag string) bool {
		if len(flag) == 0 {
			return true
		}
		switch flag[0] {
		case 'c', 'l', 'm', 'w':
			return inner_validate(flag[1:])
		}
		return false
	}

	return inner_validate(flag[1:]), nil
}

func parse_flag(flag string, flags [4]bool) ([4]bool, error) {
	updated_flags := flags
	is_val, err := validate_flag(flag)

	if err != nil {
		return [4]bool{}, err
	}

	if !is_val {
		return [4]bool{}, errors.New("unkown option \"" + flag + "\"")
	}

	for _, f := range flag[1:] {
		switch f {
		case 'c':
			updated_flags[0] = true
		case 'l':
			updated_flags[1] = true
		case 'w':
			updated_flags[2] = true
		case 'm':
			updated_flags[3] = true
		}
	}
	return updated_flags, nil
}

func parse_files_and_flags(input_args []string) ([]string, [4]bool, error) {
	var files []string
	flags := [4]bool{false, false, false, false}
	initial_setup := true

	// Parse cli args
	for _, arg := range input_args[1:] {
		if arg[0] != '-' {
			if !slices.Contains(files, arg) {
				files = append(files, arg)
			}
		} else {
			initial_setup = false

			var err error
			flags, err = parse_flag(arg, flags)

			if err != nil {
				return nil, [4]bool{}, err
			}
		}
	}

	if initial_setup {
		flags[0] = true
		flags[1] = true
		flags[2] = true
		flags[3] = false
	}
	if len(files) == 0 {
		files = append(files, "stdin")
	}

	return files, flags, nil
}

func format_result_string(stats statistics, file string, flags [4]bool) string {
	var sb strings.Builder

	if flags[1] {
		sb.WriteString(fmt.Sprintf("%7d", stats.lines))
	}
	if flags[2] {
		sb.WriteString(fmt.Sprintf("%7d", stats.words))
	}
	if flags[3] {
		sb.WriteString(fmt.Sprintf("%7d", stats.chars))
	}
	if flags[0] {
		sb.WriteString(fmt.Sprintf("%7d", stats.bytes))
	}
	if !(flags[0] || flags[1] || flags[2] || flags[3]) {
		sb.WriteString(fmt.Sprintf("%7d", stats.lines))
		sb.WriteString(fmt.Sprintf("%7d", stats.words))
		sb.WriteString(fmt.Sprintf("%7d", stats.bytes))
	}
	if file != "stdin" {
		sb.WriteString(" " + file)
	}
	sb.WriteRune('\n')

	return sb.String()
}

func main() {
	var total_stats statistics
	files, flags, err := parse_files_and_flags(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		stats := count_stats(file)
		fmt.Print(format_result_string(stats, file, flags))

		total_stats.bytes += stats.bytes
		total_stats.chars += stats.chars
		total_stats.lines += stats.lines
		total_stats.words += stats.words
	}

	if len(files) > 1 {
		fmt.Print(format_result_string(total_stats, "total", flags))
	}
}
