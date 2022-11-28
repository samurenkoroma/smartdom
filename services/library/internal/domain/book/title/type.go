package title

type Title string

func (t Title) String() string {
	return string(t)
}

func New(title string) (*Title, error) {
	n := Title(title)
	return &n, nil
}
