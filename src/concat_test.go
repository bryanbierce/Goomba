package Concat

import (
	"fmt"
	"testing"

	equals "github.com/bryanbierce/Goomba/src/equals"
)

// func TestString(t *testing.T) {
// 	fmt.Println("------ Concat Strings ------")
// 	fmt.Println("should return a function")
// 	var result string
// 	starter := "This starts"
// 	var curried func(...string) string = String(starter)
// 	fmt.Println("should accept 1 new string and return it appended to the first")
// 	one := " a string"
// 	result = curried(one)
// 	if result != starter+one {
// 		t.Errorf("Expected 'this starts a string' and got %v", result)
// 	}
// 	fmt.Println("should accept any number of new strings and return them appended to the first in order")
// 	two := " which will grow"
// 	three := " as long as it is asked."
// 	four := " As long as this function works."
// 	result = curried(one, two, three, four)
// 	if result != starter+one+two+three+four {
// 		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
// 	}
// }

func TestStringS(t *testing.T) {
	fmt.Println("------ Concat Slice of Strings ------")
	fmt.Println("should return a function")
	var result []string
	starter := []string{"this", "is", "a"}
	var curried func(...[]string) []int = StringS(starter)

	fmt.Println("should accept 1 new slice and return it appended to the first")
	one := []string{"slice", "of", "strings"}
	result = curried(one)
	if equals.StringS(result, []string{"this", "is", "a", "slice", "of", "strings"}) {
		t.Errorf("Expected 'this starts a string' and got %v", result)
	}

	fmt.Println("should accept any number of new strings and return them appended to the first in order")
	two := []string{"that", "just"}
	three := []string{"keeps", "on"}
	four := []string{"growing", "longer"}
	result = curried(one, two, three, four)
	expected := []string{"this", "is", "a", "slice", "of", "strings", "that", "just", "keeps", "on", "growing", "longer"}
	if equals.StringS(result, expected) {
		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
	}
}

func TestIntS(t *testing.T) {
	fmt.Println("------ Concat Slice of Ints ------")
	fmt.Println("should return a function")
	var result []int
	starter := []int{1, 2, 3}
	var curried func(...[]int) []int = IntS(starter)

	fmt.Println("should accept 1 new slice and return it appended to the first")
	one := []int{4, 5, 6}
	result = curried(one)
	if equals.IntS(result, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Expected 'this starts a string' and got %v", result)
	}

	fmt.Println("should accept any number of new strings and return them appended to the first in order")
	two := []int{7, 8}
	three := []int{9, 10}
	four := []int{11, 12}
	result = curried(one, two, three, four)
	if equals.IntS(result, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}) {
		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
	}
}

func TestFloatS(t *testing.T) {
	fmt.Println("------ Concat Slice of Float64 ------")
	fmt.Println("should return a function")
	var result []float64
	starter := []float64{1, 2, 3}
	var curried func(...[]float64) []float64 = FloatS(starter)

	fmt.Println("should accept 1 new slice and return it appended to the first")
	one := []float64{4, 5, 6}
	result = curried(one)
	if equals.FloatS(result, []float64{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Expected 'this starts a string' and got %v", result)
	}

	fmt.Println("should accept any number of new strings and return them appended to the first in order")
	two := []float64{7, 8}
	three := []float64{9, 10}
	four := []float64{11, 12}
	result = curried(one, two, three, four)
	if equals.FloatS(result, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}) {
		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
	}
}

func TestBoolS(t *testing.T) {
	fmt.Println("------ Concat Slice of Bool ------")
	fmt.Println("should return a function")
	var result []bool
	starter := []bool{true, false}
	var curried func(...[]bool) []bool = BoolS(starter)

	fmt.Println("should accept 1 new slice and return it appended to the first")
	one := []bool{true, false}
	result = curried(one)
	if equals.BoolS(result, []bool{true, false, true, false}) {
		t.Errorf("Expected 'this starts a string' and got %v", result)
	}

	fmt.Println("should accept any number of new strings and return them appended to the first in order")
	two := []bool{true, false}
	three := []bool{true, false}
	four := []bool{true, false}
	result = curried(one, two, three, four)
	expected := []bool{true, false, true, false, true, false, true, false, true, false}
	if equals.BoolS(result, expected) {
		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
	}
}

func TestRuneS(t *testing.T) {
	fmt.Println("------ Concat Slice of Bool ------")
	fmt.Println("should return a function")
	var result []rune
	starter := []rune{'a', 'b'}
	var curried func(...[]rune) []rune = RuneS(starter)

	fmt.Println("should accept 1 new slice and return it appended to the first")
	one := []rune{'a', 'b'}
	result = curried(one)
	if equals.RuneS(result, []rune{'a', 'b', 'a', 'b'}) {
		t.Errorf("Expected 'this starts a string' and got %v", result)
	}

	fmt.Println("should accept any number of new strings and return them appended to the first in order")
	two := []rune{'a', 'b'}
	three := []rune{'a', 'b'}
	four := []rune{'a', 'b'}
	result = curried(one, two, three, four)
	expected := []rune{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}
	if equals.RuneS(result, expected) {
		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
	}
}

func TestByteS(t *testing.T) {
	fmt.Println("------ Concat Slice of Bytes ------")
	fmt.Println("should return a function")
	var result []byte
	starter := []byte{'a', 'b'}
	var curried func(...[]byte) []byte = ByteS(starter)

	fmt.Println("should accept 1 new slice and return it appended to the first")
	one := []byte{'a', 'b'}
	result = curried(one)
	if equals.ByteS(result, []byte{'a', 'b', 'a', 'b'}) {
		t.Errorf("Expected 'this starts a string' and got %v", result)
	}

	fmt.Println("should accept any number of new strings and return them appended to the first in order")
	two := []byte{'a', 'b'}
	three := []byte{'a', 'b'}
	four := []byte{'a', 'b'}
	result = curried(one, two, three, four)
	expected := []byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}
	if equals.ByteS(result, expected) {
		t.Errorf("Expected 'This starts a string which will grow as long as it is asked. As long as this function works.' and got %v", result)
	}
}
