package customErr

type UserAuthenticationError struct {
	Message string
}

func (e UserAuthenticationError) Error() string {
	return e.Message
}

type ActivityNotFoundError struct {
	Message string
}

func (e ActivityNotFoundError) Error() string {
	return e.Message
}

// TODO こっから追加していく予定
