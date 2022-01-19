package helpers

func RegionRulesCoef(countryCode string) float64 {
	// Nordic region country codes
	var nordicRegion = []string{"SE", "NO", "DK", "FI"}
	// EU country codes
	var euRegion = []string{"BE", "BG", "CZ", "DK", "DE", "EE", "IE", "EL", "ES", "FR", "HR", "IT", "CY", "LV", "LT", "LU", "HU", "MT", "NL", "AT", "PL", "PT", "RO", "SI", "SK", "FI", "SE"}

	// check is country code of Nordic region
	for i := range nordicRegion {
		if countryCode == nordicRegion[i] {
			return 1
		}
	}

	// check is country code of the EU
	for i := range euRegion {
		if countryCode == euRegion[i] {
			return 1.5
		}
	}

	// country code is outside the EU
	return 2.5
}
