package courses

type Video struct {
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	URLs        URLCollection `json:"url_collection" bson:"url_collection"`
}

type Photo struct {
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	URLs        URLCollection `json:"url_collection" bson:"url_collection"`
}

type Document struct {
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	URLs        URLCollection `json:"url_collection" bson:"url_collection"`
}
