func simpleArraySum(ar []int32) int32 {
	var result int32
	for _, elem := range ar {
		result += elem
	}
	return result
}