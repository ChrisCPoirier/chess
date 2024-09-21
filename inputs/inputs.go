package inputs

type Player interface {
	Ask(string) (string, error)
	Name() string
}
