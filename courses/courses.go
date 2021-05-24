package courses

type Course struct {
	ID          string         `json:"id"`
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description,omitempty"`
	Preview     *URLCollection `json:"photo,omitempty"`
	Lessons     []Lesson       `json:"lessons"`
}

type Lesson struct {
	ID           string     `json:"id"`
	NextLessonID string     `json:"next,omitempty"`
	Title        string     `json:"title,omitempty"`
	Description  string     `json:"description,omitempty"`
	Documents    []Document `json:"documents,omitempty"`
}
