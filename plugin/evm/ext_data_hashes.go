package evm

import (
	_ "embed"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

var (
	//go:embed songbird_ext_data_hashes.json
	rawSongbirdExtDataHashes []byte
	songbirdExtDataHashes    map[common.Hash]common.Hash

	//go:embed flare_ext_data_hashes.json
	rawFlareExtDataHashes []byte
	flareExtDataHashes    map[common.Hash]common.Hash
)

func init() {
	if err := json.Unmarshal(rawSongbirdExtDataHashes, &songbirdExtDataHashes); err != nil {
		panic(err)
	}
	rawSongbirdExtDataHashes = nil
	if err := json.Unmarshal(rawFlareExtDataHashes, &flareExtDataHashes); err != nil {
		panic(err)
	}
	rawFlareExtDataHashes = nil
}
