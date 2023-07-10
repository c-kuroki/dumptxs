package app

import (
	"context"
	"log"

	"github.com/c-kuroki/dumptxs/pkg/adapters"
)

type DumpTxsApp struct {
	client adapters.Client
	repos  []adapters.Repository
	log    *log.Logger
}

func NewDumpTxsApp(client adapters.Client, repos []adapters.Repository, log *log.Logger) *DumpTxsApp {
	return &DumpTxsApp{
		client: client,
		repos:  repos,
		log:    log,
	}
}

func (app *DumpTxsApp) DumpRange(ctx context.Context, from, to int64) error {
	var err error
	for i := from; i <= to; i++ {
		err = app.DumpBlock(ctx, &i)
		if err != nil {
			app.log.Printf("error dumping block# %d : %v\n", i, err.Error())
			return err
		}
	}
	return nil
}

func (app *DumpTxsApp) DumpBlock(ctx context.Context, block *int64) error {
	if block == nil {
		app.log.Println("Pulling current block")
	} else {
		app.log.Printf("Pulling block %d\n", *block)
	}
	// get block
	blockDump, err := app.client.Block(ctx, block)
	if err != nil {
		app.log.Printf("error pulling block : %v\n", err.Error())
		return err
	}
	// traverse repos to store the block dump
	for _, rep := range app.repos {
		err := rep.CreateBlock(ctx, blockDump)
		if err != nil {
			app.log.Printf("error storing block# %d into %s : %v\n", blockDump.BlockNumber, rep.Name(), err.Error())
			return err
		}
	}
	return nil
}
