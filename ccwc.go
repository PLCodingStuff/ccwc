package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	var bytes uint64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		bytes += (uint64(len(line)) + 1)
	}

	return bytes
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
		if input_args[flag][1] != 'c' {
			fmt.Println("Invalid flag")
		}
		fmt.Println(read_bytes(file))
	}
}

func main() {
	parse_input()
}
