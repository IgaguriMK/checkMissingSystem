package checker

import "testing"

func TestIsAutogenPlanets_Empty(t *testing.T) {
	ag := IsAutogenPlanets("Test System", make([]string, 0))

	if ag != true {
		t.Fail()
	}
}

func TestIsAutogenPlanets_True(t *testing.T) {
	ag := IsAutogenPlanets(
		"Test System",
		[]string{
			"Test System A",
			"Test System A 1",
			"Test System A 2",
			"Test System A 2 a",
			"Test System B 1",
			"Test System B 2",
			"Test System AB 1",
			"Test System AB 2",
		},
	)

	if ag != true {
		t.Fail()
	}
}

func TestIsAutogenPlanets_False(t *testing.T) {
	ag := IsAutogenPlanets(
		"Test System",
		[]string{
			"Test System A",
			"Test System A 1",
			"Test System A 2",
			"Jameson",
			"Test System B 2",
			"Test System AB 1",
			"Test System AB 2",
		},
	)

	if ag != false {
		t.Fail()
	}
}
