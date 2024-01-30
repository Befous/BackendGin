package testing

import (
	"fmt"
	"testing"

	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/utils"
)

func TestBox(t *testing.T) {
	mconn := utils.SetConnection("mongoenv", "befous")
	geospatial := models.Geospatial{
		Coordinates: [][]float64{
			{103.54786554087926, -1.6487359545296698},
			{103.62772892345659, -1.6034217927083034},
		},
	}
	box := utils.Box(mconn, "geojson", geospatial)
	fmt.Println(box)
}
