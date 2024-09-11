package reload

func Ponctuation(s *[]string) {
	slice := *s

	i := 0
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
