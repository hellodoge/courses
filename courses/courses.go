package courses

type Course struct {
	ID          string   `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description,omitempty"`
	Preview     *Photo   `json:"photo,omitempty"`
	Lessons     []Lesson `json:"lessons"`
}

type Lesson struct {
	ID           string     `json:"id"`
	NextLessonID string     `json:"next,omitempty"`
	Title        string     `json:"title,omitempty"`
	Description  string     `json:"description,omitempty"`
	Photos       []Photo    `json:"photos,omitempty"`
	Videos       []Video    `json:"videos,omitempty"`
	Documents    []Document `json:"documents,omitempty"`
}
