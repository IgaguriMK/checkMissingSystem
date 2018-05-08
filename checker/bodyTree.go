package checker

import (
	"fmt"
	"strconv"
	"strings"
)

type BodyTree struct {
	Name   string
	Childs []BodyTree
}

func (bt BodyTree) GetAll() []string {
	return bt.getAllInternal(make([]string, 0), "")
}

func (bt BodyTree) getAllInternal(res []string, prefix string) []string {
	name := bt.Name

	if prefix != "" {
		name = prefix + " " + name
	}

	res = append(res, name)

	for _, ch := range bt.Childs {
		res = ch.getAllInternal(res, name)
	}

	return res
}

func (bt BodyTree) Index() (string, int) {
	n := bt.Name

	if n == "" {
		return "", 0
	}

	ns := strings.Split(n, " ")
	indexStr := ns[len(ns)-1]

	prefix := ""
	if len(ns) > 1 {
		prefix = strings.Join(ns[:len(ns)-1], " ") + " "
	}

	if i, err := strconv.Atoi(indexStr); err == nil {
		if i <= 0 {
			panic("Negative indexStr")
		}

		return prefix, i
	}

	r := indexStr[0]
	switch {
	case 'A' <= r && r <= 'Z':
		return prefix, 1 + int(r-'A')
	case 'a' <= r && r <= 'z':
		return prefix, 1 + int(r-'a')
	}

	panic(fmt.Sprintf("Suould not reach: BodyTree#Index(), bt = %+v", bt))
}

func (bt BodyTree) GetTier() Tier {
	if bt.Name == "" {
		return SingleStar
	}

	ns := strings.Split(bt.Name, " ")
	indexStr := ns[len(ns)-1]

	if i, err := strconv.Atoi(indexStr); err == nil {
		if i <= 0 {
			panic("Negative indexStr")
		}

		return Planet
	}

	r := indexStr[0]
	switch {
	case 'A' <= r && r <= 'Z':
		return BinaryStar
	case 'a' <= r && r <= 'z':
		return Satellite
	}

	panic(fmt.Sprintf("Invalid index name: %q in %q", indexStr, bt.Name))
}

type Tier int

const (
	SingleStar Tier = iota
	BinaryStar
	Planet
	Satellite
)

func (t Tier) String() string {
	switch t {
	case SingleStar:
		return "SingleStar"
	case BinaryStar:
		return "BinaryStar"
	case Planet:
		return "Planet"
	case Satellite:
		return "Satellite"
	}

	panic(fmt.Sprintf("Illegal Tier value %d", int(t)))
}

func (t Tier) IndexName(index int) string {
	if t == SingleStar {
		return ""
	}

	switch t {
	case BinaryStar:
		r := 'A' + rune(index) - 1
		return string([]rune{r})
	case Planet:
		return strconv.Itoa(index)
	case Satellite:
		r := 'a' + rune(index) - 1
		return string([]rune{r})
	}

	panic("Should not reach")
}
