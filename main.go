package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dereference-xyz/trickle/config"
	"github.com/dereference-xyz/trickle/decode"
	"github.com/dereference-xyz/trickle/load"
	"github.com/dereference-xyz/trickle/model"
	"github.com/dereference-xyz/trickle/node"
	"github.com/dereference-xyz/trickle/service"
	"github.com/dereference-xyz/trickle/store"
	"github.com/dereference-xyz/trickle/store/sqlite"
)

func main() {
	configFile := flag.String("config", "", "path to config yaml file")
	flag.Parse()

	if *configFile == "" {
		panic("Please specify config file arg.")
	}

	cfg, err := config.Parse(*configFile)
	if err != nil {
		panic(err)
	}

	idlJson, err := os.ReadFile(cfg.Chains[0].Solana.Programs[0].IDL)
	if err != nil {
		panic(err)
	}

	programType, err := model.FromIDL(idlJson)
	if err != nil {
		panic(err)
	}

	accountStore, err := store.NewAccountStore(sqlite.NewDriver(cfg.Database.SQLite.File))
	err = accountStore.AutoMigrate(programType)
	if err != nil {
		panic(err)
	}

	solanaNode := node.NewSolanaGo(cfg.Chains[0].Solana.Node)
	decodeEngine := decode.NewV8Engine()
	loader := load.NewLoader(solanaNode, decodeEngine, accountStore)

	decoder, err := decode.NewAnchorAccountDecoder(config.DecoderFilePath, string(idlJson))
	if err != nil {
		panic(err)
	}

	err = loader.Load(decoder, cfg.Chains[0].Solana.Programs[0].ProgramId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data loaded successfully.")
	fmt.Println("Running service...")

	srv := service.NewService(accountStore, programType)
	err = srv.Router().Run()
	if err != nil {
		panic(err)
	}
}
