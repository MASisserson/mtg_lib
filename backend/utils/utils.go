package utils

func SliceToMap(slice []string) map[string]bool {
	resultMap := make(map[string]bool)
	for _, val := range slice {
		resultMap[val] = true
	}
	return resultMap
}
