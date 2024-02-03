package testing

import (
	"fmt"
	"testing"

	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/utils"
)

func TestEncode(t *testing.T) {
	mconn := utils.SetConnection("mongoenv", "befous")
	geospatial := models.Geospatial{
		Coordinates: []float64{
			103.62074450557095, -1.632735059500547,
		},
		Radius: 0.00003,
	}

	test, err := helpers.GetCenterSphereDoc(mconn, "geojson", geospatial)
	if err != nil {
		fmt.Println(err)
	}
	result := utils.GeojsonNameString(test)
	fmt.Println(result)
}
