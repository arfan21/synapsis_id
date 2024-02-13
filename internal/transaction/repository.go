package transaction

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/entity"
	transactionrepo "github.com/arfan21/synapsis_id/internal/transaction/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Begin(ctx context.Context) (tx pgx.Tx, err error)
	WithTx(tx pgx.Tx) *transactionrepo.Repository

	Create(ctx context.Context, data entity.Transaction) (id uuid.UUID, err error)
	CreateDetail(ctx context.Context, data []entity.TransactionDetail) (err error)
}
