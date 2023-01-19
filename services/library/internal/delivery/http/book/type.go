package book

type ID struct {
	Value string `json:"id" uri:"id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
}
