package utils

import (
	"os"

	"github.com/Befous/BackendGin/helpers"
	"github.com/Befous/BackendGin/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = models.DBInfo{
		DBString: os.Getenv(mongoenv),
		DBName:   dbname,
	}
	return helpers.MongoConnect(DBmongoinfo)
}

func SetConnection2dsphere(mongoenv, dbname, collname string) *mongo.Database {
	var DBmongoinfo = models.DBInfo{
		DBString:       os.Getenv(mongoenv),
		DBName:         dbname,
		CollectionName: collname,
	}
	return helpers.Create2dsphere(DBmongoinfo)
}
