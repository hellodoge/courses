package courses

const (
	TypeVideo    = "video"
	TypePhoto    = "photo"
	TypeDocument = "document"
)

type Document struct {
	Type        string        `json:"type" binding:"required" bson:"type"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	URLs        URLCollection `json:"url_collection" bson:"url_collection"`
}
