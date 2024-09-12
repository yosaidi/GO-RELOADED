package reload

func AdjustSingleQuotes(s *[]string) {
	slice := *s
	count := 0
	i := 0
	if len(slice) == 1 {
		return
	}
	for i < len(slice) {

		if (slice[0][0] == '\'' || slice[i] == "'") && count == 0 {
			count++
			slice[i] += slice[i+1]
			slice = append(slice[:i+1], slice[i+2:]...)

		} else if slice[i] == "'" && count%2 == 1 {
			slice[i-1] += slice[i]
			slice = append(slice[:i], slice[i+1:]...)
			count++

		} else if slice[i] == "'" && count%2 == 0 {
			count++
			if i+1 < len(slice) {

				slice[i] += slice[i+1]
				slice = append(slice[:i+1], slice[i+2:]...)
			}
		} else {
			i++
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
	if count%2 != 0 {

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
