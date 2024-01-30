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
	}

	test := helpers.GetCenterDoc(mconn, "geojson", geospatial)
	fmt.Println(test)
}
