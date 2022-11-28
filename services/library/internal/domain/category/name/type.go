package name

type Name string

func (n Name) String() string {
	return string(n)
}

func New(name string) (*Name, error) {
	n := Name(name)
	return &n, nil
}
