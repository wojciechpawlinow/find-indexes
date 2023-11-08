package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

// LoadValuesToSlice function reads the file and append values to the slice
func LoadValuesToSlice(path string, sPtr *[]int) (int, error) {
	if sPtr == nil {
		return 0, errors.New("slice pointer is nil")
	}

	// try to open the file
	f, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open the file: %w", err)
	}
	defer f.Close()

	// read the file line by line
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {

		// read single line
		line := scanner.Text()

		// convert to int
		v, err := strconv.Atoi(line)
		if err != nil {
			return i, fmt.Errorf("failed to parse line: %s at: %d, error: %w", line, i, err)
		}
		i++

		// append to slice under the given address
		*sPtr = append(*sPtr, v)
	}

	if err = scanner.Err(); err != nil {
		return i, fmt.Errorf("failed to scan: %w", err)
	}

	_, _ = color.New(color.FgGreen).Printf("\n=> loaded %d numbers to the slice\n\n", i)

	return i, nil
}
