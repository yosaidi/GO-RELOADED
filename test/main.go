package main

import (
	"reload"
	"strconv"
	"strings"
	"unicode"
)

func ProccessText(str string) string {
	strsplit := strings.Fields(str)
	vowelCheck(&strsplit)
	SpecificCaseModifier(&strsplit, "(cap,", Capitalize)
	Case(&strsplit)
	Base(&strsplit, reload.ConvertBase)
	SpecificCaseModifier(&strsplit, "(low,", strings.ToLower)
	SpecificCaseModifier(&strsplit, "(up,", strings.ToUpper)
	Ponctuation(&strsplit)
	AdjustSingleQuotes(&strsplit)
	CompleteSingleQuotes(&strsplit)
	return (strings.Join(strsplit, " "))

}

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

func SpecificCaseModifier(s *[]string, prefix string, f func(string) string) {
	slice := *s
	i := 1
	for i < len(slice) {

		if strings.HasPrefix(slice[i-1], prefix) && strings.HasSuffix(slice[i], ")") {
			num := slice[i][:len(slice[i])-1]
			prvnum, err := strconv.Atoi(num)
			if err != nil {
				return
			}

			for j := i - prvnum - 1; j < i; j++ {
				if j >= 0 {
					slice[j] = f(slice[j])
				}
			}
			slice = append(slice[:i-1], slice[i+1:]...)
		} else {
			i++
		}
	}
	*s = slice
}
func Ponctuation(s *[]string) {
	slice := *s
	pcts := []string{";", ",", "!", "?", ":", ";"}
	pctMap := make(map[string]bool)
	i := 1
	for _, ch := range pcts {
		pctMap[ch] = true
	}
	for i < len(slice) {

		if isPunctuation(slice[i]) {
			slice[i-1] += slice[i]
			slice = append(slice[:i], slice[i+1:]...)
		} else if pctMap[string(slice[i][0])] {
			slice[i-1] += string(slice[i][0])
			slice[i] = slice[i][1:]
		} else {
			i++
		}
	}
	*s = slice
}
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
	for i := 0; i < len(slice); i++ {
		if slice[i][0] == '\'' && slice[i][len(slice[i])-1] != '\'' {
			slice[i] += "'"
		} else if slice[i][0] != '\'' && slice[i][len(slice[i])-1] == '\'' {
			slice[i] = "'" + slice[i]
		}
	}
	*s = slice
}
func vowelCheck(s *[]string) {
	slice := *s
	vowels := []string{"a", "e", "i", "o", "u"}
	vowelMap := make(map[string]bool)
	for _, char := range vowels {
		vowelMap[string(char)] = true
	}
	for i := 0; i < len(slice)-1; i++ {
		if (slice[i] == "a" || slice[i] == "A") && vowelMap[string(slice[i+1][0])] {
			slice[i] += "n"
		}
	}
	*s = slice
}

func isPunctuation(s string) bool {
	for _, char := range s {
		if !unicode.IsPunct(char) {
			return false
		}
	}
	return true
}

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}
