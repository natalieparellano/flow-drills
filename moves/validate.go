package moves

import (
	"fmt"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

func Validate(inputs []byte) ([]Move, error) {
	var moves []Move

	err := yaml.Unmarshal(inputs, &moves)
	if err != nil {
		return moves, errors.Wrap(err, "yaml unmarshal failed")
	}

	for _, move := range moves {
		err = validate(move)
		if err != nil {
			return []Move{}, errors.Wrap(err, fmt.Sprintf("invalid move: %+v", move))
		}
	}

	return moves, nil
}

func ValidValues(propertyType string) []string {
	switch propertyType {
	case "timing":
		return []string{"split", "together"}
	case "direction":
		return []string{"same", "opposite"}
	case "position":
		return []string{"hip", "shoulder", "head"}
	case "forward leg":
		return []string{"left", "right"}
	case "poi":
		return []string{"left", "right"}
	}
	return []string{}
}

func validate(move Move) error {
	for idx, property := range move.Properties {
		valids := ValidValues(property.Type)
		if len(valids) == 0 {
			return errors.New(fmt.Sprintf("invalid property: %s", property.Type))
		}

		if len(property.Values) == 0 {
			move.Properties[idx].Values = valids
		}
		err := validateValues(move.Properties[idx].Values, valids)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateValues(inputs, valids []string) error {
	for _, input := range inputs {
		if !isMember(input, valids) {
			return errors.New(fmt.Sprintf("invalid property value: %s is not a member of %v", input, valids))
		}
	}
	return nil
}

func isMember(el string, arr []string) bool {
	for _, x := range arr {
		if el == x {
			return true
		}
	}
	return false
}
