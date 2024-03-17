package testing

import (
	"fmt"
	"testing"

	"github.com/Befous/BackendGin/models"
	"github.com/Befous/BackendGin/utils"
)

func TestAsss(t *testing.T) {
	mconn := utils.SetConnection("mongoenv", "befous")

	comics := models.Comics{
		ID:                1,
		Title:             "Title",
		Alternative_Title: "AltTitle",
		Description:       "Description",
		Type:              "Manga",
		Genres:            []string{"GL", "Adventure"},
		Tags:              []string{"Magic", "Sword"},
		Rating:            1.0,
		Review:            "asdwdas",
	}
	utils.InsertComics(mconn, "comics", comics)
}

func TestAssss(t *testing.T) {
	mconn := utils.SetConnection("mongoenv", "befous")

	comics := models.Comics{
		Type:   "Manga",
		Genres: []string{"GL", "Adventure"},
		Tags:   []string{"Magic", "Sword"},
	}
	a, err := utils.GetAllFilteredComics(mconn, "comics", comics)
	fmt.Print(a)
	fmt.Print(err)
}
