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
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly-io-test/testframework"
	"time"
)

func SendOntToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	err := SendOntCrossEth(ctx, status, GetRandAmount(config.DefConfig.OntValLimit, 1))
	if err != nil {
		log.Error("SendOntToEthChain, SendOntCrossEth error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendEOntToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendEOntCrossOnt(ctx, status, config.DefConfig.EthOntx, GetRandAmount(config.DefConfig.OntValLimit, 1)); err != nil {
		log.Errorf("SendEOntToOntChain, SendEOntCrossOnt error: %v", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendEthToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendEthCrossOnt(ctx, status, 1); err != nil {
		log.Errorf("SendEthToOntChain error: %v", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendEthoToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendEthoCrossEth(ctx, status, config.DefConfig.OntEth, GetRandAmount(config.DefConfig.EthValLimit, 1)); err != nil {
		log.Errorf("SendEthoToEthChain, SendEthoCrossEth error: %v", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendBtcoToBtcChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	err := SendBtcoCrossBtc(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract))
	if err != nil {
		log.Errorf("SendBtcoToBtcChain, SendBtcoCrossBtc error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendBtcToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	err := SendBtcCrossOnt(ctx, status, int64(GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)))
	if err != nil {
		log.Errorf("SendBtcToOntChain, SendBtcCrossOnt error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendBtcToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendBtcCrossEth(ctx, status, int64(GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract))); err != nil {
		log.Errorf("SendBtcToEthChain, SendBtcCrossEth error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendBtceToBtcChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendBtceCrossBtc(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit,
		config.DefConfig.BtcMinOutputValFromContract)); err != nil {
		log.Errorf("SendBtceToBtcChain, SendBtcCrossBtc error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendBtcoToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendBtcoCrossBtce(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit,
		config.DefConfig.BtcMinOutputValFromContract)); err != nil {
		log.Errorf("SendBtcoToEthChain, SendBtcoCrossBtce error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendBtceToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendBtceCrossBtco(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit,
		config.DefConfig.BtcMinOutputValFromContract)); err != nil {
		log.Errorf("SendBtceToOntChain, SendBtceCrossBtco error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

// send btc to travel btc->eth->ont->btc and btc->ont->eth->btc
func BtcCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendBtcCrossEth(ctx, status, int64(amt)); err != nil {
				log.Errorf("BtcCircle, SendBtcCrossEth error: %v", err)
				return false
			}
			if err := SendBtcCrossOnt(ctx, status, int64(amt)); err != nil {
				log.Errorf("BtcCircle, SendBtcCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("send %d btc to eth/ont, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("btc1 is on ethereum now; btc2 is on ontology now")

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendBtcFromEthereumToCosmos(ctx, status, amt); err != nil {
				log.Errorf("BtcCircle, SendBtcFromEthereumToCosmos error: %v", err)
				return false
			}
			if err := SendBtcFromOntologyToCosmos(ctx, status, amt); err != nil {
				log.Errorf("BtcCircle, SendBtcFromOntologyToCosmos error: %v", err)
				return false
			}
		}
		log.Infof("send %d btco to cosmos and btce to cosmos, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("btc1 is on cosmos now; btc2 is on cosmos now")

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendBtcFromCosmosToOntology(ctx, status, amt); err != nil {
				log.Errorf("BtcCircle, SendBtcFromCosmosToOntology error: %v", err)
				return false
			}
			if err := SendBtcFromCosmosToEthereum(ctx, status, amt); err != nil {
				log.Errorf("BtcCircle, SendBtcFromCosmosToEthereum error: %v", err)
				return false
			}
		}
		log.Infof("send %d btc from cosmos to Ethereum and to Ontology, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("btc1 is on ontology now; btc2 is on ethereum now")

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendBtcoCrossBtc(ctx, status, amt); err != nil {
				log.Errorf("BtcCircle, SendBtcoCrossBtc error: %v", err)
				return false
			}
			if err := SendBtceCrossBtc(ctx, status, amt); err != nil {
				log.Errorf("BtcCircle, SendBtceCrossBtc error: %v", err)
				return false
			}
		}
		log.Infof("send %d btc from Ethereum and Ontology to bitcoin, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("btc1&2 is back now")
	}

	status.SetItSuccess(1)
	return true
}

func OntCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// ont->eth
			if err := SendOntCrossEth(ctx, status, amt); err != nil {
				log.Errorf("OntCircle, SendOntCrossEth error: %v", err)
				return false
			}
			if err := SendOntCrossCosmos(ctx, status, amt); err != nil {
				log.Errorf("OntCircle, SendOntCrossCosmos error: %v", err)
				return false
			}
		}
		log.Infof("OntCircle, send %d ont to eth and to cosmos, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// ont->eth
			if err := SendOntFromEthereumToCosmos(ctx, status, amt); err != nil {
				log.Errorf("OntCircle, SendOntFromEthereumToCosmos error: %v", err)
				return false
			}
			if err := SendOntFromCosmosToEthereum(ctx, status, amt); err != nil {
				log.Errorf("OntCircle, SendOntFromCosmosToEthereum error: %v", err)
				return false
			}
		}
		log.Infof("OntCircle, send %d ont from ethereum to cosmos and from cosmos to ethereum, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOntFromCosmosToOntology(ctx, status, amt); err != nil {
				log.Errorf("OntCircle, SendOntFromCosmosToOntology error: %v", err)
				return false
			}
			if err := SendEOntCrossOnt(ctx, status, config.DefConfig.EthOntx, amt); err != nil {
				log.Errorf("OntCircle, SendEOntCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("OntCircle, send %d ont from cosmos to ontology, and from ethereum to ontology, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("OntCircle, ont all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func OngCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OngValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOngCrossEth(ctx, status, amt); err != nil {
				log.Errorf("OngCircle, SendOngCrossEth error: %v", err)
				return false
			}
			if err := SendOngCrossCosmos(ctx, status, amt); err != nil {
				log.Errorf("OngCircle, SendOngCrossCosmos error: %v", err)
				return false
			}
		}
		log.Infof("OngCircle, send %d ong to eth and to cosmos, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOngFromEthereumToCosmos(ctx, status, amt); err != nil {
				log.Errorf("OngCircle, SendOngFromEthereumToCosmos error: %v", err)
				return false
			}
			if err := SendOngFromCosmosToEthereum(ctx, status, amt); err != nil {
				log.Errorf("OngCircle, SendOngFromCosmosToEthereum error: %v", err)
				return false
			}
		}
		log.Infof("OngCircle, send %d ong from ethereum to cosmos and from cosmos to ethereum, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOngFromCosmosToOntology(ctx, status, amt); err != nil {
				log.Errorf("OngCircle, SendOngFromCosmosToOntology error: %v", err)
				return false
			}
			if err := SendOngeCrossOnt(ctx, status, config.DefConfig.EthOngx, amt); err != nil {
				log.Errorf("OngCircle, SendOngeCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("OngCircle, send %d ong from cosmos to ontology, and from ethereum to ontology, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("OngCircle, ong all received ( batch: %d )", i)
	}
	status.SetItSuccess(1)
	return true
}

func EthCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.EthValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// eth->ont
			if err := SendEthCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("EthCircle, SendEthCrossOnt error: %v", err)
				return false
			}
			if err := SendEthCrossCosmos(ctx, status, amt); err != nil {
				log.Errorf("EthCircle, SendEthCrossCosmos error: %v", err)
				return false
			}
		}
		log.Infof("EthCircle, send %d eth to ont and to cosmos, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// eth->ont
			if err := SendEthFromOntologyToCosmos(ctx, status, amt); err != nil {
				log.Errorf("EthCircle, SendEthFromOntologyToCosmos error: %v", err)
				return false
			}
			if err := SendEthFromCosmosToOntology(ctx, status, amt); err != nil {
				log.Errorf("EthCircle, SendEthFromCosmosToOntology error: %v", err)
				return false
			}
		}
		log.Infof("EthCircle, send %d eth from ontology to cosmos and from cosmos to ontology, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// etho->eth
			if err := SendEthFromCosmosToEthereum(ctx, status, amt); err != nil {
				log.Errorf("EthCircle, SendEthFromCosmosToEthereum error: %v", err)
				return false
			}
			if err := SendEthoCrossEth(ctx, status, config.DefConfig.OntEth, amt); err != nil {
				log.Errorf("EthCircle, SendEthoCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("EthCircle, send %d eth from cosmos and from ontology to ethereum, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("EthCircle, eth all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func Erc20Circle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.Erc20ValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendERC20CrossOnt(ctx, status, config.DefConfig.EthErc20, amt); err != nil {
				log.Errorf("Erc20Circle, SendERC20CrossOnt error: %v", err)
				return false
			}
			if err := SendErc20CrossCosmos(ctx, status, config.DefConfig.EthErc20, amt); err != nil {
				log.Errorf("Erc20Circle, SendErc20CrossCosmos error: %v", err)
				return false
			}
		}
		log.Infof("Erc20Circle, send %d Erc20 to ont and to cosmos, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendErc20FromOntologyToCosmos(ctx, status, amt); err != nil {
				log.Errorf("Erc20Circle, SendErc20FromOntologyToCosmos error: %v", err)
				return false
			}
			if err := SendErc20FromCosmosToOntology(ctx, status, amt); err != nil {
				log.Errorf("Erc20Circle, SendErc20FromCosmosToOntology error: %v", err)
				return false
			}
		}
		log.Infof("Erc20Circle, send %d erc20 from ontology to cosmos and from cosmos to ontology, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendErc20FromCosmosToEthereum(ctx, status, amt); err != nil {
				log.Errorf("Erc20Circle, SendErc20FromCosmosToEthereum error: %v", err)
				return false
			}
			if err := SendOERC20CrossEth(ctx, status, config.DefConfig.OntErc20, amt); err != nil {
				log.Errorf("Erc20Circle, SendOERC20CrossEth error: %v", err)
				return false
			}
		}
		log.Infof("Erc20Circle, send %d erc20 from cosmos and from ontology to ethereum, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
		log.Infof("Erc20Circle, erc20 all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func Oep4Circle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.Oep4ValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// oep4->eth
			if err := SendOEP4CrossEth(ctx, status, config.DefConfig.OntOep4, amt); err != nil {
				log.Errorf("Oep4Circle, SendOEP4CrossEth error: %v", err)
				return false
			}
		}
		log.Infof("send %d Oep4 to eth, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// oep4e->ont
			if err := SendEOEP4CrossOnt(ctx, status, config.DefConfig.EthOep4, amt); err != nil {
				log.Errorf("Oep4Circle, SendEOEP4CrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("send %d Oep4 to eth, waiting for confirmation...", config.DefConfig.TxNumPerBatch)
		WaitUntilClean(status)
	}

	status.SetItSuccess(1)
	return true
}

func SendOngToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendOngCrossEth(ctx, status, GetRandAmount(config.DefConfig.OngValLimit, 1)); err != nil {
		log.Errorf("SendOngToEthChain, SendOngCrossEth error: %s", err)
		return false
	}

	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendOngeToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendOngeCrossOnt(ctx, status, config.DefConfig.EthOngx, GetRandAmount(config.DefConfig.OngValLimit, 1)); err != nil {
		log.Errorf("SendOngeToOntChain, SendOngeCrossOnt error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendERC20ToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendERC20CrossOnt(ctx, status, config.DefConfig.EthErc20, GetRandAmount(config.DefConfig.Erc20ValLimit, 1)); err != nil {
		log.Errorf("SendOngeToOntChain, SendOngeCrossOnt error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendOERC20ToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendOERC20CrossEth(ctx, status, config.DefConfig.OntErc20, GetRandAmount(config.DefConfig.Erc20ValLimit, 1)); err != nil {
		log.Errorf("SendOngeToOntChain, SendOngeCrossOnt error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendOEP4ToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendOEP4CrossEth(ctx, status, config.DefConfig.OntOep4, GetRandAmount(config.DefConfig.Oep4ValLimit, 1)); err != nil {
		log.Errorf("SendOEP4ToEthChain, SendOEP4CrossEth error: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendOEP4eToOntChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendEOEP4CrossOnt(ctx, status, config.DefConfig.EthOep4, GetRandAmount(config.DefConfig.Oep4ValLimit, 1)); err != nil {
		log.Errorf("SendOEP4eToOntChain, SendEOEP4CrossOnt error: %s", err)
		return false
	}
	status.SetItSuccess(1)
	return true
}

func SendBtcToEthInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendBtcCrossEth(ctx, status, int64(GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract))); err != nil {
			log.Errorf("SendBtcToEthInBatch, SendBtcCrossEth %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("btc is all received")

	status.SetItSuccess(1)
	return true
}

func SendBtcToOntInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendBtcCrossOnt(ctx, status, int64(GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract))); err != nil {
			log.Errorf("SendBtcToOntInBatch, SendBtcCrossOnt %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("btc is all received")

	status.SetItSuccess(1)
	return true
}

func SendBtceToBtcInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendBtceCrossBtc(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)); err != nil {
			log.Errorf("SendBtceToBtcInBatch, SendBtceCrossBtc %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("ebtc is all received")

	status.SetItSuccess(1)
	return true
}

func SendBtcoToBtcInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendBtcoCrossBtc(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)); err != nil {
			log.Errorf("SendBtcoToBtcInBatch, SendBtcoCrossBtc %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("obtc is all received")

	status.SetItSuccess(1)
	return true
}

func SendBtcoToBtceInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendBtcoCrossBtce(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)); err != nil {
			log.Errorf("SendBtcoToBtceInBatch, SendBtcoCrossBtce %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("obtc is all received")

	status.SetItSuccess(1)
	return true
}

func SendBtceToBtcoInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendBtceCrossBtco(ctx, status, GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)); err != nil {
			log.Errorf("SendBtceToBtcoInBatch, SendBtceCrossBtco %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("ebtc is all received")

	status.SetItSuccess(1)
	return true
}

func SendOntToEthInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendOntCrossEth(ctx, status, GetRandAmount(config.DefConfig.OntValLimit, 1)); err != nil {
			log.Errorf("SendOntToEthInBatch, SendOntCrossEth %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("ont is all received")

	status.SetItSuccess(1)
	return true
}

func SendOnteToOntInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendEOntCrossOnt(ctx, status, config.DefConfig.EthOntx, GetRandAmount(config.DefConfig.OntValLimit, 1)); err != nil {
			log.Errorf("SendOnteToOntInBatch, SendEOntCrossOnt %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("ont is all received")

	status.SetItSuccess(1)
	return true
}

func SendEthToOntInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendEthCrossOnt(ctx, status, GetRandAmount(config.DefConfig.EthValLimit, 1)); err != nil {
			log.Errorf("SendEthToOntInBatch, SendEthCrossOnt %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("eth is all received")

	status.SetItSuccess(1)
	return true
}

func SendEthoToEthInBatch(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	for range tick.C {
		if err := SendEthoCrossEth(ctx, status, config.DefConfig.OntEth, GetRandAmount(config.DefConfig.EthValLimit, 1)); err != nil {
			log.Errorf("SendEthToOntInBatch, SendEthCrossOnt %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("eth is all received")

	status.SetItSuccess(1)
	return true
}

func BtcOntCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	amt := GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)
	cnt := uint64(0)
	tick := time.NewTicker(time.Second * time.Duration(config.DefConfig.BatchInterval))
	defer tick.Stop()

	for range tick.C {
		if err := SendBtcCrossOnt(ctx, status, int64(amt)); err != nil {
			log.Errorf("BtcOntCircle, SendBtcCrossOnt %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("btc is all received")

	cnt = uint64(0)
	for range tick.C {
		if err := SendBtcoCrossBtc(ctx, status, amt); err != nil {
			log.Errorf("BtcOntCircle, SendBtcoCrossBtc %d: %v", cnt+1, err)
			return false
		}
		cnt++
		if cnt == config.DefConfig.BatchTxNum {
			break
		}
	}

	WaitUntilClean(status)
	log.Infof("obtc is all received")

	status.SetItSuccess(1)
	return true
}

func SendBtcToCosmosAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for n := uint64(0); n < config.DefConfig.BatchTxNum; n++ {
		amtArr := make([]uint64, config.DefConfig.TxNumPerBatch)
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			amtArr[i] = GetRandAmount(config.DefConfig.BtcValLimit, config.DefConfig.BtcMinOutputValFromContract)
			if err := SendBtcCrossCosmos(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendBtcToCosmosAndBack, SendBtcCrossCosmos failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("btc->cosmos all received, next send them back")
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			if err := SendBtcFromCosmosToBitcoin(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendBtcToCosmosAndBack, SendBtcFromCosmosToBitcoin failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("btc from cosmos to bitcoin all received: ( batch: %d )", n)
	}

	status.SetItSuccess(1)
	return true
}

func SendEthToCosmosAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for n := uint64(0); n < config.DefConfig.BatchTxNum; n++ {
		amtArr := make([]uint64, config.DefConfig.TxNumPerBatch)
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			amtArr[i] = GetRandAmount(config.DefConfig.EthValLimit, 1)
			if err := SendEthCrossCosmos(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendEthToCosmosAndBack, SendEthCrossCosmos failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("eth->cosmos all received, next send them back")
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			if err := SendEthFromCosmosToEthereum(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendEthToCosmosAndBack, SendEthFromCosmosToEthereum failed: %v", err)
				return false
			}
		}
		WaitUntilClean(status)
		log.Info("eth from cosmos to ethereum all received ( batch: %d )", n)
	}

	status.SetItSuccess(1)
	return true
}

func SendErc20ToCosmosAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for n := uint64(0); n < config.DefConfig.BatchTxNum; n++ {
		amtArr := make([]uint64, config.DefConfig.TxNumPerBatch)
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			amtArr[i] = GetRandAmount(config.DefConfig.Erc20ValLimit, 1)
			if err := SendErc20CrossCosmos(ctx, status, config.DefConfig.EthErc20, amtArr[i]); err != nil {
				log.Errorf("SendErc20ToCosmosAndBack, SendErc20CrossCosmos failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("erc20->cosmos all received, next send them back")
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			if err := SendErc20FromCosmosToEthereum(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendErc20ToCosmosAndBack, SendErc20FromCosmosToEthereum failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("erc20 from cosmos to ethereum all received ( batch: %d )", n)
	}
	status.SetItSuccess(1)
	return true
}

func SendOntToCosmosAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for n := uint64(0); n < config.DefConfig.BatchTxNum; n++ {
		amtArr := make([]uint64, config.DefConfig.TxNumPerBatch)
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			amtArr[i] = GetRandAmount(config.DefConfig.OntValLimit, 1)
			if err := SendOntCrossCosmos(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendOntToCosmosAndBack, SendOntCrossCosmos failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("ont->cosmos all received, next send them back")
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			if err := SendOntFromCosmosToOntology(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendOntToCosmosAndBack, SendOntToCosmosAndBack failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Infof("ont from cosmos to ontology all received ( batch: %d )", n)
	}
	status.SetItSuccess(1)
	return true
}

func SendOngToCosmosAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for n := uint64(0); n < config.DefConfig.BatchTxNum; n++ {
		amtArr := make([]uint64, config.DefConfig.TxNumPerBatch)
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			amtArr[i] = GetRandAmount(config.DefConfig.OngValLimit, 1)
			if err := SendOngCrossCosmos(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendOngToCosmosAndBack, SendOngCrossCosmos failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("ong->cosmos all received, next send them back")
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			if err := SendOngFromCosmosToOntology(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendOngToCosmosAndBack, SendOngFromCosmosToOntology failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Infof("ong from cosmos to ontology all received ( batch: %d )", n)
	}
	status.SetItSuccess(1)
	return true
}

func SendOep4ToCosmosAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for n := uint64(0); n < config.DefConfig.BatchTxNum; n++ {
		amtArr := make([]uint64, config.DefConfig.TxNumPerBatch)
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			amtArr[i] = GetRandAmount(config.DefConfig.Oep4ValLimit, 1)
			if err := SendOep4CrossCosmos(ctx, status, config.DefConfig.OntOep4, amtArr[i]); err != nil {
				log.Errorf("SendOep4ToCosmosAndBack, SendOep4CrossCosmos failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Info("oep4->cosmos all received, next send them back")
		for i := uint64(0); i < config.DefConfig.TxNumPerBatch; i++ {
			if err := SendOep4FromCosmosToOntology(ctx, status, amtArr[i]); err != nil {
				log.Errorf("SendOep4ToCosmosAndBack, SendOep4FromCosmosToOntology failed: %v", err)
				return false
			}
		}

		WaitUntilClean(status)
		log.Infof("oep4 from cosmos to ontology all received ( batch: %d )", n)
	}
	status.SetItSuccess(1)
	return true
}

func SendZeroOntToEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	if err := SendOntCrossEth(ctx, status, 0); err == nil {
		log.Errorf("SendZeroOntToEth, SendOntCrossEth failed: err should not be nil")
		return false
	}
	if err := SendOntCrossCosmos(ctx, status, 0); err == nil {
		log.Errorf("SendZeroOntToEth, SendOntCrossCosmos failed: %v", err)
		return false
	}
	log.Info("all success!")
	status.SetItSuccess(1)
	return true
}

func OntCircleWithoutCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			// ont->eth
			if err := SendOntCrossEth(ctx, status, amt); err != nil {
				log.Errorf("OntCircle, SendOntCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("OntCircle, send %d ont to eth, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendEOntCrossOnt(ctx, status, config.DefConfig.EthOntx, amt); err != nil {
				log.Errorf("OntCircle, SendEOntCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("OntCircle, send %d ont from ethereum to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("OntCircle, ont all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func USDTCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.USDTValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendUSDTCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("USDTCircle, SendUSDTCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("USDTCircle, send %d usdt to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOUSDTCrossEth(ctx, status, amt); err != nil {
				log.Errorf("USDTCircle, SendOUSDTCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("USDTCircle, send %d usdt from ontology to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("USDTCircle, usdt all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func WBTCCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.WBTCValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendWBTCCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("WBTCCircle, SendWBTCCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("WBTCCircle, send %d wbtc to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOWBTCCrossEth(ctx, status, amt); err != nil {
				log.Errorf("WBTCCircle, SendOWBTCCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("WBTCCircle, send %d wbtc from ontology to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("WBTCCircle, wbtc all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func DAICircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.USDTValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendDAICrossOnt(ctx, status, amt); err != nil {
				log.Errorf("DAICircle, SendDAICrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("DAICircle, send %d dai to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendODAICrossEth(ctx, status, amt); err != nil {
				log.Errorf("DAICircle, SendODAICrossEth error: %v", err)
				return false
			}
		}
		log.Infof("DAICircle, send %d dai from ontology to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("DAICircle, dai all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func NeoToEthCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.NeoValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendNeoCrossEth(ctx, status, amt); err != nil {
				log.Errorf("NeoToEthCircle, SendNeoCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("NeoToEthCircle, send %d NEO to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendENeoCrossNeo(ctx, status, amt); err != nil {
				log.Errorf("NeoToEthCircle, SendENeoCrossNeo error: %v", err)
				return false
			}
		}
		log.Infof("NeoToEthCircle, send %d NEO from ethereum to NEO, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("NeoToEthCircle, neo all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func NeoToOntCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.NeoValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendNeoCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("NeoToOntCircle, SendNeoCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("NeoToOntCircle, send %d NEO to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendONeoCrossNeo(ctx, status, amt); err != nil {
				log.Errorf("NeoToOntCircle, SendONeoCrossNeo error: %v", err)
				return false
			}
		}
		log.Infof("NeoToOntCircle, send %d NEO from ontology to NEO, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("NeoToOntCircle, neo all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func OntToNeoAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOntCrossNeo(ctx, status, amt); err != nil {
				log.Errorf("OntToNeoAndBack, SendOntCrossNeo error: %v", err)
				return false
			}
		}
		log.Infof("OntToNeoAndBack, send %d ont to NEO, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendNOntCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("OntToNeoAndBack, SendNOntCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("OntToNeoAndBack, send %d ont from NEO to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("OntToNeoAndBack, ont all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func OntdToEthAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntdValLimit, config.DefConfig.OntdValFloor)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOntdCrossEth(ctx, status, amt); err != nil {
				log.Errorf("OntdToEthAndBack, SendOntdCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("OntdToEthAndBack, send %d ontd to Ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendEOntdCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("OntdToEthAndBack, SendEOntdCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("OntdToEthAndBack, send %d ontd from Ethereum to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("OntdToEthAndBack, ontd all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func OntdToNeoAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntdValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOntdCrossNeo(ctx, status, amt); err != nil {
				log.Errorf("OntdToNeoAndBack, SendOntdCrossNeo error: %v", err)
				return false
			}
		}
		log.Infof("OntdToNeoAndBack, send %d ontd to NEO, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendNOntdCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("OntToNeoAndBack, SendNOntCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("OntdToNeoAndBack, send %d ontd from NEO to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("OntdToNeoAndBack, ontd all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func EOntdToNeoAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntdValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendEOntdCrossNeo(ctx, status, amt); err != nil {
				log.Errorf("EOntdToNeoAndBack, SendEOntdCrossNeo error: %v", err)
				return false
			}
		}
		log.Infof("EOntdToNeoAndBack, send %d ontd to NEO, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendNOntdCrossEth(ctx, status, amt); err != nil {
				log.Errorf("EOntdToNeoAndBack, SendNOntdCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("EOntdToNeoAndBack, send %d ontd from NEO to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("EOntdToNeoAndBack, ontd all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func EthToNeoAndBack(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.OntValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendEthCrossNeo(ctx, status, amt); err != nil {
				log.Errorf("EthToNeoAndBack, SendEthCrossNeo error: %v", err)
				return false
			}
		}
		log.Infof("EthToNeoAndBack, send %d eth to NEO, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendNEthCrossEth(ctx, status, amt); err != nil {
				log.Errorf("EthToNeoAndBack, SendNEthCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("EthToNeoAndBack, send %d eth from NEO to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("EthToNeoAndBack, eth all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func USDCCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.USDCValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendUSDCCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("USDCCircle, SendUSDCCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("USDCCircle, send %d USDC to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendOUSDCCrossEth(ctx, status, amt); err != nil {
				log.Errorf("USDCCircle, SendOUSDCCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("USDCCircle, send %d USDC from ontology to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("USDCCircle, USDC all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}

func RenBTCCircle(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	for i := uint64(0); i < config.DefConfig.BatchTxNum; i++ {
		amt := GetRandAmount(config.DefConfig.RenBTCValLimit, 1)
		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendRenBTCCrossOnt(ctx, status, amt); err != nil {
				log.Errorf("RenBTCCircle, SendRenBTCCrossOnt error: %v", err)
				return false
			}
		}
		log.Infof("RenBTCCircle, send %d renBTC to ontology, waiting for confirmation...", amt)
		WaitUntilClean(status)

		for j := uint64(0); j < config.DefConfig.TxNumPerBatch; j++ {
			if err := SendORenBTCCrossEth(ctx, status, amt); err != nil {
				log.Errorf("RenBTCCircle, SendORenBTCCrossEth error: %v", err)
				return false
			}
		}
		log.Infof("RenBTCCircle, send %d renBTC from ontology to ethereum, waiting for confirmation...", amt)
		WaitUntilClean(status)
		log.Infof("RenBTCCircle, renBTC all received ( batch: %d )", i)
	}

	status.SetItSuccess(1)
	return true
}