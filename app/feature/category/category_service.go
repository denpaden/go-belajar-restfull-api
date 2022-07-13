package category

import (
	"context"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/web"
)

type CategoryService interface {
	Save(ctx context.Context, request web.CategoryWebCreateRequest) web.CategoryWebResponse
	Update(ctx context.Context, request web.CategoryWebUpdateRequest) web.CategoryWebResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryWebResponse
	FindAll(ctx context.Context) []web.CategoryWebResponse
}
