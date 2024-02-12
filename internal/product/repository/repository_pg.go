package productrepo

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/entity"
	dbpostgres "github.com/arfan21/synapsis_id/pkg/db/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db    dbpostgres.Queryer
	rawDb *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:    db,
		rawDb: db,
	}
}

func (r Repository) Begin(ctx context.Context) (tx pgx.Tx, err error) {
	return r.rawDb.Begin(ctx)
}

func (r Repository) WithTx(tx pgx.Tx) *Repository {
	r.db = tx
	return &r
}

func (r Repository) Create(ctx context.Context, data entity.Product) (err error) {
	query := `
		INSERT INTO products (customer_id, category_id, name, stok, price)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = r.db.Exec(ctx, query,
		data.CustomerID,
		data.CategoryID,
		data.Name,
		data.Stok,
		data.Price,
	)

	if err != nil {
		err = fmt.Errorf("product.repository.Create: failed to create product: %w", err)
		return
	}

	return
}

func (r Repository) IsCategoryExist(ctx context.Context, id string) (exist bool, err error) {
	query := `
		SELECT EXISTS (SELECT id FROM product_categories WHERE id = $1)
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&exist)
	if err != nil {
		err = fmt.Errorf("product.repository.IsCategoryExist: failed to check category exist: %w", err)
		return
	}

	return
}

func (r Repository) GetCategories(ctx context.Context) (result []entity.ProductCategory, err error) {
	query := `
		SELECT id, name FROM product_categories
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		err = fmt.Errorf("product.repository.GetCategories: failed to get categories: %w", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var category entity.ProductCategory

		err = rows.Scan(
			&category.ID,
			&category.Name,
		)

		if err != nil {
			err = fmt.Errorf("product.repository.GetCategories: failed to scan category: %w", err)
			return
		}

		result = append(result, category)
	}

	if rows.Err() != nil {
		err = fmt.Errorf("product.repository.GetCategories: failed after scan category: %w", err)
		return
	}

	return
}
