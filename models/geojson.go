package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Properties struct {
	Name string `json:"name" bson:"name"`
}

type Geometry struct {
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeoJson struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   Geometry   `json:"geometry" bson:"geometry"`
}

type FullGeoJson struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   Geometry           `json:"geometry" bson:"geometry"`
}

type Point struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Max         float64   `json:"max,omitempty" bson:"max,omitempty"`
	Min         float64   `json:"min,omitempty" bson:"min,omitempty"`
}

type Polyline struct {
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
}

type Polygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
}

type GeoJsonPoint struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   struct {
		Coordinates []float64 `json:"coordinates" bson:"coordinates"`
		Type        string    `json:"type" bson:"type"`
	} `json:"geometry" bson:"geometry"`
}

type GeoJsonLineString struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   struct {
		Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
		Type        string      `json:"type" bson:"type"`
	} `json:"geometry" bson:"geometry"`
}

type GeoJsonPolygon struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   struct {
		Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
		Type        string        `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"geometry" bson:"geometry"`
}
