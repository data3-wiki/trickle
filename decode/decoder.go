package decode

import (
	"fmt"
	"os"
)

type Decoder interface {
	Code() string
	FilePath() string
}

type AnchorAccountDecoder struct {
	jsCode   string
	idlJson  string
	filePath string
}

func NewAnchorAccountDecoder(decoderFilePath, idlJson string) (*AnchorAccountDecoder, error) {
	jsCode, err := os.ReadFile(decoderFilePath)
	if err != nil {
		return nil, err
	}

	return &AnchorAccountDecoder{
		jsCode:   string(jsCode),
		idlJson:  idlJson,
		filePath: decoderFilePath,
	}, nil
}

func (dec *AnchorAccountDecoder) Code() string {
	return fmt.Sprintf(`
		%s
		decoder.decode = value => decoder.decodeAccount(%s, value)
	`, dec.jsCode, dec.idlJson)
}

func (dec *AnchorAccountDecoder) FilePath() string {
	return dec.filePath
}
