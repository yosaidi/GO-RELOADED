package main

import (
	"fmt"
	"os"

	reload "reload/processtext"
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
	text = reload.ProccessText(text)
	err = os.WriteFile(output, []byte(text), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}





