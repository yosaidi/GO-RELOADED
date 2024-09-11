package reload

import (
	"strconv"
	"strings"
)

func Case(s *[]string, prefix string, f func(string) string) {
	slice := *s
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], prefix) {
			if strings.Contains(slice[i], prefix+",") {
				number, err := strconv.Atoi(slice[i+1][:FindParenthese(slice[i+1])])
				if err != nil || number > len(slice) || number < 0 {
					continue
				}
				for j := i - number; j < i; j++ {
					if i > 0 {

						slice[j] = f(slice[j])
					}
				}
				slice[i-1] += slice[i+1][FindParenthese(slice[i+1])+1:]
				slice = append(slice[:i], slice[i+2:]...)
			} else {
				if i == 0 {
					slice = slice[1:]
				} else {
					slice[i-1] += slice[i][len(prefix)+1:]
					slice[i-1] = f(slice[i-1])
					slice = append(slice[:i], slice[i+1:]...)
				}
			}
		}
	}
	*s = slice
}

func FindParenthese(s string) int {
	count := 0

	for i, ch := range s {
		if ch == ')' {
			count = i
		}
	}
	return count
}
