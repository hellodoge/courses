package service

type Error struct {
	error      error
	userError  string
	logMessage string
}

func (e Error) Error() string {
	if e.error == nil {
		return "No internal error"
	}
	return e.error.Error()
}

func (e Error) IsSystemError() bool {
	return e.error != nil
}

func (e Error) Log() string {
	return e.logMessage + " (" + e.Error() + ")"
}

func (e Error) HasUserError() bool {
	return e.userError != ""
}

func (e Error) UserError() string {
	return e.userError
}