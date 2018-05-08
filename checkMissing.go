package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/IgaguriMK/checkMissingSystem/checker"
)

func main() {
	var radius float64
	flag.Float64Var(&radius, "r", 200, "Search radius")

	flag.Parse()

	args := flag.Args()

	inputName := "bubblebodies.txt"
	if len(args) >= 1 {
		inputName = args[0]
	}

	ch := load(inputName, radius)

	for sys := range ch {
		bodyNames, autogen := checker.Suffixes(sys.Name, sys.Bodies)
		if !autogen {
			continue
		}

		trees := checker.BuildTree(bodyNames)

		if checker.CheckMissing(trees) {
			fmt.Printf("%s:\n", sys.Name)

			for _, n := range checker.GetAllTrees(trees) {
				fmt.Printf("    %s %s\n", sys.Name, n)
			}

			fmt.Println()
		}
	}

	fmt.Println("End")
}

type System struct {
	Name   string
	Bodies []string
}

func load(inputName string, radius float64) chan System {
	f, err := os.Open(inputName)
	if err != nil {
		log.Fatal("Open file error: ", err)
	}

	sc := bufio.NewScanner(f)

	if !sc.Scan() {
		log.Fatal("Too short input.")
	}

	ch := make(chan System, 16)

	go func() {
		defer f.Close()

		lastid := -1
		var bodies []string
		var lastSys string
		isIn := false

		radiusSq := radius * radius

		for sc.Scan() {
			line := sc.Text()

			id, x, y, z, sysname, bodyname := parseLine(line)

			if id != lastid {
				if isIn {
					ch <- System{
						Name:   lastSys,
						Bodies: bodies,
					}
				}

				lastid = id
				lastSys = sysname
				bodies = make([]string, 0)
				isIn = x*x+y*y+z*z < radiusSq
			}

			bodies = append(bodies, bodyname)
		}

		if isIn {
			ch <- System{
				Name:   lastSys,
				Bodies: bodies,
			}
		}

		close(ch)
	}()

	return ch
}

func parseLine(line string) (int, float64, float64, float64, string, string) {
	fields := strings.Split(line, "\t")

	if len(fields) != 6 {
		log.Fatalf("Invalid line, length is %d, %v\n", len(fields), fields)
	}

	id, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatalf("Can't parse id: %q", fields[0])
	}

	x, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		log.Fatalf("Can't parse x: %q", fields[1])
	}

	y, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		log.Fatalf("Can't parse x: %q", fields[2])
	}

	z, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		log.Fatalf("Can't parse x: %q", fields[3])
	}

	sysname := fields[4]
	bodyname := fields[5]

	return id, x, y, z, sysname, bodyname
}
