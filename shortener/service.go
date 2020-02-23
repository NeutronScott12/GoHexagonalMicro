package shortener

type RedirectService interface {
	Find(code string) (*Redirec, error)
	Store(redirect *Redirect) error
}