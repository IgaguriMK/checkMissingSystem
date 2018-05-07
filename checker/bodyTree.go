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
