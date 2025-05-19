package main

import (
	"testing"
)

const test_file string = "test.txt"

func TestByteCount(t *testing.T) {
	const expected_result uint64 = 342190

	result := read_bytes(test_file)
	if result != expected_result {
		t.Errorf("read_bytes(\"%s\") FAIL: expected result %d, got %d.\n", test_file, expected_result, result)
	}
	t.Logf("PASS\n")
}

func TestLineCount(t *testing.T) {
	const expected_result uint64 = 7145

	result := read_lines(test_file)
	if result != expected_result {
		t.Errorf("read_lines(\"%s\") FAIL: expected result %d, got %d.\n", test_file, expected_result, result)
	}
	t.Logf("PASS\n")
}

func TestWordCount(t *testing.T) {
	const expected_result uint64 = 58164

	result := read_words(test_file)
	if result != expected_result {
		t.Errorf("read_word(\"%s\") FAIL: expected result %d, got %d.\n", test_file, expected_result, result)
	}
	t.Logf("PASS\n")
}
