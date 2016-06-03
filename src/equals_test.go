package Equals

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("------ Equals String ------")
	fmt.Println("should take two strings and return a boolean")
	one := "a string"
	two := "a string"
	var result bool
	result = String(one, two)
	fmt.Println("should return true if the strings read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the strings do not read the same")
	result = String(one, "a different string")
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestStringS(t *testing.T) {
	fmt.Println("------ Equals String Slice ------")
	fmt.Println("should take two string slices and return a boolean")
	one := []string{"one", "two", "three"}
	two := []string{"one", "two", "three"}
	var result bool
	result = StringS(one, two)
	fmt.Println("should return true if the string slices read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the string slices do not read the same")
	result = StringS(one, []string{"four", "five"})
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestInt(t *testing.T) {
	fmt.Println("------ Equals Int ------")
	fmt.Println("should take two ints and return a boolean")
	one := 5
	two := 5
	var result bool
	result = Int(one, two)
	fmt.Println("should return true if the ints read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the ints do not read the same")
	result = Int(one, 2)
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestIntS(t *testing.T) {
	fmt.Println("------ Equals Int Slice ------")
	fmt.Println("should take two ints slices and return a boolean")
	one := []int{1, 2, 3}
	two := []int{1, 2, 3}
	var result bool
	result = IntS(one, two)
	fmt.Println("should return true if the int slices read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the int slices do not read the same")
	result = IntS(one, []int{4, 5})
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestFloat(t *testing.T) {
	fmt.Println("------ Equals Float64 ------")
	fmt.Println("should take two floats and return a boolean")
	one := 5.0
	two := 5.0
	var result bool
	result = Float(one, two)
	fmt.Println("should return true if the floats read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the floats do not read the same")
	result = Float(one, 2.0)
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestFloatS(t *testing.T) {
	fmt.Println("------ Equals Float64 Slice ------")
	fmt.Println("should take two float slices and return a boolean")
	one := []float64{1.0, 2.0, 3.0}
	two := []float64{1.0, 2.0, 3.0}
	var result bool
	result = FloatS(one, two)
	fmt.Println("should return true if the float slices read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the floats do not read the same")
	result = FloatS(one, []float64{4.0, 5.0})
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestBool(t *testing.T) {
	fmt.Println("------ Equals Bool ------")
	fmt.Println("should take two bools and return a boolean")
	one := true
	two := true
	var result bool
	result = Bool(one, two)
	fmt.Println("should return true if the bools read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the bools do not read the same")
	result = Bool(one, false)
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestBoolS(t *testing.T) {
	fmt.Println("------ Equals Bool Slice------")
	fmt.Println("should take two bool slices and return a boolean")
	one := []bool{true, false}
	two := []bool{true, false}
	var result bool
	result = BoolS(one, two)
	fmt.Println("should return true if the bool slices read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the bool slices do not read the same")
	result = BoolS(one, []bool{false, true})
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestRune(t *testing.T) {
	fmt.Println("------ Equals Rune ------")
	fmt.Println("should take two runes and return a boolean")
	one := 'a'
	two := 'a'
	var result bool
	result = Rune(one, two)
	fmt.Println("should return true if the runes read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the runes do not read the same")
	result = Rune(one, 'b')
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestRuneS(t *testing.T) {
	fmt.Println("------ Equals Rune ------")
	fmt.Println("should take two rune slices and return a boolean")
	one := []rune{'a', 'b'}
	two := []rune{'a', 'b'}
	var result bool
	result = RuneS(one, two)
	fmt.Println("should return true if the rune slices read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the rune slices do not read the same")
	result = RuneS(one, []rune{'c', 'd'})
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestByte(t *testing.T) {
	fmt.Println("------ Equals Byte ------")
	fmt.Println("should take two byte and return a boolean")
	var one byte = 'a'
	var two byte = 'a'
	var result bool
	result = Byte(one, two)
	fmt.Println("should return true if the byte read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the byte do not read the same")
	result = Byte(one, 'b')
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}

func TestByteS(t *testing.T) {
	fmt.Println("------ Equals Byte Slice ------")
	fmt.Println("should take two byte slices and return a boolean")
	one := []byte{'a', 'b'}
	two := []byte{'a', 'b'}
	var result bool
	result = ByteS(one, two)
	fmt.Println("should return true if the byte slices read the same")
	if result != true {
		t.Errorf("Expected 'true' and got 'false'")
	}
	fmt.Println("should return false if the byte slices do not read the same")
	result = ByteS(one, []byte{'c', 'd'})
	if result != false {
		t.Errorf("Expected 'false' and got 'true'")
	}
}
