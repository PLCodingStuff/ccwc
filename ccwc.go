package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

func parse_dots(input_string string) [2][]byte {
	var filename_parse [2][]byte
	var word uint8 = 0
	char := 0

	for input_char := 0; input_char < len(input_string); input_char++ {
		filename_parse[word] = append(filename_parse[word], input_string[input_char])
		if input_string[input_char] == 46 {
			word++
			char = 0
		}
		char++
	}

	return filename_parse
}

func read_bytes(filename string) uint64 {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var bytes uint64 = 0

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {

			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		bytes += uint64(len(line))
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

	var lines uint64 = 0

	for {
		_, err := reader.ReadBytes('\n')
		if err != nil {

			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		lines++
	}

	return lines
}

func iterate_words(s string, sep byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		var start, c, s_length uint64 = 0, 0, uint64(len(s))

		for c = 0; c < s_length; c++ {
			if c == s_length || s[c] == sep {
				if !yield(s[start:c]) {
					return
				}
				start = uint64(c) + 1
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

	var words uint64 = 0

	for {

		bytes_line, err := reader.ReadBytes('\n')
		if err != nil {

			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		string_line := string(bytes_line)

		// fmt.Println(string_line)
		iterate_words(string_line, ' ')(func(word string) bool {
			// fmt.Println(word)
			words++
			return true
		})
		fmt.Println()
	}

	return words
}

func parse_input() {
	input_args := os.Args

	if len(input_args) < 2 {
		fmt.Println("ERRRO: No file provided")
		return
	}

	file := input_args[len(input_args)-1]

	// Check if file is `txt`
	parsed_filename := parse_dots(file)
	if parsed_filename[1] == nil || string(parsed_filename[1]) != "txt" {
		fmt.Println("ERROR: Invalid file type")
		return
	}

	// Parse flags
	for flag := 1; flag < len(input_args)-1; flag++ {
		if input_args[flag][0] != 45 {
			fmt.Println("ERROR: Invalid flag")
			return
		}
		switch input_args[flag][1] {
		case 'c':
			fmt.Printf("%d", read_bytes(file))
		case 'l':
			fmt.Printf("%d", read_lines(file))
		case 'w':
			fmt.Printf("%d", read_words(file))
		}
		fmt.Printf(" %s\n", file)
	}
}

func main() {
	parse_input()
}
