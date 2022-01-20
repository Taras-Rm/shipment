package helpers

import (
	"github.com/biter777/countries"
)

func RegionRulesCoef(countryCode string) float64 {
	// Nordic region country codes
	var nordicRegion = []string{"SE", "NO", "DK", "FI"}

	// country info by code
	country := countries.ByName(countryCode)
	// get country region
	countryRegion := country.Region().String()

	// check is country code of Nordic region
	for i := range nordicRegion {
		if countryCode == nordicRegion[i] {
			return 1
		}
	}

	// check is country code of the EU
	if countryRegion == "Europe" {
		return 1.5
	}

	// country code is outside the EU
	return 2.5
}
