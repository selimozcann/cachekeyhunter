package types

type Signals struct {
	StatusCode int
	Headers    map[string]string
	BodyHash   string
	Age        int
	Cache      string
}

type Variant struct {
	Name    string
	Headers map[string]string
	Query   map[string]string
}

type Finding struct {
	URL      string
	Severity string
	Detail   string
	Evidence string
}
