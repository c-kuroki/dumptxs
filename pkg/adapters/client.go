package adapters

import (
	"context"

	"github.com/c-kuroki/dumptxs/pkg/model"
)

type Client interface {
	Block(ctx context.Context, block *int64) (*model.BlockDump, error)
}
