package reload

import (
	reload "reload/modifications"

	"strings"
)

func ProccessText(str string) string {
	strsplit := strings.Fields(str)
	reload.VowelCheck(&strsplit)
	reload.SpecificCaseModifier(&strsplit, "(cap,", reload.Capitalize)
	reload.Case(&strsplit)
	reload.Base(&strsplit)
	reload.SpecificCaseModifier(&strsplit, "(low,", strings.ToLower)
	reload.SpecificCaseModifier(&strsplit, "(up,", strings.ToUpper)
	reload.DeleteUnwantedFlags(&strsplit)
	reload.AdjustSingleQuotes(&strsplit)
	reload.Ponctuation(&strsplit)
	reload.CompleteSingleQuotes(&strsplit)
	return (strings.Join(strsplit, " "))

}
