package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// var declaration
	line := ""
	start_num := ""
	end_num := ""
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // loop through file lines
		line = scanner.Text()
		for _, char := range line { // loop through chars in string
			if char >= '0' && char <= '9' {
				if start_num == "" { //
					start_num = string(char)
					end_num = start_num
				} else {
					end_num = string(char)
				}
			}
		}

		// form the 2 digit number and convert to sum as an int
		start_num += end_num
		num, err := strconv.Atoi(start_num)
		if err != nil {
			log.Fatal(err)
		}
		sum += num

		start_num = ""
		end_num = ""
	}

	fmt.Println(sum)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
