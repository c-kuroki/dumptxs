package cosmos

import (
	"context"
	"log"

	"github.com/c-kuroki/dumptxs/pkg/model"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
)

type CosmosClient struct {
	client *rpchttp.HTTP
	log    *log.Logger
}

func NewCosmosClient(rpcURL string, log *log.Logger) (*CosmosClient, error) {
	client, err := client.NewClientFromNode(rpcURL)
	if err != nil {
		return nil, err
	}
	return &CosmosClient{
		client: client,
		log:    log,
	}, nil
}

func (cli *CosmosClient) Block(ctx context.Context, blocknum *int64) (*model.BlockDump, error) {
	res, err := cli.client.Block(ctx, blocknum)
	if err != nil {
		return nil, err
	}
	decoder := DefaultDecoder
	var jsonstr []string
	for _, tx := range res.Block.Data.Txs {
		txd, err := decoder.Decode(tx)
		if err != nil {
			cli.log.Printf("tx decoding error : %v\n", err.Error())
			continue
		}

		jsonb, err := txd.MarshalToJSON()
		if err != nil {
			cli.log.Printf("json marshal error :>%v\n", err.Error())
			continue
		}
		jsonstr = append(jsonstr, string(jsonb))
	}
	return &model.BlockDump{
		BlockNumber: res.Block.Header.Height,
		Dump:        jsonstr,
	}, nil
}
