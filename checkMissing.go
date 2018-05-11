package main

import (
	"flag"
	"fmt"

	"github.com/IgaguriMK/checkMissingSystem/checker"
	"github.com/IgaguriMK/checkMissingSystem/loader"
)

func main() {
	var radius float64
	flag.Float64Var(&radius, "r", 200, "Search radius")
	var maxCount int
	flag.IntVar(&maxCount, "n", 1000, "Limit of output")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		args = append(args, "bubblebodies.txt")
	}

	ch := loader.LoadAll(args, radius)

	count := 0
	for sys := range ch {
		bodyNames, autogen := checker.Suffixes(sys.Name, sys.Bodies)
		if !autogen {
			continue
		}

		trees := checker.BuildTree(bodyNames)

		if missingName, found := checker.CheckMissing(trees, sys.Name); found {
			fmt.Printf("%s (%d):\n", sys.Name, sys.ID)
			fmt.Printf("    Missing %q\n", missingName)
			fmt.Println("  in")

			for _, n := range checker.GetAllTrees(trees) {
				fmt.Printf("    %s %s\n", sys.Name, n)
			}

			fmt.Println()

			count++
			if count == maxCount {
				fmt.Println("Too many output")
				return
			}
		}
	}

	fmt.Println("Total", count, "systems.")
}
