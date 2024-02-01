package utils

import (
	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(mongoenv *mongo.Database, collname string, datauser models.User) interface{} {
	return helpers.InsertOneDoc(mongoenv, collname, datauser)
}

func GetAllUser(mconn *mongo.Database, collname string) []models.User {
	return helpers.GetAllDoc[[]models.User](mconn, collname)
}

func FindUser(mconn *mongo.Database, collname string, userdata models.User) models.User {
	filter := bson.M{"username": userdata.Username}
	return helpers.GetOneDoc[models.User](mconn, collname, filter)
}

func IsPasswordValid(mconn *mongo.Database, collname string, userdata models.User) bool {
	filter := bson.M{"username": userdata.Username}
	res := helpers.GetOneDoc[models.User](mconn, collname, filter)
	hashChecker := helpers.CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}

func UsernameExists(mconn *mongo.Database, collname string, userdata models.User) bool {
	filter := bson.M{"username": userdata.Username}
	return helpers.DocExists[models.User](mconn, collname, filter, userdata)
}

func UpdateUser(mconn *mongo.Database, collname string, datauser models.User) interface{} {
	filter := bson.M{"username": datauser.Username}
	return helpers.ReplaceOneDoc(mconn, collname, filter, datauser)
}

func DeleteUser(mconn *mongo.Database, collname string, userdata models.User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return helpers.DeleteOneDoc(mconn, collname, filter)
}
