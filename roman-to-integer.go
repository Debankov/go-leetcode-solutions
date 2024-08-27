func romanToInt(s string) int {
	result := 0
	var numMap = make(map[string]int)
	numMap["I"] = 1
	numMap["V"] = 5
	numMap["X"] = 10
	numMap["L"] = 50
	numMap["C"] = 100
	numMap["D"] = 500
	numMap["M"] = 1000
	numMap["IV"] = 4
	numMap["IX"] = 9
	numMap["XL"] = 40
	numMap["XC"] = 90
	numMap["CD"] = 400
	numMap["CM"] = 900

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if numMap[string(s[i])] < numMap[string(s[i+1])] {
				result += numMap[string(s[i])+string(s[i+1])]
				i++
			} else {
				result += numMap[string(s[i])]
			}
		} else {
			result += numMap[string(s[i])]
		}
		fmt.Println(result)

	}

	return result
}
