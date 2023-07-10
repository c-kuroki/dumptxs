## dumptxs

Cli tool to pull blocks/txs from a Cosmos Hub node

## Requirements

- make

- golang v1.19 or higher

- [golangci-lint](https://golangci-lint.run/)

## Service configuration

- Configuration is obtained from environment vars.

| Var                 | Default                                                           | Description                             |
| ------------------- | ----------------------------------------------------------------- | --------------------------------------- |
| RPC_URL              | https://cosmos-rpc.quickapi.com:443 | Cosmos Hub RPC URL |
| DB_URL  |  postgres://postgres@localhost:5432/dumptxs?sslmode=disable | PostgreSQL database connection URL, if empty, will not dump into the DB |
| DUMP_PATH |  ./ | Folder to dump json files |

## Compile and test

- `make` will build the cli

- `make test` run code tests

- `make lint` execute code linter

## Usage 

```
dumptxs -h                       # prints help
dumptxs <block from> <block to>  # Dump transactions in a range of blocks
dumptxs <block>                  # Dump transactions of an specific block
dumptxs                          # Dump transactions in current block
```

