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
package testframework

import (
	"fmt"
	"sync"
	"time"

	poly_go_sdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-io-test/chains/btc"
	"github.com/polynetwork/poly-io-test/chains/cosmos"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/chains/neo"
	"github.com/polynetwork/poly-io-test/chains/ont"
)

//TestFrameworkContext is the context for test case
type TestFrameworkContext struct {
	Framework *TestFramework
	Cases     []TestCase
	RcSdk     *poly_go_sdk.PolySdk
	Status    *CtxStatus
	// invokers
	EthInvoker  *eth.EInvoker
	BscInvoker  *eth.EInvoker
	HecoInvoker *eth.EInvoker
	BtcInvoker  *btc.BtcInvoker
	OntInvoker  *ont.OntInvoker
	CMInvoker   *cosmos.CosmosInvoker
	NeoInvoker  *neo.NeoInvoker
}

//NewTestFrameworkContext return a TestFrameworkContext instance
func NewTestFrameworkContext(fw *TestFramework, caseArr []TestCase, rcSdk *poly_go_sdk.PolySdk, eInvkr *eth.EInvoker, bscInvkr *eth.EInvoker, hecoInvkr *eth.EInvoker,
	btcInvkr *btc.BtcInvoker, ontInvkr *ont.OntInvoker, cmInvoker *cosmos.CosmosInvoker, neoInvoker *neo.NeoInvoker) *TestFrameworkContext {
	ctx := &TestFrameworkContext{
		Framework:   fw,
		Cases:       caseArr,
		RcSdk:       rcSdk,
		EthInvoker:  eInvkr,
		BscInvoker:  bscInvkr,
		HecoInvoker: hecoInvkr,
		BtcInvoker:  btcInvkr,
		OntInvoker:  ontInvkr,
		CMInvoker:   cmInvoker,
		NeoInvoker:  neoInvoker,
	}
	ctx.Status = NewCtxStatus(ctx)
	return ctx
}

type CtxStatus struct {
	lock    *sync.Mutex
	ctx     *TestFrameworkContext
	caseMap map[int]*CaseStatus
}

func NewCtxStatus(ctx *TestFrameworkContext) *CtxStatus {
	return &CtxStatus{
		lock:    &sync.Mutex{},
		ctx:     ctx,
		caseMap: make(map[int]*CaseStatus),
	}
}

func (status *CtxStatus) AddCase(idx int) *CaseStatus {
	status.lock.Lock()
	defer status.lock.Unlock()

	val := NewCaseStatus(idx)
	status.caseMap[idx] = val
	return val
}

func (status *CtxStatus) IsTxPending(tx string) (bool, int) {
	status.lock.Lock()
	defer status.lock.Unlock()

	for i, v := range status.caseMap {
		if ok := v.Has(tx); ok {
			return true, i
		}
	}
	return false, 0
}

func (status *CtxStatus) Del(tx string) {
	status.lock.Lock()
	defer status.lock.Unlock()

	for _, v := range status.caseMap {
		if ok := v.Has(tx); ok {
			v.Del(tx)
			return
		}
	}
}

func (status *CtxStatus) DelWithIndex(tx string, idx int) {
	status.lock.Lock()
	defer status.lock.Unlock()
	status.caseMap[idx].Del(tx)
}

func (status *CtxStatus) GetCaseMap() map[int]*CaseStatus {
	status.lock.Lock()
	defer status.lock.Unlock()
	return status.caseMap
}

func (status *CtxStatus) GetCaseStatus(idx int) *CaseStatus {
	status.lock.Lock()
	defer status.lock.Unlock()
	return status.caseMap[idx]
}

func (status *CtxStatus) Info() map[int]string {
	status.lock.Lock()
	defer status.lock.Unlock()

	res := make(map[int]string)
	for idx, cs := range status.caseMap {
		str := cs.Info()
		if str == "" {
			res[idx] = "no tx for now"
			continue
		}
		if str == "success!" || str == "failed!" {
			res[idx] = str
			continue
		}
		res[idx] = fmt.Sprintf("----------------------------case (index: %d, name: %s)----------------------------\n"+
			"{\n%s}\n%d tx not confirmed", idx, status.ctx.Framework.getTestCaseName(status.ctx.Cases[idx-1]), str, cs.Len())
	}
	return res
}

type TxInfo struct {
	Ty        string
	StartTime time.Time
}

type CaseStatus struct {
	lock      *sync.Mutex
	CaseIdx   int
	txMap     map[string]*TxInfo
	isSuccess int
}

func NewCaseStatus(idx int) *CaseStatus {
	return &CaseStatus{
		lock:      &sync.Mutex{},
		CaseIdx:   idx,
		txMap:     make(map[string]*TxInfo),
		isSuccess: 0,
	}
}

func (cs *CaseStatus) Has(k string) bool {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	_, ok := cs.txMap[k]
	return ok
}

func (cs *CaseStatus) Del(k string) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	delete(cs.txMap, k)
}

func (cs *CaseStatus) BatchDel(keys []string) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	for _, k := range keys {
		delete(cs.txMap, k)
	}
}

func (cs *CaseStatus) AddTx(k string, v *TxInfo) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	cs.txMap[k] = v
}

func (cs *CaseStatus) Info() string {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	info := ""
	if cs.isSuccess == 1 {
		info = "success!"
	} else if cs.isSuccess == -1 {
		info = "failed!"
	} else {
		for k, v := range cs.txMap {
			info += fmt.Sprintf("\t[ txhash: %s, type: %s, sec_not_confirm: %.1f ]\n", k, v.Ty, time.Now().Sub(v.StartTime).Seconds())
		}
	}

	return info
}

func (cs *CaseStatus) GetMapCopy() map[string]*TxInfo {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	cp := make(map[string]*TxInfo)
	for k, v := range cs.txMap {
		cp[k] = v
	}

	return cp
}

func (cs *CaseStatus) Len() int {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	return len(cs.txMap)
}

func (cs *CaseStatus) SetItSuccess(status int) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	cs.isSuccess = status
}
