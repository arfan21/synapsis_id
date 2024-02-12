package cartrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/pkg/constant"
	dbpostgres "github.com/arfan21/synapsis_id/pkg/db/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r Repository) Create(ctx context.Context, cart entity.Cart) (err error) {
	query := `
		INSERT INTO carts (customer_id, product_id)
		VALUES ($1, $2)
	`

	_, err = r.db.Exec(ctx, query, cart.CustomerID, cart.ProductID)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLUniqueViolation {
				err = constant.ErrProductAlreadyAddedToCart
			}
		}

		err = fmt.Errorf("cart.repository.Create: failed to create cart: %w", err)
		return
	}

	return
}

func (r Repository) GetByCustomerID(ctx context.Context, customerID string) (data []entity.Cart, err error) {
	query := `
		SELECT 
			c.id, 
			c.customer_id, 
			c.product_id,
			c.created_at,
			p.name,
			p.price,
			p.stok
		FROM carts c
		JOIN products p ON c.product_id = p.id
		WHERE c.customer_id = $1
	`

	rows, err := r.db.Query(ctx, query, customerID)
	if err != nil {
		err = fmt.Errorf("cart.repository.GetByCustomerID: failed to get cart by customer id: %w", err)
		return
	}

	for rows.Next() {
		var cart entity.Cart
		err = rows.Scan(
			&cart.ID,
			&cart.CustomerID,
			&cart.ProductID,
			&cart.CreatedAt,
			&cart.Product.Name,
			&cart.Product.Price,
			&cart.Product.Stok,
		)
		if err != nil {
			err = fmt.Errorf("cart.repository.GetByCustomerID: failed to scan cart: %w", err)
			return
		}

		data = append(data, cart)
	}

	if rows.Err() != nil {
		err = fmt.Errorf("product.repository.GetProducts: failed after scan products: %w", err)
		return
	}

	return
}

func (r Repository) Delete(ctx context.Context, customerID, productID string) (err error) {
	query := `
		DELETE FROM carts
		WHERE customer_id = $1 AND product_id = $2
	`

	_, err = r.db.Exec(ctx, query, customerID, productID)
	if err != nil {
		err = fmt.Errorf("cart.repository.Delete: failed to delete cart: %w", err)
		return
	}

	return
}
