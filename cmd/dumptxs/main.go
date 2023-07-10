package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/c-kuroki/dumptxs/pkg/adapters"
	"github.com/c-kuroki/dumptxs/pkg/adapters/cosmos"
	"github.com/c-kuroki/dumptxs/pkg/adapters/file"
	pgInfra "github.com/c-kuroki/dumptxs/pkg/adapters/pg"
	"github.com/c-kuroki/dumptxs/pkg/app"
	"github.com/c-kuroki/dumptxs/pkg/config"
	"github.com/go-pg/pg/v10"
)

const helpMsg = `dumptxs -h                       # prints this help
dumptxs <block from> <block to>  # Dump transactions in a range of blocks
dumptxs <block>                  # Dump transactions of an specific blocks
dumptxs                          # Dump transactions in current block`

func main() {
	displayHelp := flag.Bool("h", false, "help")
	flag.Parse()
	if displayHelp != nil && *displayHelp {
		fmt.Println(helpMsg)
		return
	}
	parms, err := checkArgs(os.Args)
	if err != nil {
		fmt.Println("invalid parameters")
		fmt.Println(helpMsg)
		return
	}

	cfg := config.MustGetConfig()
	ctx := context.Background()

	logger := log.Default()
	cli, err := cosmos.NewCosmosClient(cfg.RPCURL, logger)
	if err != nil {
		log.Fatalf("error connecting to cosmos : %v\n", err)
	}
	fileRepository := file.NewFileRepository(cfg.FILEPATH)
	repos := []adapters.Repository{fileRepository}
	if len(cfg.DBURL) > 0 {
		opt, err := pg.ParseURL(cfg.DBURL)
		if err != nil {
			log.Fatalf("error db: %v\n", err)
		}
		db := pg.Connect(opt)
		pgRepository := pgInfra.NewPgRepository(db)
		repos = append(repos, pgRepository)
	}
	app := app.NewDumpTxsApp(cli, repos, logger)

	if len(parms) == 2 {
		_ = app.DumpRange(ctx, parms[0], parms[1])
	} else {
		// dump current block
		_ = app.DumpBlock(ctx, nil)
	}
}

func checkArgs(args []string) ([]int64, error) {
	var parms []int64
	switch len(args) {
	case 1:
		return parms, nil
	case 2:
		from, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			return nil, err
		}
		parms = append(parms, from)
		parms = append(parms, from)
	case 3:
		from, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			return nil, err
		}
		to, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			return nil, err
		}
		parms = append(parms, from)
		parms = append(parms, to)
	default:
		return nil, fmt.Errorf("invalid number of parameters")
	}
	return parms, nil
}
