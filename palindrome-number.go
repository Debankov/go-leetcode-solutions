func isPalindrome(x int) bool {
    if x < 0{
        return false
    }
    check := x
    reversedX := 0
	for x != 0 {
		digit := x % 10
		reversedX = reversedX*10 + digit
		x /= 10
	}
	if reversedX == check {
		return true
	} else {
		return false
	}