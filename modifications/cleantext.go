package reload

func CleanText(s *[]string, flag string) {
	slice := *s
	i := 1
	for i < len(slice) {
		if slice[i] == flag {

			slice = append(slice[:i], slice[i+1:]...)
		} else {
			i++
		}
	}

	*s = slice
}
