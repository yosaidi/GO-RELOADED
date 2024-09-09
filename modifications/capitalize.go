package reload

import(
	"strings"
)

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
}