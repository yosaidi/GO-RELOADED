package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	texttochange, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(texttochange)
	text = ProccessText(text)
	err = os.WriteFile(output, []byte(text), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ProccessText(str string) string {
	strsplit := strings.Fields(str)
	vowelCheck(&strsplit)
	SpecificCaseModifier(&strsplit, "(cap,", Capitalize)
	Case(&strsplit)
	Base(&strsplit, ConvertBase)
	SpecificCaseModifier(&strsplit, "(low,", strings.ToLower)
	SpecificCaseModifier(&strsplit, "(up,", strings.ToUpper)
	AdjustSingleQuotes(&strsplit)
	Ponctuation(&strsplit)
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
	return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := toDecimal(nbr, baseFrom)
	return toBase(decimal, baseTo)
}

func toDecimal(nbr, baseFrom string) int {
	result := 0
	myPower := len(nbr) - 1
	baseLen := len(baseFrom)
	for i := 0; i < len(nbr); i++ {
		digit := basicAtoi(string(nbr[i]))
		if digit >= baseLen {
			panic("Invalid digit in input number for the given base")
		}
		result += digit * power(baseLen, myPower)
		myPower--
	}
	return result
}

func toBase(nbr int, base string) string {
	if nbr == 0 {
		return string(base[0])
	}
	strSlice := []rune(base)
	baseLen := len(strSlice)
	var result []rune
	for nbr > 0 {
		remainder := nbr % baseLen
		result = append([]rune{strSlice[remainder]}, result...)
		nbr /= baseLen
	}
	return string(result)
}

func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func basicAtoi(ch string) int {
	switch ch {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "A", "a":
		return 10
	case "B", "b":
		return 11
	case "C", "c":
		return 12
	case "D", "d":
		return 13
	case "E", "e":
		return 14
	case "F", "f":
		return 15
	default:
		return 0
	}
}
