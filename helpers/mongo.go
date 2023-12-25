package helpers

import (
	"context"
	"fmt"

	"github.com/Befous/BackendGin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn models.DBInfo) (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI((mconn.DBString))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}
	return client.Database(mconn.DBName)
}

func Create2dsphere(mconn models.DBInfo) (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI((mconn.DBString))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "geometry", Value: "2dsphere"},
		},
	}

	_, err = client.Database(mconn.DBName).Collection(mconn.CollectionName).Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Printf("Error creating geospatial index: %v", err)
	}

	return client.Database(mconn.DBName)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("AIteung Mongo, InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T) {
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GetOneDoc: %v\n", err)
	}
	return
}

func GetOneLatestDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	err = db.Collection(collection).FindOne(context.TODO(), filter, opts).Decode(&doc)
	if err != nil {
		return
	}
	return
}

func GetAllDocByFilter[T any](db *mongo.Database, collection string, filter bson.M) (doc T) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
	}
	return
}

func GetAllDoc[T any](db *mongo.Database, collection string) (doc T) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
	}
	return
}

func GetAllDistinctDoc(db *mongo.Database, filter bson.M, fieldname, collection string) (doc []any) {
	ctx := context.TODO()
	doc, err := db.Collection(collection).Distinct(ctx, fieldname, filter)
	if err != nil {
		fmt.Printf("GetAllDistinctDoc: %v\n", err)
	}
	return
}

func ReplaceOneDoc(db *mongo.Database, collection string, filter bson.M, doc interface{}) (updatereseult *mongo.UpdateResult) {
	updatereseult, err := db.Collection(collection).ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		fmt.Printf("ReplaceOneDoc: %v\n", err)
	}
	return
}

func DeleteOneDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteOneDoc: %v\n", err)
	}
	return
}

func DeleteDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDoc : %v\n", err)
	}
	return
}

func GetRandomDoc[T any](db *mongo.Database, collection string, size uint) (result []T, err error) {
	filter := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: size}}}},
	}
	ctx := context.Background()
	cursor, err := db.Collection(collection).Aggregate(ctx, filter)
	if err != nil {
		return
	}

	err = cursor.All(ctx, &result)

	return
}

func DocExists[T any](db *mongo.Database, collname string, filter bson.M, doc T) (result bool) {
	err := db.Collection(collname).FindOne(context.Background(), filter).Decode(&doc)
	return err == nil
}

func GetGeoIntersectsDoc(db *mongo.Database, collname string, coordinates models.Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
			},
		},
	}
	var doc models.FullGeoJson
	err := db.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GeoIntersects: %v\n", err)
	}
	return "Koordinat anda bersinggungan dengan " + doc.Properties.Name
}

func GetGeoWithinDoc(db *mongo.Database, collname string, coordinates models.Polygon) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": coordinates.Coordinates,
				},
			},
		},
	}
	var doc models.FullGeoJson
	err := db.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GeoWithin: %v\n", err)
	}
	return "Koordinat anda berada di " + doc.Properties.Name
}

func GetNearDoc(db *mongo.Database, collname string, coordinates models.Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
				"$maxDistance": 1000,
			},
		},
	}
	var doc models.FullGeoJson
	err := db.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("Near: %v\n", err)
	}
	return "Koordinat anda dekat dengan " + doc.Properties.Name
}

func GetNearSphereDoc(db *mongo.Database, collname string, coordinates models.Point) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
				"$maxDistance": 1000,
			},
		},
	}
	var doc models.FullGeoJson
	err := db.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("NearSphere: %v\n", err)
	}
	return "Koordinat anda dekat dengan " + doc.Properties.Name
}

func GetBoxDoc(db *mongo.Database, collname string, coordinates models.Polyline) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$box": coordinates.Coordinates,
			},
		},
	}
	var doc models.FullGeoJson
	err := db.Collection(collname).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("Box: %v\n", err)
	}
	return "Box anda berada pada " + doc.Properties.Name
}
