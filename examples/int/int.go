package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/jesse-amano/rbst"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	var tree rbst.RBST

	fmt.Println("Welcome to the rbst example for ints.")
	fmt.Println("Please provide the first integer for the search tree:")
	prompt()

	for reader.Scan() {
		line := reader.Text()
		if len(line) == 0 {
			prompt()
			continue
		}
		if line[0] == 'q' {
			fmt.Println("Final tree:")
			fmt.Println(tree.String())
			break
		}

		var elem int64
		if line[0] == 'r' {
			elem = rand.Int63() - (1 << 62)
		} else {
			var err error
			elem, err = strconv.ParseInt(line, 10, 64)
			if err != nil {
				fmt.Printf("%q doesn't seem to be an integer. Please try again.\n", line)
				prompt()
				continue
			}
		}

		tree.Insert(intElem(elem))
		fmt.Println("Current tree:")
		fmt.Println(tree.String())
		fmt.Println("Please provide the next integer (q to quit, r for random):")
		prompt()
	}

	if err := reader.Err(); err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Println("Your integers from least to greatest:")
	fmt.Println(tree.Flatten())
}

type intElem int64

func (e intElem) Less(f interface{}) bool {
	return e < f.(intElem)
}

func prompt() (int, error) {
	return fmt.Print("\t» ")
}
