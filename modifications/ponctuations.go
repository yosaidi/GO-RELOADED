package reload



func Ponctuation(s *[]string) {
	slice := *s
	puncs := []string{",", ".", "!", "?", ":", ";"}
	//punc at end
	for i, word := range slice {
		for _, punc := range puncs {
			if string(word[0]) == punc && slice[len(slice)-1] == slice[i] {
				slice[i-1] = slice[i-1] + word
				slice = slice[:len(slice)-1]
			}
		}
	}
	//middle punctuation
	for i, word := range slice {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && slice[i] != slice[len(slice)-1] {
				slice[i-1] = slice[i-1] + word
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
	// punc in the middle connect to word
	for i, word := range slice {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				slice[i-1] += punc
				slice[i] = word[1:]
			}
		}
	}
	*s = slice
}
