package helpers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

func GetGeoIntersectsDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
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

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Geo Intersects: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang bersinggungan dengan koordinat anda adalah: " + result
}

func GetGeoWithinDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
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

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GeoWithin: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang berada di dalam koordinat anda adalah: " + result
}

func GetNearDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
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

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Near: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang berdekatan dengan koordinat anda adalah: " + result
}

func GetNearSphereDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
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

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Near Sphere: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang berdekatan dengan koordinat anda adalah: " + result
}

func GetBoxDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$box": geospatial.Coordinates,
			},
		},
	}

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Box: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang berada di dalam box anda adalah: " + result
}

func GetCenterDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$center": []interface{}{geospatial.Coordinates, geospatial.Radius},
			},
		},
	}

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Center: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang berada di dalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius, 'f', -1, 64) + " adalah: " + result
}

func GetCenterSphereDoc(db *mongo.Database, collname string, geospatial models.Geospatial) (result string) {
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{geospatial.Coordinates, geospatial.Radius},
			},
		},
	}

	var docs []models.FullGeoJson
	cur, err := db.Collection(collname).Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Center Sphere: %v\n", err)
		return ""
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var doc models.FullGeoJson
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Printf("Decode Err: %v\n", err)
			continue
		}
		docs = append(docs, doc)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Cursor Err: %v\n", err)
		return ""
	}

	// Ambil nilai properti Name dari setiap dokumen
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties.Name)
	}

	// Gabungkan nilai-nilai dengan koma
	result = strings.Join(names, ", ")

	return "Geojson yang berada di dalam lingkaran dengan radius " + strconv.FormatFloat(geospatial.Radius, 'f', -1, 64) + " adalah: " + result
}
