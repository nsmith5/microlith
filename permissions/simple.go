package permissions

type simple int

func (s *simple) CanI(thing string) (bool, error) {
	_, ok := okthings[thing]
	return ok, nil
}

func NewSimple() Service {
	return new(simple)
}
