package productrepo

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/pkg/constant"
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

func (r Repository) queryRowsProductWithFilter(ctx context.Context, query string, filter model.GetListProductRequest) (rows pgx.Rows, err error) {
	var filterArgs []any
	var whereQuery string

	if len(filter.Name) != 0 {
		filterName := "%" + strings.ToLower(filter.Name) + "%"
		filterArgs = append(filterArgs, filterName)
		whereQuery += "LOWER(p.name) LIKE $" + strconv.Itoa(len(filterArgs)) + " AND "
	}

	if filter.CategoryID.Valid {
		filterArgs = append(filterArgs, filter.CategoryID)
		whereQuery += "p.category_id = $" + strconv.Itoa(len(filterArgs)) + " AND "
	}

	// if filterArgsLen  > 0, add WHERE statement and remove last AND
	if filterArgsLen := len(filterArgs); filterArgsLen > 0 {
		whereQuery = "WHERE " + whereQuery[:len(whereQuery)-len(" AND ")] + " "
	}

	query += whereQuery

	if !filter.DisableOffset {
		filterArgs = append(filterArgs, filter.Limit)
		query += "LIMIT $" + strconv.Itoa(len(filterArgs)) + " "

		offset := (filter.Page - 1) * filter.Limit
		filterArgs = append(filterArgs, offset)
		query += "OFFSET $" + strconv.Itoa(len(filterArgs)) + " "
	}

	return r.db.Query(ctx, query, filterArgs...)
}

func (r Repository) GetProducts(ctx context.Context, filter model.GetListProductRequest) (result []entity.Product, err error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.stok,
			p.price,
			pc.id AS category_id,
			pc.name AS category_name,
			c.id AS owner_id,
			c.fullname AS owner_name
		FROM
			products p
			JOIN product_categories pc ON pc.id = p.category_id
			JOIN customers c ON c.id = p.customer_id
	`

	rows, err := r.queryRowsProductWithFilter(ctx, query, filter)
	if err != nil {
		err = fmt.Errorf("product.repository.GetProducts: failed to get products: %w", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var product entity.Product

		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Stok,
			&product.Price,
			&product.Category.ID,
			&product.Category.Name,
			&product.Customer.ID,
			&product.Customer.Fullname,
		)

		if err != nil {
			err = fmt.Errorf("product.repository.GetProducts: failed to scan product: %w", err)
			return
		}

		result = append(result, product)
	}

	if rows.Err() != nil {
		err = fmt.Errorf("product.repository.GetProducts: failed after scan products: %w", err)
		return
	}

	return
}

func (r Repository) GetTotalProduct(ctx context.Context, filter model.GetListProductRequest) (result int, err error) {
	query := `
		SELECT
			COUNT(p.id)
		FROM
			products p
			JOIN product_categories pc ON pc.id = p.category_id
			JOIN customers c ON c.id = p.customer_id
	`
	filter.DisableOffset = true
	rows, err := r.queryRowsProductWithFilter(ctx, query, filter)
	if err != nil {
		err = fmt.Errorf("product.repository.GetTotalProduct: failed to get total product: %w", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&result)
	}

	if rows.Err() != nil {
		err = fmt.Errorf("product.repository.GetProducts: failed after scan total product: %w", err)
		return
	}

	return
}

func (r Repository) IsProductExist(ctx context.Context, id string) (exist bool, err error) {
	query := `
		SELECT EXISTS (SELECT id FROM products WHERE id = $1)
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&exist)
	if err != nil {
		err = fmt.Errorf("product.repository.IsProductExist: failed to check product exist: %w", err)
		return
	}

	return
}

func (r Repository) GetProductByID(ctx context.Context, id string) (result entity.Product, err error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.stok,
			p.price,
			pc.id AS category_id,
			pc.name AS category_name,
			c.id AS owner_id,
			c.fullname AS owner_name
		FROM
			products p
			JOIN product_categories pc ON pc.id = p.category_id
			JOIN customers c ON c.id = p.customer_id
		WHERE
			p.id = $1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(
		&result.ID,
		&result.Name,
		&result.Stok,
		&result.Price,
		&result.Category.ID,
		&result.Category.Name,
		&result.Customer.ID,
		&result.Customer.Fullname,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrProductNotFound
		}
		err = fmt.Errorf("product.repository.GetProductByID: failed to get product by id: %w", err)
		return
	}

	return
}

func (r Repository) UpdateStok(ctx context.Context, data entity.Product) (err error) {
	query := `
		UPDATE products
		SET stok = $1, updated_at = NOW()
		WHERE id = $2 AND stok >= (stok - $1)
	`

	cmd, err := r.db.Exec(ctx, query, data.Stok, data.ID)
	if err != nil {
		err = fmt.Errorf("product.repository.BatchUpdateStok: failed to batch deduct stok: %w", err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = fmt.Errorf("product.repository.BatchUpdateStok: nothing updated: %w", constant.ErrProductNotFoundOrStok)
		return err
	}

	return
}
