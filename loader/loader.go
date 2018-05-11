package loader

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type System struct {
	Name   string
	Bodies []string
}

func Load(inputName string, radius float64) chan System {
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
