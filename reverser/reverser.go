package reverser

func Reverse(str string) string {
	reversed := ""
	for _, s := range str {
		reversed = string(s) + reversed
	}
	return reversed
}
