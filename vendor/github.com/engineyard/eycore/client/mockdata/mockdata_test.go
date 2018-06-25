package mockdata

import (
	"testing"
)

var testArray []string

func setupTestArray() {
	testArray = []string{"one", "two", "three"}
}

func TestPop(t *testing.T) {
	setupTestArray()

	result, remainder := Pop(testArray)

	if len(remainder) != 2 {
		t.Error("Remainder should only be 2, got", len(remainder))
	}

	if result != testArray[2] {
		t.Error("Expected result to be last element of original array, got", result)
	}
}

func TestPeek(t *testing.T) {
	setupTestArray()

	result := Peek(testArray)

	if result != testArray[2] {
		t.Error("Expected result to be last element of original array, got", result)
	}
}
