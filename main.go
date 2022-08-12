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
	"gorm.io/driver/sqlite"
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

	decoderCode, err := os.ReadFile(config.DecoderFilePath)
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

	accountStore, err := store.NewAccountStore(sqlite.Open(cfg.Database.SQLite.File))
	err = accountStore.AutoMigrate(programType)
	if err != nil {
		panic(err)
	}

	solanaNode := node.NewSolanaNode(cfg.Chains[0].Solana.Node)
	decodeEngine := decode.NewV8Engine()
	loader := load.NewLoader(solanaNode, decodeEngine, accountStore)
	decoder := decode.NewAnchorAccountDecoder(string(decoderCode), string(idlJson), config.DecoderFilePath)

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
