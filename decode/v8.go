package decode

import (
	"encoding/base64"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	v8 "rogchap.com/v8go"
)

type V8Engine struct {
	isolate *v8.Isolate
}

func NewV8Engine() *V8Engine {
	return &V8Engine{
		isolate: v8.NewIsolate(),
	}
}

func (eng *V8Engine) DecodeAccount(decoder Decoder, account *rpc.KeyedAccount) (string, error) {
	ctx := v8.NewContext(eng.isolate)
	defer ctx.Close()
	ctx.RunScript(decoder.Code(), decoder.FilePath())

	data := base64.StdEncoding.EncodeToString(account.Account.Data.GetBinary())
	val, err := ctx.RunScript(fmt.Sprintf("decoder.decode('%s')", data), "main.js")
	if err != nil {
		return "", err
	}

	json, err := val.MarshalJSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}
