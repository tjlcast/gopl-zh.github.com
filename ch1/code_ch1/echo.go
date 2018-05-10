package main

import (
	"fmt"
	"os"
	"strings"
)


func main()  {
	input := os.Args

	execrise1(input)

	execrise2(input)
}

func execrise1(input []string)  {
	for idx, val := range input {
		fmt.Println(string(idx) + "-" + string(val)) ;
	}
}

func execrise2(input []string) {
	joinedStr := strings.Join(input, " ")
	fmt.Println(joinedStr)
}



