package decode

import "fmt"

type Decoder interface {
	Code() string
	FilePath() string
}

type AnchorAccountDecoder struct {
	jsCode   string
	idlJson  string
	filePath string
}

func NewAnchorAccountDecoder(jsCode string, idlJson string, filePath string) *AnchorAccountDecoder {
	return &AnchorAccountDecoder{
		jsCode:   jsCode,
		idlJson:  idlJson,
		filePath: filePath,
	}
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
