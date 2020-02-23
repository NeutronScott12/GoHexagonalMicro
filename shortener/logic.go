package shortener

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)