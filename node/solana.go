package node

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type SolanaNode interface {
	GetProgramAccounts(programId string) (rpc.GetProgramAccountsResult, error)
}

type SolanaGo struct {
	client *rpc.Client
}

func NewSolanaGo(endpoint string) *SolanaGo {
	return &SolanaGo{
		client: rpc.New(endpoint),
	}
}

func (node *SolanaGo) GetProgramAccounts(programId string) (rpc.GetProgramAccountsResult, error) {
	accounts, err := node.client.GetProgramAccounts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(programId),
	)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
