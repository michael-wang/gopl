package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// Map from "dup line" to "map from filename to boolean" to keep track of files with dup line.
	// Notice Golang does not have set, I guess because set can be done with map from any type to bool.
	infiles := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", infiles)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, infiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			files := infiles[line]
			sep := ""
			for filename := range files {
				fmt.Print(sep + filename)
				sep = ", "
			}
			fmt.Print("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, infiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		files, ok := infiles[line]
		if !ok {
			files = make(map[string]bool)
			infiles[line] = files
		}
		files[filename] = true
	}
}
