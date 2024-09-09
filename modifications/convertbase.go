package reload

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
