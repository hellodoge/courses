package courses

type Course struct {
	ID          string       `json:"id" bson:"_id"`
	Title       string       `json:"title" bson:"title"`
	Description string       `json:"description,omitempty" bson:"description,omitempty"`
	Lessons     []LessonInfo `json:"lessons" bson:"lessons"`
}

type LessonInfo struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
}

type Lesson struct {
	ID           string     `json:"id" bson:"_id"`
	NextLessonID string     `json:"next,omitempty" bson:"next,omitempty"`
	Title        string     `json:"title,omitempty" bson:"title,omitempty"`
	Description  string     `json:"description,omitempty" bson:"description,omitempty"`
	Photos       []Photo    `json:"photos,omitempty" bson:"photos,omitempty"`
	Videos       []Video    `json:"videos,omitempty" bson:"videos,omitempty"`
	Documents    []Document `json:"documents,omitempty" bson:"documents,omitempty"`
}
