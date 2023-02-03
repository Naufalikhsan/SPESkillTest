package helper

func NeedleInHaystack(haystack []string, needle string) int {
	for i, hay := range haystack {
		if hay == needle {
			return i
		}
	}

	return 0
}
