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
package ont

import (
	"bytes"
	"crypto/elliptic"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/sm2"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"strings"
	"time"
)

func GetAccountByPassword(sdk *ontology_go_sdk.OntologySdk, path string, pwd []byte) (*ontology_go_sdk.Account, error) {
	wallet, err := sdk.OpenWallet(path)
	if err != nil {
		return nil, fmt.Errorf("open wallet error: %v", err)
	}
	user, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		return nil, fmt.Errorf("getDefaultAccount error: %v", err)
	}
	return user, nil
}

func GetCurveLabel(name string) (byte, error) {
	switch strings.ToUpper(name) {
	case strings.ToUpper(elliptic.P224().Params().Name):
		return 1, nil
	case strings.ToUpper(elliptic.P256().Params().Name):
		return 2, nil
	case strings.ToUpper(elliptic.P384().Params().Name):
		return 3, nil
	case strings.ToUpper(elliptic.P521().Params().Name):
		return 4, nil
	case strings.ToUpper(sm2.SM2P256V1().Params().Name):
		return 20, nil
	case strings.ToUpper(btcec.S256().Name):
		return 5, nil
	default:
		panic("err")
	}
}

func GetOntNoCompressKey(key keypair.PublicKey) []byte {
	var buf bytes.Buffer
	switch t := key.(type) {
	case *ec.PublicKey:
		switch t.Algorithm {
		case ec.ECDSA:
			// Take P-256 as a special case
			if t.Params().Name == elliptic.P256().Params().Name {
				return ec.EncodePublicKey(t.PublicKey, false)
			}
			buf.WriteByte(byte(0x12))
		case ec.SM2:
			buf.WriteByte(byte(0x13))
		}
		label, err := GetCurveLabel(t.Curve.Params().Name)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(label)
		buf.Write(ec.EncodePublicKey(t.PublicKey, false))
	default:
		panic("err")
	}
	return buf.Bytes()
}

func GetOntAccByPwd(acc, pwd string) (*ontology_go_sdk.Account, error) {
	wallet, err := ontology_go_sdk.OpenWallet(acc)
	if err != nil {
		return nil, fmt.Errorf("open wallet error: %v", err)
	}
	user, err := wallet.GetDefaultAccount([]byte(pwd))
	if err != nil {
		return nil, fmt.Errorf("getDefaultAccount error: %v", err)
	}

	return user, nil
}

func WaitOntTx(txhash common.Uint256, ont *ontology_go_sdk.OntologySdk) {
	tick := time.NewTicker(100 * time.Millisecond)
	var h uint32
	startTime := time.Now()
	for range tick.C {
		h, _ = ont.GetBlockHeightByTxHash(txhash.ToHexString())
		curr, _ := ont.GetCurrentBlockHeight()
		if h > 0 && curr > h {
			break
		}

		if startTime.Add(100 * time.Millisecond); startTime.Second() > 300 {
			panic(fmt.Errorf("tx( %s ) is not confirm for a long time ( over %d sec )",
				txhash.ToHexString(), 300))
		}
	}
}
