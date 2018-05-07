package checker

import "strings"

func BuildTree(systemName string, bodyNames []string) []BodyTree {
	if len(bodyNames) == 0 {
		return nil // For dummy
	}

	res := make([]BodyTree, 0)

	for len(bodyNames) > 0 {
		selected := bodyNames[0]
		bodyNames = bodyNames[1:]

		name := strings.TrimPrefix(selected, systemName)
		name = strings.TrimLeft(name, " ")

		var childNames []string
		childNames, bodyNames = filterByPrefix(bodyNames, selected+" ")
		childs := BuildTree(selected, childNames)

		res = append(
			res,
			BodyTree{
				Name:   name,
				Childs: childs,
			},
		)
	}

	return res
}

func filterByPrefix(strs []string, prefix string) ([]string, []string) {
	ts := make([]string, 0)
	fs := make([]string, 0)

	for _, s := range strs {
		if strings.HasPrefix(s, prefix) {
			ts = append(ts, s)
		} else {
			fs = append(fs, s)
		}
	}

	return ts, fs
}
