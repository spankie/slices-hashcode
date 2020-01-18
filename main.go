package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// create output directory
	err := os.Mkdir("outputs", os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create outputs directory: %v\n", err)
	}
	// Read the `input` directory so that we don't have to
	// modify the code whenever we want to test other inputs
	_ = filepath.Walk("inputs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// for ix := range files {
			file, err := os.Open(path)
			if err != nil {
				// we dont want to alter the whole code if just one file does not open
				return err
				// log.Fatalf("failed to open file: %s", err)
				// os.Exit(3)
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

			pizzasAdded := simulate(maxAndNo, sliceNos, &types)
			out(pizzasAdded, path)
			// }
		}
		return nil
	})
	// fmt.Printf("Error reading the input directory: %v", err)
}

// extract returns the integer equivalents of numbers  on each line of the input file...translated into a slice
func extract(slice []string) *[]int {

	var tmp []int
	for ix := range slice {
		ref, err := strconv.Atoi(slice[ix])
		// Because the error isn't supposed to occur at all, i'll handle it here
		if err != nil {
			log.Fatal(err)
			return nil
		}
		tmp = append(tmp, ref)
	}
	return &tmp
}

// simulate does the main calculations of the program...it sorts the number of pizza slices slice from highest to lowest
// and then adds from the first element whilst checking if the accumulated total isn't
// More than the given maximum constraint, and then returns the numbers of
// Different pizzas to order and which types to order.
func simulate(maxNo *[]int, slices *[]int, types *map[int]int) *[]int {

	maxAndNo := *maxNo
	sliceNos := *slices
	max := maxAndNo[0]

	total := 0
	// var pizzaSlice []int
	var pizzasAdded []int

	// using this to loop through the slice from the back
	// this is because the larger values are at the end of the slice...
	// from the example file we have [2 5 6 8] right, the larger numbers are at the end of the slice
	for key := len(sliceNos) - 1; key >= 0; key-- {
		val := sliceNos[key]
		if (total + val) <= max {
			total += val
			// pizzaSlice = append(pizzaSlice, val)
			pizzasAdded = append(pizzasAdded, key)
		}
	}
	fmt.Println("The total is ", total)

	// Now because the output requires that the kinds of pizzas we buy to be listed in ascending order...i sort the pizzasAdded
	// slice in ascending order
	sort.Slice(pizzasAdded, func(i, j int) bool {
		return pizzasAdded[i] < pizzasAdded[j]
	})
	return &pizzasAdded
}

// out will write out output to files, named relative to the input file's name
func out(types *[]int, filename string) {
	// replace all occurence of "in" with "out"
	filename = strings.Replace(filename, "in", "out", -1)

	fmt.Println("Writing answer to file: ", filename)
	fmt.Println(len(*types), " types of pizza")
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Cannot create file for saving result: ", err)
	}
	_, err = f.Write([]byte(strconv.Itoa(len(*types)) + "\n" + strings.Trim(fmt.Sprint(types), "&[]")))
	if err != nil {
		fmt.Println("Cannot write result to file: ", err)
	}
	f.Sync()
	f.Close()
	fmt.Println()
}
