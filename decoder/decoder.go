package decoder

import (
	"encoding/base64"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	v8 "rogchap.com/v8go"
)

type Decoder struct {
	isolate *v8.Isolate
}

func NewDecoder() *Decoder {
	return &Decoder{
		isolate: v8.NewIsolate(),
	}
}

func (dec *Decoder) Decode(account *rpc.KeyedAccount) (string, error) {
	ctx := v8.NewContext(dec.isolate)
	defer ctx.Close()
	ctx.RunScript(`
		function decode(value) {
			return "decoded_" + value
		}
	`, "decoder.js")

	data := base64.StdEncoding.EncodeToString(account.Account.Data.GetBinary())
	val, err := ctx.RunScript(fmt.Sprintf("decode('%s')", data), "main.js")
	if err != nil {
		return "", err
	}
	return val.String(), nil
}
