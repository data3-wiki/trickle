package decode

import (
	"fmt"
	"os"
)

// Interface for a javascript decoder.
type Decoder interface {
	// Javascript code implementing the decoder.
	// The code should be completely self-contained and is expected to be packaged as a
	// var-type library.
	// See the V8 engine implementation for the expected function interface of the decoder.
	Code() string
	// File path of the decoder for error reporting / logging.
	FilePath() string
}

// Decoder for Solana Anchor framework account data.
type AnchorAccountDecoder struct {
	// Javascript code of the decoder.
	jsCode string
	// Anchor framework IDL json spec for the program being decoded.
	idlJson string
	// File path of the decoder.
	filePath string
}

// Creat new anchor account decoder with the given decoder and IDL json spec.
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

// Return javascript code implementing the decoder.
func (dec *AnchorAccountDecoder) Code() string {
	return fmt.Sprintf(`
		%s
		decoder.decode = value => decoder.decodeAccount(%s, value)
	`, dec.jsCode, dec.idlJson)
}

// Return file path of the decoder.
func (dec *AnchorAccountDecoder) FilePath() string {
	return dec.filePath
}
