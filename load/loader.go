package load

import (
	"fmt"
	"os"
	"strings"

	"github.com/dereference-xyz/trickle/decode"
	"github.com/dereference-xyz/trickle/model"
	"github.com/dereference-xyz/trickle/node"
	"github.com/dereference-xyz/trickle/store"
)

type Loader struct {
	solanaNode   *node.SolanaNode
	decodeEngine *decode.V8Engine
	accountStore *store.AccountStore
}

func NewLoader(
	solanaNode *node.SolanaNode,
	decodeEngine *decode.V8Engine,
	accountStore *store.AccountStore) *Loader {
	return &Loader{
		solanaNode:   solanaNode,
		decodeEngine: decodeEngine,
		accountStore: accountStore,
	}
}

func (ld *Loader) Load(decoder decode.Decoder, programId string) error {
	accounts, err := ld.solanaNode.GetProgramAccounts(programId)
	if err != nil {
		return err
	}

	decodedAccounts := []*model.Account{}
	decodingErrors := []string{}
	for _, acc := range accounts {
		da, err := ld.decodeEngine.DecodeAccount(decoder, acc)
		if err != nil {
			decodingErrors = append(decodingErrors, err.Error())
		} else {
			decodedAccounts = append(decodedAccounts, da)
		}
	}

	if len(decodingErrors) > 0 {
		// TODO: Change to logger or return stats to caller.
		fmt.Fprintf(
			os.Stderr,
			"Errors:\n%s\n%d out of %d succeeded.\n",
			strings.Join(decodingErrors, "\n"),
			len(decodedAccounts),
			len(accounts))
	}

	if len(decodedAccounts) == 0 {
		return fmt.Errorf("No accounts were decoded.")
	}

	err = ld.accountStore.Create(decodedAccounts)
	if err != nil {
		return err
	}

	return nil
}
