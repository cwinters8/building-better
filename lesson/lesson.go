package lesson

type Lesson struct {
	Name string
}

func New(name string) *Lesson {
	return &Lesson{
		Name: name,
	}
}
