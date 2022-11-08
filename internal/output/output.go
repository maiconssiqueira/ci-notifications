package output

func KeysByValue(m map[string]bool, value bool) []string {
	var keys []string
	for k, v := range m {
		if value == v {
			keys = append(keys, k)
		}
	}
	return keys
}
