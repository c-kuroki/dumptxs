package file

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/c-kuroki/dumptxs/pkg/model"
)

const repName = "file"

type FileRepository struct {
	path string
}

func NewFileRepository(filePath string) *FileRepository {
	return &FileRepository{path: filePath}
}

func (rep *FileRepository) CreateBlock(ctx context.Context, blockDump *model.BlockDump) error {
	f, err := os.Create(fmt.Sprintf("%s/%d.json", rep.path, blockDump.BlockNumber))
	if err != nil {
		return err
	}
	defer f.Close()

	bdBytes, err := json.Marshal(blockDump)
	if err != nil {
		return err
	}
	_, err = f.Write(bdBytes)
	if err != nil {
		return err
	}
	_ = f.Sync()
	return nil
}

func (rep *FileRepository) Name() string {
	return repName
}
