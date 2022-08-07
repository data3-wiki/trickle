package node

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type SolanaNode struct {
	client *rpc.Client
}

func NewSolanaNode(endpoint string) *SolanaNode {
	return &SolanaNode{
		client: rpc.New(endpoint),
	}
}

func (node *SolanaNode) GetProgramAccounts(programId string) (rpc.GetProgramAccountsResult, error) {
	accounts, err := node.client.GetProgramAccounts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(programId),
	)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
