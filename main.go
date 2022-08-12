package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dereference-xyz/trickle/decode"
	"github.com/dereference-xyz/trickle/model"
	"github.com/dereference-xyz/trickle/node"
	"github.com/dereference-xyz/trickle/store"
	"github.com/gagliardetto/solana-go/rpc"
	"gorm.io/driver/sqlite"
)

func main() {
	idlJsonFile := flag.String("idl", "", "path to idl.json file")
	programId := flag.String("programId", "", "program id to pull data for")
	flag.Parse()

	decoderFilePath := "js/decoder/anchor/dist/decoder.js"
	decoderCode, err := os.ReadFile(decoderFilePath)
	if err != nil {
		panic(err)
	}

	idlJson, err := os.ReadFile(*idlJsonFile)
	if err != nil {
		panic(err)
	}

	accountTypes, err := model.FromIDL(idlJson)
	if err != nil {
		panic(err)
	}

	// TODO: Add CLI flag for db path.
	accountStore, err := store.NewAccountStore(sqlite.Open("./test.db"))
	err = accountStore.AutoMigrate(accountTypes)
	if err != nil {
		panic(err)
	}

	solana := node.NewSolanaNode(rpc.MainNetBeta_RPC)
	accounts, err := solana.GetProgramAccounts(*programId)
	if err != nil {
		panic(err)
	}

	dec := decode.NewV8Engine()
	decoder := decode.NewAnchorAccountDecoder(string(decoderCode), string(idlJson), decoderFilePath)

	decodedAccounts := []*model.Account{}
	decodingErrors := []string{}
	for _, acc := range accounts {
		da, err := dec.DecodeAccount(decoder, acc)
		if err != nil {
			decodingErrors = append(decodingErrors, err.Error())
		} else {
			decodedAccounts = append(decodedAccounts, da)
		}
	}

	if len(decodingErrors) > 0 {
		fmt.Fprintf(
			os.Stderr,
			"Errors:\n%s\n%d out of %d succeeded.\n",
			strings.Join(decodingErrors, "\n"),
			len(decodedAccounts),
			len(accounts))
	}

	if len(decodedAccounts) == 0 {
		panic("No accounts were decoded.")
	}

	err = accountStore.Create(decodedAccounts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data loaded successfully.")
}
