package moves

import "math/rand"

type Move struct {
	Name       string         `yaml:"name"`
	Properties []MoveProperty `yaml:"properties"`
	Variations []string       `yaml:"variations"`
}

type MoveProperty struct {
	Type         string   `yaml:"type"`
	Values       []string `yaml:"values"`
	CurrentIndex int
}

func randomValue(values []string) string {
	n := rand.Int() % len(values)
	return values[n]
}
