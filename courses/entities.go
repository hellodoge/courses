package courses

type Video struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	URLCollection
}

type Photo struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	URLCollection
}

type Document struct {
	URLCollection
}
