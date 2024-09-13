package reload

import (
	"strconv"
	"strings"
)

func Case(s *[]string, prefix string, f func(string) string) {
	slice := *s
	if len(slice) == 1 {
		return
	}
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], prefix) {
			if strings.Contains(slice[i], prefix+",") && CheckFlag(slice[i+1][:len(slice[i+1])-1]) {
				number, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				if err != nil || number < 0 {
					continue
				}
				start := 0
				if i-number >= 0 {
					start = i - number
				}

				for j := start; j < i; j++ {
					slice[j] = f(slice[j])
				}

				slice[i-1] += slice[i+1][FindParenthese(slice[i+1])+1:]
				slice = append(slice[:i], slice[i+2:]...)
			} else {
				if i == 0 {
					continue
				} else if slice[i] == prefix+")" {
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

func CheckFlag(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true

}
