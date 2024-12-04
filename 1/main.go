package main

import (
	"bufio"
	"fmt"
	// "fmt"
	"os"
	// "slices"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input")
	check(err)

	c1 := make([]int, 1)
	c2 := make([]int, 1)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		v1, err := strconv.ParseInt(parts[0], 10, 64)
		check(err)
		c1 = append(c1, int(v1))

		v2, err := strconv.ParseInt(parts[1], 10, 64)
		check(err)
		c2 = append(c2, int(v2))
		// fmt.Printf("Number 1 is %d\n", v1)
		// fmt.Printf("Number 2 is %d\n", v2)

	}
	sort.Slice(c1, func(i int, j int) bool {
		return c1[i] < c1[j]
	})
	sort.Slice(c2, func(i int, j int) bool {
		return c2[i] < c2[j]
	})
	res := make([]int, 1)
	for i := range c1 {
		dif := abs(c1[i] - c2[i])
		res = append(res, dif)
	}
	var total int
	for _, num := range res {
		total += num
	}
	fmt.Printf("The total is %d\n", total)

	sim := make([]int, 1)
	for _, val := range c1 {
		tot := 0
		for _, n := range c2 {

			if val == n {
				tot += 1
			}
		}
		x := val * tot
		sim = append(sim, x)
	}
	var tot_sim int
	for _, num := range sim {
		tot_sim += num
	}
	fmt.Printf("The sim is %d\n", tot_sim)

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
