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
package testcase

import (
	"github.com/polynetwork/poly-io-test/testframework"
)

//TestCase list
func init() {
	testframework.TFramework.RegTestCase("SendOntToEthChain", SendOntToEthChain)
	testframework.TFramework.RegTestCase("SendOnteToOntChain", SendEOntToOntChain) // Onte means ONT on ethereum.
	testframework.TFramework.RegTestCase("SendOngToEthChain", SendOngToEthChain)
	testframework.TFramework.RegTestCase("SendOngeToOntChain", SendOngeToOntChain)
	testframework.TFramework.RegTestCase("SendOEP4ToEthChain", SendOEP4ToEthChain)
	testframework.TFramework.RegTestCase("SendOEP4eToOntChain", SendOEP4eToOntChain)

	testframework.TFramework.RegTestCase("SendEthToOntChain", SendEthToOntChain)
	testframework.TFramework.RegTestCase("SendEthoToEthChain", SendEthoToEthChain) // Etho means ETH on ontology.
	testframework.TFramework.RegTestCase("SendERC20ToOntChain", SendERC20ToOntChain)
	testframework.TFramework.RegTestCase("SendERC20oToEthChain", SendOERC20ToEthChain)

	testframework.TFramework.RegTestCase("SendBtcoToBtcChain", SendBtcoToBtcChain)
	testframework.TFramework.RegTestCase("SendBtcToOntChain", SendBtcToOntChain)
	testframework.TFramework.RegTestCase("SendBtcToEthChain", SendBtcToEthChain)
	testframework.TFramework.RegTestCase("SendBtceToBtcChain", SendBtceToBtcChain)
	testframework.TFramework.RegTestCase("SendBtcoToEthChain", SendBtcoToEthChain)
	testframework.TFramework.RegTestCase("SendBtceToOntChain", SendBtceToOntChain)

	testframework.TFramework.RegTestCase("SendBtcToEthInBatch", SendBtcToEthInBatch)
	testframework.TFramework.RegTestCase("SendBtcToOntInBatch", SendBtcToOntInBatch)
	testframework.TFramework.RegTestCase("SendBtceToBtcInBatch", SendBtceToBtcInBatch)
	testframework.TFramework.RegTestCase("SendBtcoToBtcInBatch", SendBtcoToBtcInBatch)
	testframework.TFramework.RegTestCase("SendBtcoToBtceInBatch", SendBtcoToBtceInBatch)
	testframework.TFramework.RegTestCase("SendBtceToBtcoInBatch", SendBtceToBtcoInBatch)
	testframework.TFramework.RegTestCase("SendOntToEthInBatch", SendOntToEthInBatch)
	testframework.TFramework.RegTestCase("SendOnteToOntInBatch", SendOnteToOntInBatch)
	testframework.TFramework.RegTestCase("SendEthToOntInBatch", SendEthToOntInBatch)
	testframework.TFramework.RegTestCase("SendEthoToEthInBatch", SendEthoToEthInBatch)

	testframework.TFramework.RegTestCase("BtcCircle", BtcCircle)
	testframework.TFramework.RegTestCase("OntCircle", OntCircle)
	testframework.TFramework.RegTestCase("EthCircle", EthCircle)
	testframework.TFramework.RegTestCase("OngCircle", OngCircle)
	testframework.TFramework.RegTestCase("Erc20Circle", Erc20Circle)
	testframework.TFramework.RegTestCase("Oep4Circle", Oep4Circle)
	testframework.TFramework.RegTestCase("OntCircleWithoutCosmos", OntCircleWithoutCosmos)

	//no ether
	testframework.TFramework.RegTestCase("BtcOntCircle", BtcOntCircle)

	// cosmos
	testframework.TFramework.RegTestCase("SendBtcToCosmosAndBack", SendBtcToCosmosAndBack)
	testframework.TFramework.RegTestCase("SendEthToCosmosAndBack", SendEthToCosmosAndBack)
	testframework.TFramework.RegTestCase("SendErc20ToCosmosAndBack", SendErc20ToCosmosAndBack)
	testframework.TFramework.RegTestCase("SendOntToCosmosAndBack", SendOntToCosmosAndBack)
	testframework.TFramework.RegTestCase("SendOngToCosmosAndBack", SendOngToCosmosAndBack)
	testframework.TFramework.RegTestCase("SendOep4ToCosmosAndBack", SendOep4ToCosmosAndBack)
	testframework.TFramework.RegTestCase("SendZeroOntToEth", SendZeroOntToEth)
	testframework.TFramework.RegTestCase("USDTCircle", USDTCircle)
	testframework.TFramework.RegTestCase("WBTCCircle", WBTCCircle)
	testframework.TFramework.RegTestCase("DAICircle", DAICircle)
	testframework.TFramework.RegTestCase("NeoToEthCircle", NeoToEthCircle)
	testframework.TFramework.RegTestCase("NeoToOntCircle", NeoToOntCircle)
	testframework.TFramework.RegTestCase("OntToNeoAndBack", OntToNeoAndBack)
	testframework.TFramework.RegTestCase("EthToNeoAndBack", EthToNeoAndBack)
	testframework.TFramework.RegTestCase("OntdToNeoAndBack", OntdToNeoAndBack)
	testframework.TFramework.RegTestCase("OntdToEthAndBack", OntdToEthAndBack)
	testframework.TFramework.RegTestCase("USDCCircle", USDCCircle)
	testframework.TFramework.RegTestCase("BnbToBsc", BnbToBsc)
	testframework.TFramework.RegTestCase("BnbToEthAndBack", BnbToEthAndBack)
	testframework.TFramework.RegTestCase("EthBnbToBscOnce", EthBnbToBscOnce)
	testframework.TFramework.RegTestCase("BnbToEthOnce", BnbToEthOnce)
	testframework.TFramework.RegTestCase("RenBTCCircle", RenBTCCircle)
	testframework.TFramework.RegTestCase("EOntdToNeoAndBack", EOntdToNeoAndBack)

	// huobi eco chain
	testframework.TFramework.RegTestCase("SendHecoHtToEthAndBack", SendHecoHtToEthAndBack)
	testframework.TFramework.RegTestCase("SendEthEtherToHecoAndBack", SendEthEtherToHecoAndBack)
	testframework.TFramework.RegTestCase("SendErc20ToHecoAndBack", SendErc20ToHecoAndBack)
	testframework.TFramework.RegTestCase("SendHecoHtToNeoAndBack", SendHecoHtToNeoAndBack)
	testframework.TFramework.RegTestCase("SendHecoHrc20ToNeoAndBack", SendHecoHrc20ToNeoAndBack)
}
