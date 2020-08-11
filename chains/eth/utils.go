/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
 */
package eth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polynetwork/poly/common"

	ethComm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type EthSigner struct {
	PrivateKey *ecdsa.PrivateKey
	Address    ethComm.Address
}

func NewEthSigner(privateKey string) (*EthSigner, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("ApproveERC20, cannot decode private key")
	}
	address := crypto.PubkeyToAddress(priKey.PublicKey)
	return &EthSigner{
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
