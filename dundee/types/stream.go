package types

type Stream struct {
	Id        string
	Franchise franchise
}

type franchise struct {
	Name string
}
