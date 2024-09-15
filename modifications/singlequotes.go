package reload

func AdjustSingleQuotes(s *[]string) {
	slice := *s
	count := 0
	if len(slice) == 1 {
		return
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "'" && count == 0 && i+1<len(slice) {
			count++
			slice[i+1] = "'" + slice[i+1]
			slice = append(slice[:i], slice[i+1:]...)
		} else if slice[i] == "'" && count%2 == 1 && i>0{
			count++
			slice[i-1] = slice[i-1] + "'"
			slice = append(slice[:i], slice[i+1:]...)
		} else if slice[i] == "'" && count%2 == 0  && i+1<len(slice){
			count++
			slice[i+1] = "'" + slice[i+1]
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	*s = slice
}
