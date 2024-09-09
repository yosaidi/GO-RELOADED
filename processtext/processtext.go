package reload

import (
	"reload/modifications"

	"strings"
)

func ProccessText(str string) string {
	strsplit := strings.Fields(str)
	reload.VowelCheck(&strsplit)
	reload.SpecificCaseModifier(&strsplit, "(cap,", reload.Capitalize)
	reload.Case(&strsplit)
	reload.Base(&strsplit, reload.ConvertBase)
	reload.SpecificCaseModifier(&strsplit, "(low,", strings.ToLower)
	reload.SpecificCaseModifier(&strsplit, "(up,", strings.ToUpper)
	reload.AdjustSingleQuotes(&strsplit)
	reload.Ponctuation(&strsplit)
	reload.CompleteSingleQuotes(&strsplit)
	return (strings.Join(strsplit, " "))

}
