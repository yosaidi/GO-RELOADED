package reload

import(
"unicode"
)
func Ponctuation(s *[]string) {
	slice := *s
	pcts := []string{";", ",", "!", "?", ":", ";"}
	pctMap := make(map[string]bool)
	i := 1
	for _, ch := range pcts {
		pctMap[ch] = true
	}
	for i < len(slice) {

		if isPunctuation(slice[i]) {
			slice[i-1] += slice[i]
			slice = append(slice[:i], slice[i+1:]...)
		} else if pctMap[string(slice[i][0])] {
			slice[i-1] += string(slice[i][0])
			slice[i] = slice[i][1:]
		} else {
			i++
		}
	}
	*s = slice
}

func isPunctuation(s string) bool {
	for _, char := range s {
		if !unicode.IsPunct(char) {
			return false
		}
	}
	return true
}

