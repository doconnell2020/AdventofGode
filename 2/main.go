package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("./short")
	check(err)
	defer f.Close()
	fmt.Println(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf(scanner.Text())
	}

	// numbers = data_to_nums(data)

}
