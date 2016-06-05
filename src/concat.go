package Concat

func StringS(base []string) (func(...[]string)([]string)) {
  result := []string
  for _, v := range base {
    result = append(result, v)
  }
	return func(lists ...[]) ([]string) {
    for _, list := range lists {
      for _, v := range list {
        result = append(result, v)
      }
    }
    return result
  }
}

func IntS(base []int) (func(...[]int)([]int)) {
  result := []int
  for _, v := range base {
    result = append(result, v)
  }
	return func(lists ...[]) ([]int) {
    for _, list := range lists {
      for _, v := range list {
        result = append(result, v)
      }
    }
    return result
  }
}

func FloatS(base []float64) (func(...[]float64)([]float64)) {
  result := []float64
  for _, v := range base {
    result = append(result, v)
  }
	return func(lists ...[]) ([]float64) {
    for _, list := range lists {
      for _, v := range list {
        result = append(result, v)
      }
    }
    return result
  }
}

func BoolS(base []bool) (func(...[]bool)([]bool)) {
  result := []bool
  for _, v := range base {
    result = append(result, v)
  }
	return func(lists ...[]) ([]bool) {
    for _, list := range lists {
      for _, v := range list {
        result = append(result, v)
      }
    }
    return result
  }
}

func RuneS(base []rune) (func(...[]rune)([]rune)) {
  result := []rune
  for _, v := range base {
    result = append(result, v)
  }
	return func(lists ...[]) ([]rune) {
    for _, list := range lists {
      for _, v := range list {
        result = append(result, v)
      }
    }
    return result
  }
}

func ByteS(base []byte) (func(...[]byte)([]byte)) {
  result := []byte
  for _, v := range base {
    result = append(result, v)
  }
	return func(lists ...[]) ([]byte) {
    for _, list := range lists {
      for _, v := range list {
        result = append(result, v)
      }
    }
    return result
  }
}
