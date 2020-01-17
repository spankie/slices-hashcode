package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var files = [...]string{
	"inputs/a_example.in",
	"inputs/b_small.in",
	"inputs/c_medium.in",
	"inputs/d_quite_big.in",
	"inputs/e_also_big.in",
}

func main() {
	for ix := range files {
		file, err := os.Open(files[ix])
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
			os.Exit(3)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		file.Close()

		maxAndNo := extract(strings.Split(lines[0], " "))
		sliceNos := extract(strings.Split(lines[1], " "))

		// Map to store the pizza types
		types := make(map[int]int)
		for ix, val := range *sliceNos {
			types[ix] = val
		}

		count, pizzasAdded := simulate(maxAndNo, sliceNos, &types)
		out(count, pizzasAdded, files[ix])
	}
}

func extract(slice []string) *[]int {

	var tmp []int
	for ix := range slice {
		ref, err := strconv.Atoi(slice[ix])
		// Because the error isn't supposed to occur at all i'll handle it here
		if err != nil {
			log.Fatal(err)
			return nil
		}
		tmp = append(tmp, ref)
	}
	return &tmp
}

func simulate(maxNo *[]int, slices *[]int, types *map[int]int) (*int, *[]int) {

	maxAndNo := *maxNo
	sliceNos := *slices
	max := maxAndNo[0]

	sort.Slice(sliceNos, func(i, j int) bool {
		return sliceNos[i] > sliceNos[j]
	})

	total := 0
	count := 0
	var pizzaSlice []int
	for _, val := range sliceNos {
		if (total + val) <= max {
			total += val
			pizzaSlice = append(pizzaSlice, val)
			count++
		}
	}
	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	fmt.Println("The total is ", total)
	fmt.Println()

	pizzatype := *types
	var pizzasAdded []int

	for _, addedSlice := range pizzaSlice {
	loop:
		for key, val := range pizzatype {
			if addedSlice == val {
				pizzasAdded = append(pizzasAdded, key)
				break loop
			}
		}
	}

	sort.Slice(pizzasAdded, func(i, j int) bool {
		return pizzasAdded[i] < pizzasAdded[j]
	})
	return &count, &pizzasAdded
}

func out(count *int, types *[]int, filename string) {
	filename = strings.TrimPrefix(filename, "inputs/")
	filename = strings.TrimSuffix(filename, ".in")
	filename = filename + "_output"

	fmt.Println(filename)
	fmt.Println(*count)
	fmt.Println(*types)
	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	fmt.Print("test")
	fmt.Print("test1")

}
