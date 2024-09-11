package main

import (
	"bufio"
	"fmt"
	"os"
	reload "reload/processtext"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	if strings.HasSuffix(os.Args[1], ".txt") && strings.HasSuffix(os.Args[2], ".txt") {

		input := os.Args[1]
		output := os.Args[2]

		texttochange, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer texttochange.Close()

		result, err := os.Create(output)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer result.Close()
		scan := bufio.NewScanner(texttochange)
		write := bufio.NewWriter(result)

		for scan.Scan() {
			line := scan.Text()
			processedline := reload.ProccessText(line)
			_, err := write.WriteString(processedline + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}

		}
		if err := scan.Err(); err != nil {
			fmt.Println(err)
			return
		}
		write.Flush()
	} else {
		fmt.Print("Invalid file format!", "\n", "Correct file format: '.txt'", "\n")
	}
}
