package utils

import (
	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertComics(mongoenv *mongo.Database, collname string, datauser models.Comics) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datauser)
}

func GetAllComics(mconn *mongo.Database, collname string) ([]models.Comics, error) {
	return helpers.GetAllDoc[models.Comics](mconn, collname)
}

func GetAllFilteredComics(mconn *mongo.Database, collname string, datacomics models.Comics) ([]models.Comics, error) {
	filter := bson.M{
		"type": datacomics.Type,
		"genres": bson.M{
			"$all": datacomics.Genres,
		},
		"tags": bson.M{
			"$all": datacomics.Tags,
		},
	}
	return helpers.GetAllDocByFilter[models.Comics](mconn, collname, filter)
}

func GetOneComics(mconn *mongo.Database, collname string, userdata models.Comics) models.Comics {
	filter := bson.M{"id": userdata.ID}
	return helpers.GetOneDoc[models.Comics](mconn, collname, filter)
}

func UpdateComics(mconn *mongo.Database, collname string, datauser models.Comics) interface{} {
	filter := bson.M{"id": datauser.ID}
	return helpers.ReplaceOneDoc(mconn, collname, filter, datauser)
}

func DeleteComics(mconn *mongo.Database, collname string, userdata models.Comics) interface{} {
	filter := bson.M{"id": userdata.ID}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
