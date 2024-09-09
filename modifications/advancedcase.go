package reload

import (
	"strconv"
	"strings"
)

func SpecificCaseModifier(s *[]string, prefix string, f func(string) string) {
	slice := *s
	i := 1
	for i < len(slice) {

		if strings.HasPrefix(slice[i-1], prefix) && strings.HasSuffix(slice[i], ")") {
			num := slice[i][:len(slice[i])-1]
			prvnum, err := strconv.Atoi(num)
			if err != nil {
				return
			}

			for j := i - prvnum - 1; j < i; j++ {
				if j >= 0 {
					slice[j] = f(slice[j])
				}
			}
			slice = append(slice[:i-1], slice[i+1:]...)
		} else {
			i++
		}
	}
	*s = slice
}
