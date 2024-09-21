package inputs

type Player interface {
	Ask(string, []string) (string, error)
	Name() string
}
