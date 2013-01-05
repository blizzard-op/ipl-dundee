package types

type Config struct {
	Port        string
	Elementals  []Elemental
	Streams_url string
}

type Elemental struct {
	Name     string
	Url      string
	Authkey  string
	Authuser string
}
