package utils

import (
	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create

func PostPoint(mconn *mongo.Database, collection string, pointdata models.GeoJsonPoint) interface{} {
	return helpers.InsertOneDoc(mconn, collection, pointdata)
}

func PostLinestring(mconn *mongo.Database, collection string, linestringdata models.GeoJsonLineString) interface{} {
	return helpers.InsertOneDoc(mconn, collection, linestringdata)
}

func PostPolygon(mconn *mongo.Database, collection string, polygondata models.GeoJsonPolygon) interface{} {
	return helpers.InsertOneDoc(mconn, collection, polygondata)
}

// Read

func GetAllBangunan(mconn *mongo.Database, collname string) []models.GeoJson {
	return helpers.GetAllDoc[[]models.GeoJson](mconn, collname)
}

func GeoIntersects(mconn *mongo.Database, collname string, coordinates models.Point) string {
	return helpers.GetGeoIntersectsDoc(mconn, collname, coordinates)
}

func GeoWithin(mconn *mongo.Database, collname string, coordinates models.Polygon) string {
	return helpers.GetGeoWithinDoc(mconn, collname, coordinates)
}

func Near(mconn *mongo.Database, collname string, coordinates models.Point) string {
	return helpers.GetNearDoc(mconn, collname, coordinates)
}

func NearSphere(mconn *mongo.Database, collname string, coordinates models.Point) string {
	return helpers.GetNearSphereDoc(mconn, collname, coordinates)
}

func Box(mconn *mongo.Database, collname string, coordinates models.Polyline) string {
	return helpers.GetBoxDoc(mconn, collname, coordinates)
}

func Center(mconn *mongo.Database, collname string, coordinates models.Point) string {
	return helpers.GetCenterDoc(mconn, collname, coordinates)
}

func CenterSphere(mconn *mongo.Database, collname string, coordinates models.Point) string {
	return helpers.GetCenterSphereDoc(mconn, collname, coordinates)
}

// Update

// Delete

func DeleteGeojson(mconn *mongo.Database, collname string, userdata models.User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
