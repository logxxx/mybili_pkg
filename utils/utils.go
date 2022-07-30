package utils

func Contains(input string, slice []string) bool {
	for _, elem := range slice {
		if elem == input {
			return true
		}
	}
	return false
}
