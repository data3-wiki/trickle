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

func (eng *V8Engine) DecodeAccount(
	programType *model.ProgramType,
	decoder Decoder,
	account *rpc.KeyedAccount) (*model.Account, error) {
	ctx := v8.NewContext(eng.isolate)
	defer ctx.Close()
	ctx.RunScript(decoder.Code(), decoder.FilePath())

	data := base64.StdEncoding.EncodeToString(account.Account.Data.GetBinary())
	val, err := ctx.RunScript(fmt.Sprintf("decoder.decode('%s')", data), "main.js")
	if err != nil {
		return nil, err
	}
	return convertValue(programType, val)
}

func convertValue(programType *model.ProgramType, val *v8go.Value) (*model.Account, error) {
	valJson, err := val.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var decoded decodedAccount
	err = json.Unmarshal(valJson, &decoded)
	if err != nil {
		return nil, err
	}

	accountType, exists := programType.AccountType(decoded.AccountType)
	if !exists {
		return nil, fmt.Errorf(
			"Unsupported account type '%s'. Decoder and schema are out of sync. Probably a bug.",
			decoded.AccountType)
	}

	for field, value := range decoded.Decoded {
		propertyType, exists := accountType.PropertyType(field)
		if !exists {
			return nil, fmt.Errorf(
				"Unsupported property name '%s'. Decoder and schema are out of sync. Probably a bug.",
				field)
		}
		converted, err := convertDecodedValue(propertyType.DataType, value)
		if err != nil {
			return nil, err
		}
		decoded.Decoded[field] = converted
	}

	return &model.Account{
		AccountType: accountType,
		Type:        decoded.AccountType,
		Data:        decoded.Decoded,
	}, nil
}
