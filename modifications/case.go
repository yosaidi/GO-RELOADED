package reload

import (
	"strconv"
	"strings"
)

func Case(s *[]string, prefix string, f func(string) string) {
	slice := *s
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], prefix) || strings.Contains(slice[i],strings.ToUpper(prefix)) {
			if strings.Contains(slice[i], prefix+",") || strings.Contains(slice[i],strings.ToUpper(prefix)+","){
				number, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
				if err != nil || number > len(slice) || number < 0 {

					continue
				}
				for j := i - number; j < i; j++ {
					slice[j] = f(slice[j])
				}
				slice = append(slice[:i], slice[i+2:]...)
			} else {
				slice[i-1] = f(slice[i-1])
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
	*s = slice
}
