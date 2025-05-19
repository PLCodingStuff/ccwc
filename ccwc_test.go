package main

import (
	"testing"
)

const test_file string = "test.txt"

func Testbyte_count(t *testing.T) {
	const expected_result uint64 = 342190
	result := read_bytes(test_file)
	if result != expected_result {
		t.Errorf("reab_bytes(\"%s\") \"FAIL: expected result %d, got %d.\n", test_file, expected_result, result)
	}
	t.Logf("PASS\n")
}
