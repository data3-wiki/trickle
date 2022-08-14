package service

import (
	"os"
	"testing"

	"github.com/dereference-xyz/trickle/config"
	"github.com/dereference-xyz/trickle/decode"
	"github.com/dereference-xyz/trickle/load"
	"github.com/dereference-xyz/trickle/model"
	"github.com/dereference-xyz/trickle/node"
	"github.com/dereference-xyz/trickle/store"
	"github.com/dereference-xyz/trickle/store/sqlite"
	"github.com/stretchr/testify/require"
)

type Deps struct {
	programType  *model.ProgramType
	solanaNode   node.SolanaNode
	accountStore *store.AccountStore
	service      *Service
	loader       *load.Loader
}

func loadTestIDL(t *testing.T) ([]byte, *model.ProgramType) {
	idlJson, err := os.ReadFile("../test/squads_mpl.json")
	require.NoError(t, err)
	programType, err := model.FromIDL(idlJson)
	require.NoError(t, err)
	return idlJson, programType
}

func initDeps(t *testing.T, deps *Deps) {
	idlJson, programType := loadTestIDL(t)
	if deps.programType == nil {
		deps.programType = programType
	}

	if deps.solanaNode == nil {
		deps.solanaNode = node.NewSolanaNode("")
	}

	if deps.accountStore == nil {
		accountStore, err := store.NewAccountStore(sqlite.NewDriver(":memory:"))
		require.NoError(t, err)
		require.NoError(t, accountStore.AutoMigrate(deps.programType))
		deps.accountStore = accountStore
	}

	if deps.service == nil {
		NewService(deps.accountStore, deps.programType)
	}

	if deps.loader == nil {
		decodeEngine := decode.NewV8Engine()
		loader := load.NewLoader(deps.solanaNode, decodeEngine, deps.accountStore)
		decoder, err := decode.NewAnchorAccountDecoder("../"+config.DecoderFilePath, string(idlJson))
		require.NoError(t, err)
		require.NoError(t, loader.Load(decoder, "0xDEADBEEF"))
	}
}

func TestDataTypes(t *testing.T) {
	var deps Deps
	initDeps(t, &deps)
	// TODO
}
