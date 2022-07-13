package helper

import (
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/domain"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryWebResponse {
	return web.CategoryWebResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryWebResponse {

	var categoryResponse []web.CategoryWebResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}
	return categoryResponse
}
