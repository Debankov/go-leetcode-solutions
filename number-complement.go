import "strconv"

func makeNumBinary(num int) string {
	resultStr := ""
	for num != 0 {
		remainder := num % 2
		resultStr = string(rune('0'+remainder)) + resultStr

		num /= 2
	}
	return resultStr
}

func findComplement(num int) int {
	binary := makeNumBinary(num)
	changeStr := ""
	resultInt := 0

	for _, v := range binary {
		switch v {
		case (49):
			changeStr += "0"
		case (48):
			changeStr += "1"
		}
	}
	parsedvalue, _ := strconv.ParseInt(changeStr, 2, 64)
	resultInt = int(parsedvalue)
	return resultInt
}