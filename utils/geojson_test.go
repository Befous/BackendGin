package utils

import (
	"fmt"
	"testing"

	"github.com/Befous/BackendGin/models"
)

func TestBox(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	coordinates := models.Polyline{
		Coordinates: [][]float64{
			{103.54786554087926, -1.6487359545296698},
			{103.62772892345659, -1.6034217927083034},
		},
	}
	box := Box(mconn, "geojson", coordinates)
	fmt.Println(box)
}
