package reverse_string

func ReverseString(input string) (output string) {

	r := []rune(input)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}

	return string(res)
}
