package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	for i := len(r) - 1; i >= 0; i-- {
		output += string(r[i])
	}
	return
}
