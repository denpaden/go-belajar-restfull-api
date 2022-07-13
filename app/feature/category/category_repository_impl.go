package category

import (
	"context"
	"database/sql"
	"errors"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/domain"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// execute query
	query := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	// dapatkan id terakhir after insert
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	// set id terakhir
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// execute query
	query := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	// execute query
	query := "DELETE from category WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, categoryId)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// execute query
	query := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("data not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// execute query
	query := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)

	}
	return categories
}
