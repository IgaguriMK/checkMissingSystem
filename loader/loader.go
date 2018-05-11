package loader

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type System struct {
	ID     int
	Name   string
	Bodies []string
}

func LoadAll(inputNames []string, radius float64) chan System {
	ch := make(chan System, 16)

	go func() {
		bufs := make([]*LoadBuffer, 0, len(inputNames))

		for _, name := range inputNames {
			lch := Load(name, radius)
			bufs = append(bufs, NewLoadBuffer(lch))
		}

		// Main scan
		for {
			// Scan all
			nbufs := make([]*LoadBuffer, 0, len(bufs))
			for _, b := range bufs {
				if b.Scan() {
					nbufs = append(nbufs, b)
				}
			}
			bufs = nbufs

			if len(bufs) == 0 {
				break
			}

			sort.Slice(bufs, func(i, j int) bool {
				return bufs[i].NextID() < bufs[j].NextID()
			})

			sys := bufs[0].Pop()

			for _, buf := range bufs[1:] {
				if sys.ID < buf.NextID() {
					break
				}

				s := buf.Pop()
				sys.Bodies = append(sys.Bodies, s.Bodies...)
			}

			ch <- sys
		}

		close(ch)
	}()

	return ch
}

type LoadBuffer struct {
	ch       chan System
	buf      System
	scanned  bool
	finished bool
}

func NewLoadBuffer(ch chan System) *LoadBuffer {
	return &LoadBuffer{
		ch: ch,
	}
}

func (lb *LoadBuffer) Scan() bool {
	if lb.finished {
		return false
	}

	if lb.scanned {
		return true
	}

	sys, ok := <-lb.ch

	if ok {
		lb.scanned = true
		lb.buf = sys
		return true
	} else {
		lb.finished = true
		return false
	}
}

func (lb *LoadBuffer) NextID() int {
	if lb.finished {
		panic("Read id from finished LoadBuffer")
	}

	if !lb.scanned {
		panic("Read id from unscanned LoadBuffer")
	}

	return lb.buf.ID
}

func (lb *LoadBuffer) Pop() System {
	if lb.finished {
		panic("Pop from finished LoadBuffer")
	}

	if !lb.scanned {
		panic("Pop from unscanned LoadBuffer")
	}

	lb.scanned = false

	return lb.buf
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
						ID:     lastid,
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
				ID:     lastid,
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
