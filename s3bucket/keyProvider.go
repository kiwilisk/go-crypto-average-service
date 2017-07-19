package s3bucket

import "encoding/hex"

type S3KeyProvider interface {
	CreateKeyFor(symbol string) *string
}

type HexKeyProvider struct{}

func NewHexKeyProvider() *HexKeyProvider {
	return &HexKeyProvider{}
}

func (HexKeyProvider) CreateKeyFor(symbol string) *string {
	bytes := []byte(symbol)
	hexEncodedSymbol := hex.EncodeToString(bytes)
	key := hexEncodedSymbol + "_" + symbol
	return &key
}
