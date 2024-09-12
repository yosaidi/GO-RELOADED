package reload

func CleanText(s *[]string, prefix string) {
	slice := *s
	i := 0
	for i < len(slice) {
		if slice[i] == prefix {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			i++
		}
	}

	*s = slice
}
