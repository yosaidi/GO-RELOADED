package reload

import (
	reload "reload/modifications"

	"strings"
)

func ProccessText(str string) string {
	strsplit := strings.Fields(str)
	reload.Case(&strsplit, "(up", strings.ToUpper)
	reload.Case(&strsplit, "(cap", reload.Capitalize)
	reload.Case(&strsplit, "(low", strings.ToLower)
	reload.Base(&strsplit)
	reload.VowelCheck(&strsplit)
	reload.Ponctuation(&strsplit)
	reload.DeleteUnwantedFlags(&strsplit)
	reload.AdjustSingleQuotes(&strsplit)
	reload.CompleteSingleQuotes(&strsplit)
	return (strings.Join(strsplit, " "))

}
