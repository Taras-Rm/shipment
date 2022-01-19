package helpers

import "errors"

func WeightClassRulesCoef(weight float64) (uint, error) {

	// determining the type of weight class
	// small / medium / large / huge
	if weight > 0 && weight < 11 {
		return 100, nil
	} else if weight >= 11 && weight < 26 {
		return 300, nil
	} else if weight >= 26 && weight < 51 {
		return 500, nil
	} else if weight >= 51 && weight <= 1000 {
		return 500, nil
	}

	return 0, errors.New("uncorrect weight")
}
