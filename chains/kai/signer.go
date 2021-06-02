package kai

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	ethComm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/kardiachain/go-kardia/types"
	"github.com/ontio/ontology/common"
)

type Signer struct {
	PrivateKey *ecdsa.PrivateKey
	Address    ethComm.Address
}

func NewSigner(privateKey string) (*Signer, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("ApproveERC20, cannot decode private key")
	}
	address := crypto.PubkeyToAddress(priKey.PublicKey)
	return &Signer{
		PrivateKey: priKey,
		Address:    address,
	}, nil
}

func DeserializeTx(rawTx string) (*types.Transaction, error) {
	txData := ethComm.FromHex(rawTx)
	tx := &types.Transaction{}
	err := rlp.DecodeBytes(txData, tx)
	if err != nil {
		return nil, fmt.Errorf("deserializeTx: err: %s", err)
	}
	return tx, nil
}

func SerializeTx(tx *types.Transaction) (string, error) {
	bf := new(bytes.Buffer)
	err := rlp.Encode(bf, tx)
	if err != nil {
		return "", fmt.Errorf("signTx: encode signed tx err: %s", err)
	}
	signedRawTx := hexutil.Encode(bf.Bytes())
	return signedRawTx, nil
}

func ReverseOntAddress(ontAddress string) string {
	addr, _ := common.AddressFromBase58(ontAddress)
	//fmt.Printf("Address Form Base58 is %s\n", "0x"+hex.EncodeToString(addr[:]))
	return hex.EncodeToString(addr[:])
}
