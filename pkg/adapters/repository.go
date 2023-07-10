package adapters

import (
	"context"

	"github.com/c-kuroki/dumptxs/pkg/model"
)

type Repository interface {
	Name() string
	CreateBlock(ctx context.Context, blockDump *model.BlockDump) error
}
