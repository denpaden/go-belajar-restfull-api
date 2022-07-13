package web

type CategoryWebUpdateRequest struct {
	Id   int
	Name string `validate:"required"`
}
