package file

import (
	"context"
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

	for _, d := range blockDump.Dump {
		_, err = f.WriteString(d)
		if err != nil {
			return err
		}
	}
	_ = f.Sync()
	return nil
}

func (rep *FileRepository) Name() string {
	return repName
}
