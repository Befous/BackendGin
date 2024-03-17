package models

type Comics struct {
	ID                int      `json:"id" bson:"id"`
	Title             string   `json:"title" bson:"title"`
	Alternative_Title string   `json:"alternative_title" bson:"alternative_title"`
	Description       string   `json:"description" bson:"description"`
	Type              string   `json:"type" bson:"type"`
	Genres            []string `json:"genres" bson:"genres"`
	Tags              []string `json:"tags" bson:"tags"`
	Rating            float32  `json:"rating,omitempty" bson:"rating,omitempty"`
	Review            string   `json:"review,omitempty" bson:"review,omitempty"`
	Image             string   `json:"image,omitempty" bson:"image,omitempty"`
}
