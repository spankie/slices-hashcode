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
		// types maps out each pizza to its number of slices
		types := make(map[int]int)
		for ix, val := range *sliceNos {
			types[ix] = val
		}

		count, pizzasAdded := simulate(maxAndNo, sliceNos, &types)
		out(count, pizzasAdded, files[ix])
	}
}

// extract returns the integer equivalents of numbers in the slice parameter...translated into a slice of ints
func extract(slice []string) *[]int {

	var tmp []int
	for ix := range slice {
		ref, err := strconv.Atoi(slice[ix])
		// Because the error isn't supposed to occur at all, i'll handle it here
		if err != nil {
			log.Fatal(err)
			os.Exit(3)
		}
		tmp = append(tmp, ref)
	}
	return &tmp
}

// simulate does the main calculations of the program...it sorts the number of pizza slices slice from highest to lowest
// and then adds from the first element whilst checking if the accumulated total isn't
// More than the given maximum constraint, and then returns the numbers of
// Different pizzas to order and which types to order.
//
// maxNo holds the integers in the first line of the file, denoting the maximum number of slices allowed and the number of types of pizza
// slices holds the integer in the second line of the input file, denoting the number of slices for each type of pizza progressively
// types maps out each pizza to its number of slices
func simulate(maxNo *[]int, slices *[]int, types *map[int]int) (*int, *[]int) {

	maxAndNo := *maxNo //TODO revise
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
		} else {
			break
		}
	}

	// pizzatypes is the map of of the original pizza to their respective number of slices
	// pizzasAdded is a slice that'll hold the different types of pizzas added
	// (remember the pizzas are named progressively with numbers e.g type1,type2 etc..)
	pizzatypes := *types
	var pizzasAdded []int

	// addedSlice holds the current pizza Slice
	for _, addedSlice := range pizzaSlice {
		// Here we range over the pizzatypes map looking for a pizza type,
		// That has the number of slices that addedSlice is currently holding,
		// Once found, we add the key, which denotes the pizza type(explained in pizzasAdded declaration above), to the pizzasAdded slice
	loop:
		for key, val := range pizzatypes {
			if addedSlice == val {
				pizzasAdded = append(pizzasAdded, key)
				break loop
			}
		}
	}

	// Now because the output requires that the kinds of pizzas we order to be listed in ascending order...i sort the pizzasAdded
	// slice in ascending order
	sort.Slice(pizzasAdded, func(i, j int) bool {
		return pizzasAdded[i] < pizzasAdded[j]
	})
	return &count, &pizzasAdded
}

// out will write out output to files, named relative to the input file's name
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
