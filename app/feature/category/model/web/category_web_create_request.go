package web

type CategoryWebCreateRequest struct {
	Name string `validate:"required"`
}
