package Map

import (
	"fmt"
	"testing"
)

func addLetter(str string, i int, list []string) string {
	return (str + "a")
}

func times2Int(num int, i int, list []int) int {
	return (num * 2)
}

func times2Int32(num int32, i int, list []int32) int32 {
	return (num * 2)
}

func times2Float(num float64, i int, list []float64) float64 {
	return (num * 2)
}

func times2Float32(num float32, i int, list []float32) float32 {
	return (num * 2)
}

func flipBool(truth bool, i int, list []bool) (res bool) {
	if truth == true {
		res = false
	} else {
		res = true
	}
	return
}

func replaceRune(rn rune, i int, list []rune) rune {
	return 'z'
}

func replaceByte(rn byte, i int, list []byte) byte {
	return 'Z'
}

func TestString(t *testing.T) {
	fmt.Println("------ MAP STRINGS ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]string) []string = String(addLetter)

	fmt.Println("maps simple functions over arrays")
	list := []string{"these", "are", "strings"}
	expected := []string{"thesea", "area", "stringsa"}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestInt(t *testing.T) {
	fmt.Println("------ MAP INT ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]int) []int = Int(times2Int)

	fmt.Println("maps simple functions over arrays")
	list := []int{1, 2, 3, 4}
	expected := []int{2, 4, 6, 8}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestInt32(t *testing.T) {
	fmt.Println("------ MAP INT32 ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]int32) []int32 = Int32(times2Int32)

	fmt.Println("maps simple functions over arrays")
	list := []int32{1, 2, 3, 4}
	expected := []int32{2, 4, 6, 8}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestFloat(t *testing.T) {
	fmt.Println("------ MAP FLOAT ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]float64) []float64 = Float(times2Float)

	fmt.Println("maps simple functions over arrays")
	list := []float64{1.0, 2.0, 3.0, 4.0}
	expected := []float64{2.0, 4.0, 6.0, 8.0}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestFlt32(t *testing.T) {
	fmt.Println("------ MAP FLOAT32 ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]float32) []float32 = Float32(times2Float32)

	fmt.Println("maps simple functions over arrays")
	list := []float32{1.0, 2.0, 3.0, 4.0}
	expected := []float32{2.0, 4.0, 6.0, 8.0}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestBool(t *testing.T) {
	fmt.Println("------ MAP BOOL ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]bool) []bool = Bool(flipBool)

	fmt.Println("maps simple functions over arrays")
	list := []bool{true, false, true, false}
	expected := []bool{false, true, false, true}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestRune(t *testing.T) {
	fmt.Println("------ MAP RUNE ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]rune) []rune = Rune(replaceRune)

	fmt.Println("runemaps simple functions over arrays")
	list := []rune{'a', 'b', 'c', 'd'}
	expected := []rune{'z', 'z', 'z', 'z'}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}

func TestByte(t *testing.T) {
	fmt.Println("------ MAP RUNE ------")
	fmt.Println("requires partial application")
	fmt.Println("returns a function which accepts a slice and returns a slice of the same type")
	var partial func([]byte) []byte = Byte(replaceByte)

	fmt.Println("runemaps simple functions over arrays")
	list := []byte{'A', 'B', 'C', 'D'}
	expected := []byte{'Z', 'Z', 'Z', 'Z'}
	res := partial(list)
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("Expected %v and got %v", expected[i], v)
		}
	}
	fmt.Println("")
}
