package path

type Path string

func (t Path) String() string {
	return string(t)
}

func New(path string) (*Path, error) {
	n := Path(path)
	return &n, nil
}
