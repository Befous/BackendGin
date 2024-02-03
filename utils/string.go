package utils

import (
	"strings"

	"github.com/Befous/BackendGin/models"
)

func GeojsonNameString(geojson []models.FullGeoJson) (result string) {
	var names []string
	for _, geojson := range geojson {
		names = append(names, geojson.Properties.Name)
	}
	result = strings.Join(names, ", ")
	return result
}
