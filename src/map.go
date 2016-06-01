package Map

// String takes a function and applies it to each value in the list, list of new values
func String(fn func(string, int, []string) string) func([]string) []string {
	return func(functor []string) (result []string) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Int takes a function and applies it to each value in the list, list of new values
func Int(fn func(int, int, []int) int) func([]int) []int {
	return func(functor []int) (result []int) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Int32 takes a function and applies it to each value in the list, list of new values
func Int32(fn func(int32, int, []int32) int32) func([]int32) []int32 {
	return func(functor []int32) (result []int32) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Float takes a function and applies it to each value in the list, list of new values
func Float(fn func(float64, int, []float64) float64) func([]float64) []float64 {
	return func(functor []float64) (result []float64) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Float32 takes a function and applies it to each value in the list, list of new values
func Float32(fn func(float32, int, []float32) float32) func([]float32) []float32 {
	return func(functor []float32) (result []float32) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Bool takes a function and applies it to each value in the list, list of new values
func Bool(fn func(bool, int, []bool) bool) func([]bool) []bool {
	return func(functor []bool) (result []bool) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Rune takes a function and applies it to each value in the list, list of new values
func Rune(fn func(rune, int, []rune) rune) func([]rune) []rune {
	return func(functor []rune) (result []rune) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}

// Byte takes a function and applies it to each value in the list, list of new values
func Byte(fn func(byte, int, []byte) byte) func([]byte) []byte {
	return func(functor []byte) (result []byte) {
		for i, v := range functor {
			result = append(result, fn(v, i, functor))
		}
		return
	}
}
