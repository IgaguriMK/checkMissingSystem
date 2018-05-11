package checker

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type BodyTree struct {
	Name   string
	Childs []BodyTree
}

func (bt BodyTree) GetAll() []string {
	res := make([]string, 0)
	return bt.getAllInternal(res, "")

}

func GetAllTrees(bts []BodyTree) []string {
	res := make([]string, 0)

	for _, bt := range bts {
		res = append(res, bt.GetAll()...)
	}

	return res
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

func (bt BodyTree) Missing() (string, bool) {
	indexMap := make(map[string][]int)

	if len(bt.Childs) == 0 {
		return "", false
	}

	for _, c := range bt.Childs {
		pf, i := c.Index()

		if i < 0 {
			log.Printf("Index parse error: %q\n", bt.Name)
			return "", false
		}

		if len(pf) == 2 { // like "A "
			return joinName(bt.Name, pf+"::BinaryStar"), true
		}

		if _, ok := indexMap[pf]; !ok {
			indexMap[pf] = make([]int, 0)
		}
	}

	for _, c := range bt.Childs {
		pf, ci := c.Index()

		indexMap[pf] = append(indexMap[pf], ci)
	}

	tier := bt.Childs[0].GetTier()

	for _, cs := range indexMap {
		sort.Ints(cs)

		offset := 1

		for j, i := range cs {
			if i == 0 {
				offset = 0
				continue
			}

			if i > j+offset {
				indexName := tier.IndexName(j + offset)
				return joinName(bt.Name, fmt.Sprintf("%s ::Body", indexName)), true
			}
		}
	}

	for _, c := range bt.Childs {
		if name, found := c.Missing(); found {
			return joinName(bt.Name, name), true
		}
	}

	return "", false
}

func joinName(na, nc string) string {
	if na == "" {
		return nc
	} else {
		return na + " " + nc
	}
}

func CheckMissing(bts []BodyTree, systemName string) (string, bool) {
	tempBt := BodyTree{
		Name:   systemName,
		Childs: bts,
	}

	return tempBt.Missing()
}

func (bt BodyTree) Index() (string, int) {
	n := bt.Name

	if n == "" {
		return "", 1
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

	if len(indexStr) > 1 {
		return "", 1
	}

	r := indexStr[0]
	switch {
	case 'A' <= r && r <= 'Z':
		return prefix, 1 + int(r-'A')
	case 'a' <= r && r <= 'z':
		return prefix, 1 + int(r-'a')
	}

	return n, -1
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
		r := 'A' + rune(index-1)
		return string([]rune{r})
	case Planet:
		return strconv.Itoa(index)
	case Satellite:
		r := 'a' + rune(index-1)
		return string([]rune{r})
	}

	panic("Should not reach")
}

func (bt BodyTree) String() string {
	if len(bt.Childs) == 0 {
		return fmt.Sprintf("%q", bt.Name)
	}

	cs := make([]string, 0, len(bt.Childs))
	for _, c := range bt.Childs {
		s := c.String()
		cs = append(cs, s)
	}

	return fmt.Sprintf(
		"%q{ %s }",
		bt.Name,
		strings.Join(cs, ", "),
	)
}
