package checker

import "strings"

func IsAutogenPlanets(systemName string, planetNames []string) bool {
	for _, pn := range planetNames {
		if !strings.HasPrefix(pn, systemName) {
			return false
		}
	}

	return true
}
