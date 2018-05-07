package checker

import "strings"

func Suffixes(systemName string, planetNames []string) ([]string, bool) {
	res := make([]string, 0, len(planetNames))

	for _, pn := range planetNames {
		if !strings.HasPrefix(pn, systemName) {
			return nil, false
		}

		pn = strings.TrimPrefix(pn, systemName)
		pn = strings.TrimLeft(pn, " ")

		res = append(res, pn)
	}

	return res, true
}
