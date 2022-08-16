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

// Expected json object returned by the decoder.
type decodedAccount struct {
	// Type of the account.
	AccountType string `json:"accountType"`
	// Decoded account data.
	Decoded map[string]interface{} `json:"decoded"`
}

// V8 javascript engine used to execute the decoders.
type V8Engine struct {
	// V8 isolate that can only be accessed by one thread at a time.
	isolate *v8.Isolate
}

// Create a new V8 javascript engine for decoding.
func NewV8Engine() *V8Engine {
	return &V8Engine{
		isolate: v8.NewIsolate(),
	}
}

// Decode account data using the specified decoder for the given program.
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

// Parse and convert the value returned by the V8 engine into Golang types.
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
