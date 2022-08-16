package node

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// Solana RPC node client API.
type SolanaNode interface {
	// Return all accounts owned by the provided program public key.
	GetProgramAccounts(programId string) (rpc.GetProgramAccountsResult, error)
}

// Solana-go implementation.
type SolanaGo struct {
	client *rpc.Client
}

// Create a new solana-go RPC client.
func NewSolanaGo(endpoint string) *SolanaGo {
	return &SolanaGo{
		client: rpc.New(endpoint),
	}
}

// Return all accounts owned by the provided program public key.
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
