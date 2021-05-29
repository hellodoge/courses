package callback

const (
	ActionGetCourseDescription = "course_description"
	ActionGetCourseLessons     = "course_lessons"
	ActionGetLesson            = "lesson"
	ActionSearch               = "search"
)

type Query struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}
