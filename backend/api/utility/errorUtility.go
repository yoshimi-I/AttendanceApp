package utility

// AuthenticationError 401エラー
type AuthenticationError struct {
	Message string
}

func (e AuthenticationError) Error() string {
	return e.Message
}

// NotFoundError 404エラー
type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

// BadRequestError 400エラー
type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

// ForbiddenError 403エラー
type ForbiddenError struct {
	Message string
}

func (e ForbiddenError) Error() string {
	return e.Message
}

// TODO こっから追加していく予定
