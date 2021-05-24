package callback

const (
	ActionGetCourseDescription = "get_course_description"
	ActionGetCourseLessons     = "get_course_lessons"
	ActionGetLesson            = "get_lesson"
)

type Query struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}
