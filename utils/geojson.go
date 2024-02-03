package utils

import (
	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create

func PostPoint(mconn *mongo.Database, collection string, pointdata models.GeoJsonPoint) (interface{}, error) {
	return helpers.InsertOneDoc(mconn, collection, pointdata)
}

func PostLinestring(mconn *mongo.Database, collection string, linestringdata models.GeoJsonLineString) (interface{}, error) {
	return helpers.InsertOneDoc(mconn, collection, linestringdata)
}

func PostPolygon(mconn *mongo.Database, collection string, polygondata models.GeoJsonPolygon) (interface{}, error) {
	return helpers.InsertOneDoc(mconn, collection, polygondata)
}

// Read

func GetAllBangunan(mconn *mongo.Database, collname string) ([]models.GeoJson, error) {
	return helpers.GetAllDoc[models.GeoJson](mconn, collname)
}

func GeoIntersects(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetGeoIntersectsDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

func GeoWithin(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetGeoWithinDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

func Near(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetNearDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

func NearSphere(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetNearSphereDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

func Box(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetBoxDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

func Center(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetCenterDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

func CenterSphere(mconn *mongo.Database, collname string, geospatial models.Geospatial) ([]models.FullGeoJson, error) {
	return helpers.GetCenterSphereDoc[models.FullGeoJson](mconn, collname, "geometry", geospatial)
}

// Update

// Delete

func DeleteGeojson(mconn *mongo.Database, collname string, userdata models.User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
