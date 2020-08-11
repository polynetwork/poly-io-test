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
package main

//
//import (
//	"fmt"
//	"os"
//	"runtime/trace"
//)
//
//func main() {
//	//ont := ontology_go_sdk.NewOntologySdk()
//	//
//	//for _, s := range []string{"73", "74", "75", "76", "78", "79", "80"} {
//	//	ont.NewRpcClient().SetAddress("http://172.168.3." + s + ":20336")
//	//	curr, _ := ont.GetCurrentBlockHeight()
//	//	h := uint32(50996)
//	//	res, err := ont.GetCrossChainMsg(h)
//	//	if err != nil {
//	//		fmt.Println(err)
//	//		return
//	//	}
//	//
//	//	hashes, _ := ont.GetBlockTxHashesByHeight(h)
//	//	numTx := len(hashes.Transactions)
//	//
//	//	//raw, _ := hex.DecodeString(res)
//	//	//ccm := &types.CrossChainMsg{}
//	//	//err = ccm.Deserialization(common.NewZeroCopySource(raw))
//	//	//if err != nil {
//	//	//	fmt.Println(err)
//	//	//	return
//	//	//}
//	//	fmt.Println(s + ": [" + res + "]")
//	//	//fmt.Printf("ccm: [ %+v ]\n", ccm)
//	//	i := 0
//	//	events, _ := ont.GetSmartContractEventByBlock(h)
//	//	for _, e := range events {
//	//		for _, n := range e.Notify {
//	//			key, ok := n.States.([]interface{})[0].(string)
//	//			if ok && key == cross_chain_manager.MAKE_FROM_ONT_PROOF {
//	//				hash, _ := common.Uint256FromHexString(e.TxHash)
//	//				i++
//	//				fmt.Printf("%d txhash: %s [ contract: %s, states: %v ]\n", i, hex.EncodeToString(hash[:]), n.ContractAddress, n.States)
//	//			}
//	//		}
//	//	}
//	//	blk, _ := ont.GetBlockByHeight(curr)
//	//	fmt.Println(curr, blk.Header.Timestamp, numTx)
//	//	//events, err := ont.GetSmartContractEventByBlock(49903)
//	//	//if err != nil {
//	//	//	fmt.Println(err)
//	//	//	return
//	//	//}
//	//	//for _, e := range events {
//	//	//	for _, n := range e.Notify {
//	//	//		arr := n.States.([]interface{})
//	//	//		contractMethod, _ := arr[0].(string)
//	//	//		if contractMethod != "verifyToOntProof" {
//	//	//			continue
//	//	//		}
//	//	//
//	//	//		fmt.Printf("alliance: %s, btc: %s\n", arr[1].(string), arr[2].(string))
//	//	//	}
//	//	//}
//	//
//	//	//b, _ := ont.GetBlockByHeight(49903)
//	//	//for _, tx := range b.Transactions {
//	//	//	hash := tx.Hash()
//	//	//	fmt.Println(hash.ToHexString(), hex.EncodeToString([]byte{byte(tx.TxType)}))
//	//	//}
//	//
//	//	//_, err := ont.GetTransaction("08780d001425bdcaa8531f896d158a24cace26a9ed449d3d4c590cf58f06c36e")
//	//	//if err != nil {
//	//	//	fmt.Println(s, err)
//	//	//}
//	//
//	//	//state, err := ont.GetMemPoolTxState("08780d001425bdcaa8531f896d158a24cace26a9ed449d3d4c590cf58f06c36e")
//	//	//if err != nil {
//	//	//	fmt.Println(err)
//	//	//}
//	//	//if state != nil {
//	//	//	for _, s := range state.State {
//	//	//		fmt.Println(s.Type, s.Height, s.ErrCode)
//	//	//	}
//	//	//}
//	//	//hash := tx.Hash()
//	//	//fmt.Println(hash.ToHexString())
//	//}
//	//
//	////w, err := ont.CreateWallet("./ont")
//	////if err != nil {
//	////	fmt.Println(err)
//	////	return
//	////}
//	////raw, _ := hex.DecodeString("5f2fe68215476abb9852cfa7da31ef00aa1468782d5ca809da5c4e1390b8ee45")
//	////acc, err := ontology_go_sdk.NewAccountFromPrivateKey(raw, signature.SHA256withECDSA)
//	////if err != nil {
//	////	fmt.Println(err)
//	////	return
//	////}
//	////wif, _ := keypair.Key2WIF(acc.PrivateKey)
//	////_, err = w.NewAccountFromWIF(wif, []byte("123123"))
//	////if err != nil {
//	////	fmt.Println(err)
//	////	return
//	////}
//	////w.Save()
//	//
//	////raw, err := hex.DecodeString("")
//	////if err != nil {
//	////	panic(err)
//	////}
//	////
//	////fmt.Println(utils.OntContractAddress.ToHexString())
//	////
//	////aa, _ := common.Uint256ParseFromBytes(raw)
//	//fmt.Println(aa.ToHexString())
//	f := func(arr []int) {
//		arr = append(arr, 1)
//		return
//	}
//	arr := make([]int, 0)
//	f(arr)
//	fmt.Println(arr)
//}

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}
