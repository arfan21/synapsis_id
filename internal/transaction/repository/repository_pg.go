package transactionrepo

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/pkg/constant"
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
	columns := []string{"transaction_id", "product_id", "qty"}

	rows := make([][]interface{}, len(data))
	for i, item := range data {
		rows[i] = []interface{}{item.TransactionID, item.ProductID, item.Qty}
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
		err = fmt.Errorf("transaction.repository.CreateDetail: failed to create transaction detail: %w", constant.ErrTxDetailInsertedNotEqual)
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
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`

	_, err = r.db.Exec(ctx, query, status, id)
	if err != nil {
		err = fmt.Errorf("transaction.repository.UpdateStatus: failed to update status: %w", err)
		return
	}

	return
}

func (r Repository) GetByCustomerID(ctx context.Context, customerID string) (result []entity.Transaction, err error) {
	query := `
		SELECT 
			tx.id, 
			tx.customer_id, 
			tx.payment_method_id, 
			tx.total_amount, 
			tx.status, 
			tx.created_at,
			tx.updated_at,
			pm.name AS payment_method_name,
			td.id AS detail_id,
			td.transaction_id AS detail_transaction_id,
			td.qty AS detail_qty,
			p.id AS product_id,
			p.name AS product_name,
			p.price AS product_price
		FROM transactions tx
			JOIN payment_methods pm ON pm.id = tx.payment_method_id
			JOIN transaction_details td ON td.transaction_id = tx.id
			JOIN products p ON p.id = td.product_id
		WHERE tx.customer_id = $1
	`

	rows, err := r.db.Query(ctx, query, customerID)
	if err != nil {
		err = fmt.Errorf("transaction.repository.GetByCustomerID: failed to get transaction by customer id: %w", err)
		return
	}

	defer rows.Close()

	mapData := make(map[uuid.UUID]entity.Transaction)
	for rows.Next() {
		var tx entity.Transaction
		var detail entity.TransactionDetail
		err = rows.Scan(
			&tx.ID,
			&tx.CustomerID,
			&tx.PaymentMethodID,
			&tx.TotalAmount,
			&tx.Status,
			&tx.CreatedAt,
			&tx.UpdatedAt,
			&tx.PaymentMethod.Name,
			&detail.ID,
			&detail.TransactionID,
			&detail.Qty,
			&detail.ProductID,
			&detail.Product.Name,
			&detail.Product.Price,
		)

		if err != nil {
			err = fmt.Errorf("transaction.repository.GetByCustomerID: failed to scan data: %w", err)
			return
		}

		if currentData, ok := mapData[tx.ID]; !ok {
			tx.TransactionDetails = append(tx.TransactionDetails, detail)
			mapData[tx.ID] = tx
		} else {
			currentData.TransactionDetails = append(currentData.TransactionDetails, detail)
			mapData[tx.ID] = currentData
		}

	}

	if err = rows.Err(); err != nil {
		err = fmt.Errorf("transaction.repository.GetByCustomerID: failed to get transaction by customer id: %w", err)
		return
	}

	for _, item := range mapData {
		result = append(result, item)
	}

	return
}
