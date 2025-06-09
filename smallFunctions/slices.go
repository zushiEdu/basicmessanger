package smallFunctions

func Contains(number int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if number == list[i] {
			return true
		}
	}
	return false
}
