package reload

func Base(s *[]string, f func(nbr, from, to string) string) {
	slice := *s
	for i := 1; i < len(slice); i++ {
		if slice[i] == "(bin)" {
			slice[i-1] = f(slice[i-1], "01", "0123456789")
			slice = append(slice[:i], slice[i+1:]...)
		}
		if slice[i] == "(hex)" {
			slice[i-1] = f(slice[i-1], "0123456789ABCDEF", "0123456789")
			slice = append(slice[:i], slice[i+1:]...)
		}

	}
	*s = slice
}
