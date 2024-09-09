package reload

import (

	"strings"
)

func Case(s *[]string) {
	slice := *s
	i := 1
	for i < len(slice) {
		if slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		} else if slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		} else if slice[i] == "(cap)" {
			slice[i-1] = Capitalize(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			i++
		}
	}
	*s = slice
}
