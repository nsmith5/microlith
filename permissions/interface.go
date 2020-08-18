package permissions

var okthings = map[string]struct{}{
	"walk": struct{}{},
	"talk": struct{}{},

	// Nah, we're not that cool
	// "walk the walk": struct{}
}

type Service interface {
	CanI(string) (bool, error)
}
