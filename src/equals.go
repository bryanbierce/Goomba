package Equals

// String compares two strings and returns a bool indicated whether they are equal
func String(a string, b string) bool {
	return a == b
}

// StringS compares two string slices and returns a bool indicated whether they are equal
func StringS(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Int compares two ints and returns a bool indicated whether they are equal
func Int(a int, b int) bool {
	return a == b
}

// IntS compares two int slices and returns a bool indicated whether they are equal
func IntS(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Float compares two float slices and returns a bool indicated whether they are equal
func Float(a float64, b float64) bool {
	return a == b
}

// FloatS compares two float slices and returns a bool indicated whether they are equal
func FloatS(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Bool compares two bools and returns a bool indicated whether they are equal
func Bool(a bool, b bool) bool {
	return a == b
}

// BoolS compares two bool slices and returns a bool indicated whether they are equal
func BoolS(a []bool, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Rune compares two rune and returns a bool indicated whether they are equal
func Rune(a rune, b rune) bool {
	return a == b
}

// RuneS compares two rune slices and returns a bool indicated whether they are equal
func RuneS(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Byte compares two bytes and returns a bool indicated whether they are equal
func Byte(a byte, b byte) bool {
	return a == b
}

// ByteS compares two byte slices and returns a bool indicated whether they are equal
func ByteS(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
