package reload

func VowelCheck(s *[]string) {
	slice := *s
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	vowelMap := make(map[string]bool)
	for _, char := range vowels {
		vowelMap[string(char)] = true
	}
	for i := 0; i < len(slice)-1; i++ {
		if (slice[i] == "a" || slice[i] == "A") && vowelMap[slice[i+1]] {
			continue
		} else if (slice[i] == "a" || slice[i] == "A") && vowelMap[string(slice[i+1][0])] {
			slice[i] += "n"
		}
	}
	*s = slice
}
