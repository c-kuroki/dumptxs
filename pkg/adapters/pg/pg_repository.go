package pg

import (
	"context"

	"github.com/c-kuroki/dumptxs/pkg/model"
	"github.com/go-pg/pg/v10/orm"
)

const repName = "db"

type PgRepository struct {
	db orm.DB
}

func NewPgRepository(db orm.DB) *PgRepository {
	return &PgRepository{db: db}
}

func (rep *PgRepository) CreateBlock(ctx context.Context, blockDump *model.BlockDump) error {
	_, err := rep.db.ModelContext(ctx, blockDump).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (rep *PgRepository) Name() string {
	return repName
}
