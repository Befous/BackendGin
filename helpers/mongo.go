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

	// Mengecek apakah indeks sudah ada
	collection := client.Database(mconn.DBName).Collection(mconn.CollectionName)
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		fmt.Printf("Error listing indexes: %v", err)
	}

	indexExists := false
	for cursor.Next(context.TODO()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			fmt.Printf("Error decoding index: %v", err)
		}
		if index["name"] == "geometry_2dsphere" {
			indexExists = true
			break
		}
	}

	// Membuat indeks jika belum ada
	if !indexExists {
		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "geometry", Value: "2dsphere"},
			},
		}

		_, err = client.Database(mconn.DBName).Collection(mconn.CollectionName).Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			fmt.Printf("Error creating geospatial index: %v", err)
		}
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

func GetGeoIntersectsDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        geospatial.Type,
					"coordinates": geospatial.Coordinates,
				},
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetGeoIntersectsDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetGeoIntersectsDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetGeoWithinDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        geospatial.Type,
					"coordinates": geospatial.Coordinates,
				},
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetGeoWithinDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetGeoWithinDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetNearDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        geospatial.Type,
					"coordinates": geospatial.Coordinates,
				},
				"$maxDistance": geospatial.Max,
				"$minDistance": geospatial.Min,
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetNearDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetNearDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetNearSphereDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        geospatial.Type,
					"coordinates": geospatial.Coordinates,
				},
				"$maxDistance": geospatial.Max,
				"$minDistance": geospatial.Min,
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetNearSphereDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetNearSphereDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetBoxDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$box": geospatial.Coordinates,
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetBoxDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetBoxDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetCenterDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$center": []interface{}{geospatial.Coordinates, geospatial.Radius},
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetCenterDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetCenterDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetCenterSphereDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result []models.FullGeoJson, err error) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{geospatial.Coordinates, geospatial.Radius},
			},
		},
	}

	ctx := context.TODO()
	cur, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetCenterSphereDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &result)
	if err != nil {
		fmt.Printf("GetCenterSphereDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return result, nil
}
