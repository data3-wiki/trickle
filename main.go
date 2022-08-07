package main

import (
	"fmt"

	"github.com/dereference-xyz/trickle/decoder"
	"github.com/dereference-xyz/trickle/node"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	solana := node.NewSolanaNode(rpc.MainNetBeta_RPC)
	accounts, err := solana.GetProgramAccounts("strmRqUCoQUgGUan5YhzUZa6KqdzwX5L6FpUxfmKg5m")
	if err != nil {
		panic(err)
	}

	dec := decoder.NewDecoder()
	for _, acc := range accounts {
		fmt.Println(dec.Decode(acc))
	}
}
