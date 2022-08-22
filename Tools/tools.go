package Tools

func Contains[T comparable](array []T, obj T) bool {
	for _, a := range array {
		if a == obj {
			return true
		}
	}
	return false
}
