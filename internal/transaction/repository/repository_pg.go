package transactionrepo

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/entity"
	dbpostgres "github.com/arfan21/synapsis_id/pkg/db/postgres"
	"github.com/google/uuid"
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

func (r Repository) Create(ctx context.Context, data entity.Transaction) (id uuid.UUID, err error) {
	query := `
		INSERT INTO transactions (customer_id, payment_method_id,  total_amount)
		VALUES ($1, $2, $3) RETURNING id
	`

	err = r.db.QueryRow(ctx, query,
		data.CustomerID,
		data.PaymentMethodID,
		data.TotalAmount,
	).Scan(&id)

	if err != nil {
		err = fmt.Errorf("transaction.repository.Create: failed to create transaction: %w", err)
		return
	}

	return
}

func (r Repository) CreateDetail(ctx context.Context, data []entity.TransactionDetail) (err error) {
	columns := []string{"transaction_id", "product_id"}

	rows := make([][]interface{}, len(data))
	for i, item := range data {
		rows[i] = []interface{}{item.TransactionID, item.ProductID}
	}

	rowsAffected, err := r.db.CopyFrom(ctx,
		pgx.Identifier{entity.TransactionDetail{}.TableName()},
		columns,
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		err = fmt.Errorf("transaction.repository.CreateDetail: failed to create transaction detail: %w", err)
		return
	}

	if rowsAffected != int64(len(data)) {
		err = fmt.Errorf("transaction.repository.CreateDetail: failed to create transaction detail: %w", err)
		return
	}

	return
}

func (r Repository) GetByID(ctx context.Context, id uuid.UUID) (result entity.Transaction, err error) {
	query := `
		SELECT id, customer_id, payment_method_id, total_amount, status, created_at
		FROM transactions
		WHERE id = $1
	`

	err = r.db.QueryRow(ctx, query, id).Scan(
		&result.ID,
		&result.CustomerID,
		&result.PaymentMethodID,
		&result.TotalAmount,
		&result.Status,
		&result.CreatedAt,
	)

	if err != nil {
		err = fmt.Errorf("transaction.repository.GetByID: failed to get transaction by id: %w", err)
		return
	}

	return
}

func (r Repository) UpdateStatus(ctx context.Context, id uuid.UUID, status entity.TransactionStatus) (err error) {
	query := `
		UPDATE transactions
		SET status = $1
		WHERE id = $2
	`

	_, err = r.db.Exec(ctx, query, status, id)
	if err != nil {
		err = fmt.Errorf("transaction.repository.UpdateStatus: failed to update status: %w", err)
		return
	}

	return
}
