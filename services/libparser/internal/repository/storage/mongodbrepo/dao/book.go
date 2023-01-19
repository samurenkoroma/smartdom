package dao

type Book struct {
	ID      string   `bson:"_id,omitempty"`
	Title   string   `bson:"title,omitempty"`
	Path    string   `bson:"path,omitempty"`
	Tags    []string `bson:"tags,omitempty"`
	Authors []string `bson:"authors,omitempty"`
	Size    uint64   `bson:"size,omitempty"`
}
