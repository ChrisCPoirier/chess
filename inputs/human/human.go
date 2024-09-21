package human

import "fmt"

func New(name string) Human {
	return Human{
		name: name,
	}
}

type Human struct {
	name string
}

func (h Human) Ask(current string) (string, error) {
	var input string

	// Read a single line of input
	fmt.Printf(`%s - what is your move: `, h.name)
	_, err := fmt.Scan(&input)

	if err != nil {
		return ``, err
	}

	return input, nil
}

func (h Human) Name() string {
	return h.name
}
