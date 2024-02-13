package paymentrepo

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

func (r Repository) GetPaymentMethods(ctx context.Context) (result []entity.PaymentMethod, err error) {
	query := "SELECT id, name FROM payment_methods"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		err = fmt.Errorf("payment.repository.GetPaymentMethods: failed to get payment methods: %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data entity.PaymentMethod
		err = rows.Scan(&data.ID, &data.Name)
		if err != nil {
			err = fmt.Errorf("payment.repository.GetPaymentMethods: failed to scan payment methods: %w", err)
			return nil, err
		}

		result = append(result, data)
	}

	if err = rows.Err(); err != nil {
		err = fmt.Errorf("payment.repository.GetPaymentMethods: failed to iterate payment methods: %w", err)
		return nil, err
	}

	return result, nil
}

func (r Repository) GetPaymentMethodByID(ctx context.Context, id string) (result entity.PaymentMethod, err error) {
	query := "SELECT id, name FROM payment_methods WHERE id = $1"
	err = r.db.QueryRow(ctx, query, id).Scan(&result.ID, &result.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = constant.ErrPaymentMethodNotFound
		}
		err = fmt.Errorf("payment.repository.GetPaymentMethodByID: failed to get payment method by id: %w", err)
		return result, err
	}

	return result, nil
}

func (r Repository) CreatePayment(ctx context.Context, payment entity.Payment) (err error) {
	query := `
		INSERT INTO payments (transaction_id)
		VALUES ($1)
	`

	_, err = r.db.Exec(ctx, query, payment.TransactionID)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLUniqueViolation {
				err = constant.ErrTransactionAlreadyPaid
			}
		}
		err = fmt.Errorf("payment.repository.CreatePayment: failed to create payment: %w", err)
		return err
	}

	return nil
}
