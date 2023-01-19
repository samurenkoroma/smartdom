package dao

type Book struct {
	Id    string `bson:"_id,omitempty"`
	Title string `bson:"title,omitempty"`
	Path  string `bson:"path,omitempty"`
}
