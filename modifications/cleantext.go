package reload

import (
	"strings"
)

func CleanText(s *[]string, prefix string) {
	slice := *s
	i := 1
	for i < len(slice) {
		if strings.HasPrefix(slice[i], prefix) && i==1 {
			slice [i-1]+= slice[i][5:]
			slice = append(slice[:i], slice[i+1:]...)
		} else if strings.HasPrefix(slice[i], prefix) && i!=1 {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			i++
		}
	}

	*s = slice
}
