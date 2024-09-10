package reload

import (
	"strconv"
)

func Base(s *[]string) {
	slice := *s

	for i := 1; i < len(slice); i++ {
		if slice[i] == "(bin)" || slice[i] == "(BIN)" {
			temp, err := strconv.ParseInt(slice[i-1], 2, 64)
			if err != nil {
				return
			}
			slice[i-1] = strconv.Itoa(int(temp))
			slice = append(slice[:i], slice[i+1:]...)

		} else if slice[i] == "(hex)" || slice[i] == "(HEX)" {
			temp, err := strconv.ParseInt(slice[i-1], 16, 64)
			if err != nil {
				return
			}
			slice[i-1] = strconv.Itoa(int(temp))
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	DeleteUnwantedFlags(&slice)
	*s = slice
}
