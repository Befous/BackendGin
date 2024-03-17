package utils

import (
	"database/sql"
	"os"
	"strconv"

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

func SetConnectionPostgres(host, user, password, dbname, port, ssl string) *sql.DB {
	xxx, _ := strconv.Atoi(os.Getenv(port))
	var DBpostgresinfo = models.PostgresInfo{
		Host:     os.Getenv(host),
		User:     os.Getenv(user),
		Password: os.Getenv(password),
		DBName:   os.Getenv(dbname),
		Port:     xxx,
		SSL:      ssl,
	}
	return helpers.PostgresConnect(DBpostgresinfo)
}

func SetConnection2dsphere(mongoenv, dbname, collname string) *mongo.Database {
	var DBmongoinfo = models.DBInfo{
		DBString:       os.Getenv(mongoenv),
		DBName:         dbname,
		CollectionName: collname,
	}
	return helpers.Create2dsphere(DBmongoinfo)
}
