package reload

func Ponctuation(s *[]string) {
	slice := *s

	i := 1
	for i < len(slice) {
		if CheckPunc(slice[i]) {
			slice[i-1] += slice[i]
			slice = append(slice[:i], slice[i+1:]...)

		} else if CheckPunc(string(slice[i][0])) {
			slice[i-1] += string(slice[i][0])
			slice[i] = slice[i][1:]
		} else {

			i++
		}
	}

	*s = slice
}
func CheckPunc(s string) bool {
	puncs := []rune{',', '.', '!', '?', ':', ';'}
	puncMap := make(map[rune]bool)
	for _, punc := range puncs {
		puncMap[punc] = true
	}
	for _, ch := range s {
		if !puncMap[ch] {
			return false
		}
	}
	return true
}

func InsertSpaceAfterPunc(s *[]string) {
	slice := *s
	for j, word := range slice {
		for i := 0; i < len(word)-1; i++ {
			if CheckPunc(string(word[i])) && !CheckPunc(string(word[i+1])) && word[len(word)-1] != '\'' {
				slice[j] = word[:i+1] + " " + word[i+1:]
				break
			}
		}
	}
	*s = slice
}
