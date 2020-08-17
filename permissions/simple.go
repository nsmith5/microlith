package permissions

type simple int

func (s *simple) CanI(thing string) bool {
	_, ok := okthings[thing]
	return ok
}

func NewSimple() Service {
	return new(simple)
}
