package option

type Setter func(e *ExampleOption)

type ExampleOption struct {
	ID     int
	Client string
}

func NewExampleOption(ss ...Setter) *ExampleOption {
	eo := new(ExampleOption)
	for _, s := range ss {
		s(eo)
	}
	return eo
}

func WithID(i int) Setter {
	return func(e *ExampleOption) {
		e.ID = i
	}
}

func WithClient(c string) Setter {
	return func(e *ExampleOption) {
		e.Client = c
	}
}
