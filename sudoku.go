package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}


// Implement our own sets
type valueSet map[int]bool

// Initialize a valueSet with 1-9
func newValueSet() valueSet {
	valueMap := valueSet{}
	for i := 1; i <= 9; i++ {
		valueMap[i] = true
	}
	return valueMap
}

// Union of two valueSets
func union(value1 valueSet, value2 valueSet) valueSet {
	myValueSet := valueSet{}
	for key, value := range value1 {
		if value {
			myValueSet[key] = value
		}
	}
	for key, value := range value2 {
		if value {
			myValueSet[key] = value
		}
	}
	return myValueSet
}

// Intersection of two valueSets
func intersection(value1 valueSet, value2 valueSet) valueSet {
	myValueSet := valueSet{}
	for key, value := range value1 {
		if (value && value2[key]) {
			myValueSet[key] = value
		}
	}
	return myValueSet
}

// Represents one square of a Sudoku puzzle.
type Cell struct {
	// The determined value (meaningless if valueKnown is false)
	value int
	// Stores which values the cell is allowed to be (meaningless if valueKnown is true)
	valueOptions map[int]bool
	// Whether the value is known
	valueKnown bool
}

func newCell() *Cell {
	cell := new(Cell)
	cell.valueOptions = newValueSet()
	return cell
}

// Returns the number of values the cell is allowed to be in
func (c *Cell) numValueOptions() int {
	num := 0
	for _, value := range c.valueOptions {
		if value {
			num += 1
		}
	}
	return num
}

// Just print the value of a Cell f
func (c *Cell) String() string {
	valueString := strconv.Itoa(c.value)
	if !(c.valueKnown) {
		valueString = "0"
	}
	return valueString
}

type Puzzle struct {
	puzzle [][]Cell
}

func newPuzzle() *Puzzle {
	myPuzzle := new(Puzzle)
	myPuzzle.puzzle = make([][]Cell, 9)
	for i := 0; i < 9; i++ {
		myPuzzle.puzzle[i] = make([]Cell, 9)
		for j := 0; j < 9; j++ {
			myPuzzle.puzzle[i][j] = *newCell()
		}
	}
	return myPuzzle
}

func (myPuzzle *Puzzle) setValue(rowNum int, colNum int, value int) {
	myPuzzle.puzzle[rowNum][colNum].value = value
	myPuzzle.puzzle[rowNum][colNum].valueKnown = true
}

func (myPuzzle *Puzzle) insertRow(row string, rowNum int) {
	for colNum, char := range row {
		value := int(char)
		myPuzzle.setValue(rowNum, colNum, value)
	}
}

func box(cells [][]Cell) string {
	var box []string
	for _, row := range cells {
		var rowValues []string
		for _, col := range row {
			rowValues = append(rowValues, col.String())
		}
		box = append(box, strings.Join(rowValues, ""))
	}
	return strings.Join(box, "\n")
}

func (myPuzzle *Puzzle) String() string {
	return box(myPuzzle.puzzle)
}

func isNumeric(row string) bool {
	for _, char := range row {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func validateRowFormat(row string) error {
	if !(isNumeric(row)) {
		return fmt.Errorf("File contains non-numeric characters in row: ", row)
	}

	if len(row) != 9 {
		return fmt.Errorf("Row does not contain 9 digits: ", row) 
	}

	return nil
}

func puzzleFromFile(filename string) *Puzzle {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myPuzzle := newPuzzle()
	
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i >= 9 {
			log.Fatal("Files contains more than 9 rows.")
		}
		err := validateRowFormat(row)
		if err != nil {
			log.Fatal(err)
		}

		myPuzzle.insertRow(row, i)
		i++
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return myPuzzle
}

func main() {
	fmt.Printf("hello, world\n")
}
