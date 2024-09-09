package main

import (
	"fmt"
	"os"
	"strings"

	reload "reload/processtext"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	if strings.HasSuffix(os.Args[1], ".txt") && strings.HasSuffix(os.Args[2], ".txt") {

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
	} else {
		fmt.Print("Invalid file format!", "\n", "Correct file format: '.txt'", "\n")
	}
}
