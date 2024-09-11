package reload

import
(
	"strings"
)
func CleanText(s *[]string, prefix string) {
	slice := *s
	i := 0
	for i < len(slice) {
		if strings.HasPrefix(slice[i],prefix)|| strings.HasPrefix(slice[i],strings.ToUpper(prefix)){
			slice = append(slice[:i], slice[i+1:]...)
		} else {
			i++
		}
	}

	*s = slice
}