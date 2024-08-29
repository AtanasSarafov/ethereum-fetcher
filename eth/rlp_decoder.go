package eth

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
)

func DecodeRLPHex(rlpHex string) ([]string, error) {
	data, err := hex.DecodeString(rlpHex)
	if err != nil {
		return nil, err
	}

	var txHashes []string
	err = rlp.DecodeBytes(data, &txHashes)
	if err != nil {
		return nil, err
	}

	return txHashes, nil
}
