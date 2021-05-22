package courses

const (
	TelegramShortcut = "telegram"
)

type URLCollection struct {
	URL       string            `json:"url" bson:"url"`
	Shortcuts map[string]string `json:"shortcuts,omitempty" bson:"shortcuts,omitempty"`
}
