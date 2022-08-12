package decode

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/dereference-xyz/trickle/model"
	"github.com/gagliardetto/solana-go/rpc"
	"rogchap.com/v8go"
	v8 "rogchap.com/v8go"
)

type decodedAccount struct {
	AccountType string                 `json:"accountType"`
	Decoded     map[string]interface{} `json:"decoded"`
}

type V8Engine struct {
	isolate *v8.Isolate
}

func NewV8Engine() *V8Engine {
	return &V8Engine{
		isolate: v8.NewIsolate(),
	}
}

func (eng *V8Engine) DecodeAccount(decoder Decoder, account *rpc.KeyedAccount) (*model.Account, error) {
	ctx := v8.NewContext(eng.isolate)
	defer ctx.Close()
	ctx.RunScript(decoder.Code(), decoder.FilePath())

	data := base64.StdEncoding.EncodeToString(account.Account.Data.GetBinary())
	val, err := ctx.RunScript(fmt.Sprintf("decoder.decode('%s')", data), "main.js")
	if err != nil {
		return nil, err
	}
	return convertValue(val)
}

func convertValue(val *v8go.Value) (*model.Account, error) {
	valJson, err := val.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var decoded decodedAccount
	err = json.Unmarshal(valJson, &decoded)
	if err != nil {
		return nil, err
	}

	// TODO: Case on PropertyType.
	for field, value := range decoded.Decoded {
		switch value.(type) {
		case string:
		default:
			serialized, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			decoded.Decoded[field] = string(serialized)
		}
	}

	return &model.Account{
		Type: decoded.AccountType,
		Data: decoded.Decoded,
	}, nil
}
