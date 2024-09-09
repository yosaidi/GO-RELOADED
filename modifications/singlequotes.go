package reload

func AdjustSingleQuotes(s *[]string) {
	slice := *s
	insidequotes := false
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == "'" {
			if !insidequotes {
				slice[i] += slice[i+1]
				slice = append(slice[:i+1], slice[i+2:]...)
				insidequotes = true
			} else if insidequotes {
				slice[i-1] += slice[i]
				slice = append(slice[:i], slice[i+1:]...)
				insidequotes = false
			}
		}

	}

	*s = slice
}

func CompleteSingleQuotes(s *[]string) {
	slice := *s
	insidequotes := false
	count := 0
	for _, word := range slice {
		for _, ch := range word {
			if ch == '\'' {
				count++
			}
		}
	}
	if count%2 != 2 {

		for i := 0; i < len(slice); i++ {
			if slice[i][0] == '\'' {
				if !insidequotes {
					insidequotes = true
					continue

				} else if insidequotes {
					slice[i-1] += "'"
					slice[i] = slice[i][1:]
					insidequotes = false
				}
			}
		}
	}
	*s = slice
}
