package utils

// In check if a string is in list.
func In(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// HasOne check if one element of first list is in second list.
func HasOne(items []string, all []string) bool {
	for _, item := range items {
		if In(item, all) {
			return true
		}
	}
	return false
}
