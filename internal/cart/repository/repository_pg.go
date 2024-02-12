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
