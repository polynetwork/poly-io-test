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
package btc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	common2 "github.com/ontio/ontology/common"
	config2 "github.com/polynetwork/btc-vendor-tools/config"
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly/common"
	"github.com/polynetwork/poly/native/service/governance/side_chain_manager"
	"github.com/polynetwork/poly/native/service/utils"
	"path"
	"strings"
)

type BtcInvoker struct {
	RChain      *poly_go_sdk.PolySdk
	RChainAcc   *poly_go_sdk.Account
	FromChainId int32
	BtcCli      *RestCli
	Signer      *BtcSigner
}

func NewBtcInvoker(rpc, wallet, pwd, btcRpc, btcUser, btcPwd, privk string) (*BtcInvoker, error) {
	invoker := &BtcInvoker{}
	invoker.RChain = poly_go_sdk.NewPolySdk()
	if err := SetUpPoly(invoker.RChain, rpc); err != nil {
		return nil, err
	}

	acc, err := GetAccountByPassword(invoker.RChain, wallet, []byte(pwd))
	if err != nil {
		return nil, err
	}
	invoker.RChainAcc = acc
	invoker.BtcCli = NewRestCli(btcRpc, btcUser, btcPwd)
	invoker.Signer, err = NewBtcSigner(privk)
	if err != nil {
		return nil, err
	}
	return invoker, nil
}

/**
Generate 3/5， 4/7 Vendors
*/
func (invoker *BtcInvoker) GenerateVendor(n, require int) (*Vendor, error) {
	privks := make([]*btcutil.WIF, n)
	addrs := make([]*btcutil.AddressPubKey, n)
	for i := 0; i < n; i++ {
		privk, err := btcec.NewPrivateKey(btcec.S256())
		if err != nil {
			return nil, err
		}
		privks[i], _ = btcutil.NewWIF(privk, config.BtcNet, true)
		addrs[i], _ = btcutil.NewAddressPubKey(privk.PubKey().SerializeCompressed(), config.BtcNet)
	}

	redeem, _ := txscript.MultiSigScript(addrs, require)
	p2sh, _ := btcutil.NewAddressScriptHash(redeem, config.BtcNet)
	hasher := sha256.New()
	hasher.Write(redeem)
	p2wsh, _ := btcutil.NewAddressWitnessScriptHash(hasher.Sum(nil), config.BtcNet)

	return &Vendor{
		Redeem:      redeem,
		HashKey:     btcutil.Hash160(redeem),
		AddressSet:  addrs,
		P2shAddr:    p2sh,
		P2wshAddr:   p2wsh,
		PrivateKeys: privks,
	}, nil
}

func (invoker *BtcInvoker) BindBtcxWithVendor(btcx string, btcxChainId uint64, vendor *Vendor) (common.Uint256, error) {
	ver := uint64(0)
	val, _ := invoker.RChain.GetStorage(utils.SideChainManagerContractAddress.ToHexString(),
		append(append(append([]byte(side_chain_manager.REDEEM_BIND),
			utils.GetUint64Bytes(1)...),
			utils.GetUint64Bytes(btcxChainId)...), vendor.HashKey...))
	if val != nil && len(val) != 0 {
		c := &side_chain_manager.ContractBinded{}
		err := c.Deserialization(common.NewZeroCopySource(val))
		if err != nil {
			return common.UINT256_EMPTY, err
		}
		ver = c.Ver + 1
	}

	btcx = strings.Replace(btcx, "0x", "", 1)
	var btcxBytes []byte
	switch btcxChainId {
	case config.DefConfig.OntChainID: //ONT ChainId: 3
		addr, err := common2.AddressFromHexString(btcx)
		if err != nil {
			return common.UINT256_EMPTY, err
		}
		btcxBytes = addr[:]
	case config.DefConfig.EthChainID: //ETH ChainID: 2
		addr, err := hex.DecodeString(btcx)
		if err != nil {
			return common.UINT256_EMPTY, err
		}
		btcxBytes = addr[:]
	case config.DefConfig.CMCrossChainId:
		btcxBytes = []byte(btcx)
	//todo：Add Cosmos + Others, need to refactor to registration mode
	default:
		return common.UINT256_EMPTY, fmt.Errorf("chain-id %d not supported", btcxChainId)
	}

	hash := btcutil.Hash160(
		append(append(append(append(vendor.Redeem, utils.GetUint64Bytes(1)...),
			btcxBytes...), utils.GetUint64Bytes(btcxChainId)...), utils.GetUint64Bytes(ver)...))
	sigs := make([][]byte, len(vendor.PrivateKeys))
	for i := 0; i < len(vendor.PrivateKeys); i++ {
		sig, err := vendor.PrivateKeys[i].PrivKey.Sign(hash)
		if err != nil {
			return common.UINT256_EMPTY, fmt.Errorf("failed to sign contract-binding hash: %v", err)
		}
		sigs[i] = sig.Serialize()
	}
	txhash, err := invoker.RChain.Native.Scm.RegisterRedeem(1, btcxChainId, vendor.Redeem, btcxBytes,
		ver, sigs, invoker.RChainAcc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return txhash, nil
}

func (invoker *BtcInvoker) BindBtcTxParam(feeRate, minChange uint64, vendor *Vendor) (common.Uint256, error) {
	ver := uint64(0)
	val, _ := invoker.RChain.GetStorage(utils.SideChainManagerContractAddress.ToHexString(),
		append(append([]byte(side_chain_manager.BTC_TX_PARAM), vendor.HashKey...),
			utils.GetUint64Bytes(1)...))
	if val != nil && len(val) > 0 {
		detail := &side_chain_manager.BtcTxParamDetial{}
		if err := detail.Deserialization(common.NewZeroCopySource(val)); err != nil {
			return common.UINT256_EMPTY, err
		}
		ver = detail.PVersion + 1
	}

	hash := btcutil.Hash160(
		append(append(append(append(vendor.Redeem, utils.GetUint64Bytes(1)...),
			utils.GetUint64Bytes(feeRate)...), utils.GetUint64Bytes(minChange)...), utils.GetUint64Bytes(ver)...))
	sigs := make([][]byte, len(vendor.PrivateKeys))
	for i := 0; i < len(vendor.PrivateKeys); i++ {
		sig, err := vendor.PrivateKeys[i].PrivKey.Sign(hash)
		if err != nil {
			return common.UINT256_EMPTY, fmt.Errorf("failed to sign contract-binding hash: %v", err)
		}
		sigs[i] = sig.Serialize()
	}

	txhash, err := invoker.RChain.Native.Scm.SetBtcTxParam(vendor.Redeem, 1, feeRate, minChange, ver,
		sigs, invoker.RChainAcc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return txhash, err
}

func (invoker *BtcInvoker) GetAccInfo() (string, error) {
	n, err := invoker.BtcCli.GetBlockCount()
	if err != nil {
		return "", err
	}
	utxos, err := invoker.BtcCli.ListUnspent(1, n, invoker.Signer.Address)
	if err != nil {
		return "", err
	}
	sum := int64(0)
	for _, u := range utxos {
		sum += u.Amount
	}
	return fmt.Sprintf("BITCOIN: acc: %s, asset: [ btc: %d ]", invoker.Signer.Address, sum), nil
}

type Vendor struct {
	PrivateKeys []*btcutil.WIF
	AddressSet  []*btcutil.AddressPubKey
	Redeem      []byte
	HashKey     []byte
	P2shAddr    *btcutil.AddressScriptHash
	P2wshAddr   *btcutil.AddressWitnessScriptHash
}

func NewVendorFromConfig() (*Vendor, error) {
	v := &Vendor{}
	strArr := strings.Split(config.DefConfig.BtcExistingVendorPrivks, ",")
	if len(strArr) == 0 {
		return nil, fmt.Errorf("no private keys")
	}
	privks := make([]*btcutil.WIF, len(strArr))
	addrs := make([]*btcutil.AddressPubKey, len(strArr))
	for i, str := range strArr {
		wif, err := btcutil.DecodeWIF(str)
		if err != nil {
			return nil, err
		}
		privks[i] = wif
		addr, _ := btcutil.NewAddressPubKey(wif.SerializePubKey(), config.BtcNet)
		addrs[i] = addr
	}
	v.PrivateKeys = privks
	v.AddressSet = addrs

	v.Redeem, _ = txscript.MultiSigScript(addrs, config.DefConfig.BtcMultiSigRequire)
	v.HashKey = btcutil.Hash160(v.Redeem)
	v.P2shAddr, _ = btcutil.NewAddressScriptHash(v.Redeem, config.BtcNet)
	hasher := sha256.New()
	hasher.Write(v.Redeem)
	v.P2wshAddr, _ = btcutil.NewAddressWitnessScriptHash(hasher.Sum(nil), config.BtcNet)

	return v, nil
}

func (v *Vendor) EncryptPrivateKeys(file, pwd string) ([]string, error) {
	pathSet := make([]string, len(v.PrivateKeys))
	for i, priv := range v.PrivateKeys {
		path := fmt.Sprintf("%s_%d", file, i+1)
		if err := encryptBtcPrivk(path, priv.String(), pwd); err != nil {
			return nil, fmt.Errorf("failed to encrypt private key %s: %v", priv.String(), err)
		}
		pathSet[i] = path
	}
	return pathSet, nil
}

func (v *Vendor) UpdateConfigFile(file, suffix, privkFile, pwd string) error {
	conf, err := config2.NewConfig(file)
	if err != nil {
		return err
	}
	conf.Redeem = hex.EncodeToString(v.Redeem)
	conf.BtcWalletPwd = pwd
	conf.BtcPrivkFile = privkFile
	conf.ConfigDBPath = path.Join(conf.ConfigDBPath, "vendor"+suffix)
	return conf.Save(file + suffix)
}
