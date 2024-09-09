package reload

func DeleteUnwantedFlags(s *[]string) {
	slice := *s
	i := 0
	for i < len(slice) {
		if slice[i] == "(hex)" || slice[i] == "(bin)" || slice[i] == "(cap)" || slice[i] == "(low)" || slice[i] == "(up)" {
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			i++
		}
	}

	*s = slice
}
