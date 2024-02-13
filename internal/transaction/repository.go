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
	GetByID(ctx context.Context, id uuid.UUID) (result entity.Transaction, err error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status entity.TransactionStatus) (err error)
}
