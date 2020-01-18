package main

import (
	"bufio"
	"fmt"
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
			file, err := os.Open(path)
			if err != nil {
				// we dont want to stop the whole app if just one file does not open
				return err
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

			pizzasAdded := simulate(maxAndNo, sliceNos)
			out(pizzasAdded, path)
		}
		return nil
	})
}

// extract returns the integer equivalents of numbers in the slice parameter...translated into a slice of ints
func extract(slice []string) *[]int {

	var tmp []int
	for ix := range slice {
		ref, err := strconv.Atoi(slice[ix])
		// Because the error isn't supposed to occur at all, i'll handle it here
		if err != nil {
			fmt.Println("Conversion failed")
			os.Exit(3) // we make this stringent because this error should never occur
		}
		tmp = append(tmp, ref)
	}
	return &tmp
}

// simulate does the main calculations of the program...it iterate the sliceNos slice from the end
// and then adds elements whilst checking if the accumulated total isn't
// More than the given maximum constraint, and then returns the a slice containing the pizza types ordered
func simulate(maxNo *[]int, slices *[]int) *[]int {

	maxAndNo := *maxNo
	sliceNos := *slices
	max := maxAndNo[0]

	total := 0
	var pizzasAdded []int

	// using this to loop through the slice from the back
	// this is because the larger values are at the end of the slice...
	// from the example file we have [2 5 6 8] right, the larger numbers are at the end of the slice
	for key := len(sliceNos) - 1; key >= 0; key-- {
		val := sliceNos[key]
		if (total + val) <= max {
			total += val
			pizzasAdded = append(pizzasAdded, key)
		}
	}

	// Now because the output requires that the kinds of pizzas we order to be listed in ascending order...i sort the pizzasAdded
	// slice in ascending order
	sort.SliceStable(pizzasAdded, func(i, j int) bool {
		return pizzasAdded[i] < pizzasAdded[j]
	})
	return &pizzasAdded
}

// out will write out output to files, named relative to the input file's name
func out(types *[]int, filename string) {
	// replace all occurence of "in" with "out"
	filename = strings.Replace(filename, "in", "out", -1)

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println("Cannot create file for saving result: ", err)
		return // return since there is no file to write to
	}
	_, err = f.Write([]byte(strconv.Itoa(len(*types)) + "\n" + strings.Trim(fmt.Sprint(*types), "&[]")))
	if err != nil {
		fmt.Println("Cannot write result to file: ", err)
	}
	f.Sync()
}
