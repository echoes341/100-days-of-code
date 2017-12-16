package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//enter a hex number -> return it in decimal form

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter a hex number (q to quit)")
		s.Scan()
		ln := s.Text()
		switch ln {
		case "q":
			return
		default:
			//take only the first
			fields := strings.Fields(ln) //strip by spaces
			var numStr string
			numStr = fields[0]
			num, err := strconv.ParseInt(numStr, 16, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Input is not a valid hex number.\n %s\n", err)
			} else {
				fmt.Printf("%x written in decimal form is: %d\n", num, num)
			}
		}
	}
}
