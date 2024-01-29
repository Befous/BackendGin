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
	coordinates := models.Point{
		Coordinates: []float64{
			103.62074450557095, -1.632735059500547,
		},
	}

	test := helpers.GetCenterDoc(mconn, "geojson", coordinates)
	fmt.Println(test)
}
