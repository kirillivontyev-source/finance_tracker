package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kirillivontyev/finance_tracker/internal/models"
)

type CategoryRepository struct {
	Conn *pgx.Conn
}

func NewCategoryRepositoty(conn *pgx.Conn) *CategoryRepository {
	return &CategoryRepository{
		Conn: conn,
	}
}

func (cr *CategoryRepository) GetCategory(ctx context.Context) ([]models.Category, error) {

	sql_Query := `SELECT * FROM categories`

	query, err := cr.Conn.Query(ctx, sql_Query)
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}

	return pgx.CollectRows(query, pgx.RowToStructByName[models.Category])

}

func (cr *CategoryRepository) CreateCategory(ctx context.Context, category models.Category) error {

	sql := `INSERT INTO categories (category_name, monthly_limit) VALUES ($1, $2)`

	comandTag, err := cr.Conn.Exec(ctx, sql, category.CategoryName, category.MonthlyLimit)
	if err != nil {
		return err
	}
	if comandTag.RowsAffected() < 1 {
		return fmt.Errorf("ни одной строки не добавлено")
	}
	return nil
}

func (cr *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {

	sql := `DELETE FROM categories WHERE categories.id = $1`

	commandTag, err := cr.Conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return fmt.Errorf("команда delete не затрагивает ни одной строки")
	}

	return nil

}

func (cr *CategoryRepository) UpdateCategoryName(ctx context.Context, id int, newName string) error {

	sql := `UPDATE categories SET category_name = $1 WHERE id = $2`

	commandTag, err := cr.Conn.Exec(ctx, sql, newName, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return fmt.Errorf("команда update не затрагивает ни одной строки")
	}

	return nil

}
