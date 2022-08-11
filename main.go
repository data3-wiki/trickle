package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dereference-xyz/trickle/decode"
	"github.com/dereference-xyz/trickle/model"
	"github.com/dereference-xyz/trickle/node"
	"github.com/gagliardetto/solana-go/rpc"
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

	_, err = model.FromIDL(idlJson)
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
	for _, acc := range accounts {
		fmt.Println(dec.DecodeAccount(decoder, acc))
	}
}
