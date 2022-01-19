package helpers

// determining the type of weight class
func WeightClassRulesCoef(weight float64) uint {

	// small / medium / large / huge
	if weight > 0 && weight < 11 {
		return 100
	} else if weight >= 11 && weight < 26 {
		return 300
	} else if weight >= 26 && weight < 51 {
		return 500
	} else if weight >= 51 && weight <= 1000 {
		return 2000
	}

	return 0
}
