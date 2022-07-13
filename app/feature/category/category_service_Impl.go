package category

import (
	"context"
	"database/sql"
	"github.com/denpaden/go-belajar-restfull-api/app/excecption"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/domain"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/web"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository CategoryRepository
	DB                 *sql.DB
	Validation         *validator.Validate
}

func logStart(msg string) string {
	return "CategoryServiceImpl Starting for " + msg
}
func logFinish(msg string) string {
	return "CategoryServiceImpl Finished for " + msg
}
func NewCategoryService(categoryRepository CategoryRepository, DB *sql.DB, validation *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validation: validation}
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request web.CategoryWebCreateRequest) web.CategoryWebResponse {
	helper.LoggerInfo(logStart("Save"))
	err := service.Validation.Struct(request)
	helper.PanicIfError(err)

	// init db & transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// define body
	domain := domain.Category{
		Id:   0,
		Name: request.Name,
	}

	// call repo function
	domain = service.CategoryRepository.Save(ctx, tx, domain)

	// return & convert to response
	helper.LoggerInfo(logFinish("Save"))
	return helper.ToCategoryResponse(domain)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryWebUpdateRequest) web.CategoryWebResponse {
	helper.LoggerInfo(logStart("Update"))
	err := service.Validation.Struct(request)
	helper.PanicIfError(err)

	// init db & transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find by id
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(excecption.NewNotFoundError(err.Error()))
	}

	// update
	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	// response
	helper.LoggerInfo(logFinish("Update"))
	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	helper.LoggerInfo(logStart("Delete"))
	// init db & transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find by id
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(excecption.NewNotFoundError(err.Error()))
	}

	// delete
	helper.LoggerInfo(logFinish("Delete"))
	service.CategoryRepository.Delete(ctx, tx, category.Id)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryWebResponse {
	helper.LoggerInfo(logStart("FindById"))
	// init db & transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find by id
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(excecption.NewNotFoundError(err.Error()))
	}

	helper.LoggerInfo(logFinish("FindById"))
	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryWebResponse {
	helper.LoggerInfo(logStart("FindAll"))
	// init db & transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find by id
	categories := service.CategoryRepository.FindAll(ctx, tx)

	helper.LoggerInfo(logFinish("FindAll"))
	return helper.ToCategoryResponses(categories)

}
