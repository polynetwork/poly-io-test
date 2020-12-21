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
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/polynetwork/poly-io-test/chains/btc"
	"github.com/polynetwork/poly-io-test/chains/cosmos"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/chains/neo"
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"

	poly_go_sdk "github.com/polynetwork/poly-go-sdk"
)

//Default TestFramework instance
var TFramework = NewTestFramework()

//TestCase type
type TestCase func(ctx *TestFrameworkContext, status *CaseStatus) bool

//TestFramework manage test case and run test case
type TestFramework struct {
	//Test case start time
	startTime time.Time
	//hold the test case for testing
	testCases []TestCase
	//Map the test case id to test case name
	testCaseNameMap map[string]string
	//Map test case id to test case
	testCasesMap map[string]TestCase
	//Map the test case result for testing
	testCaseRes map[string]bool
	//relayer chain sdk object
	rcSdk *poly_go_sdk.PolySdk
	// invokers
	ethInvoker    *eth.EInvoker
	bscInvoker    *eth.EInvoker
	btcInvoker    *btc.BtcInvoker
	ontInvoker    *ont.OntInvoker
	cosmosInvoker *cosmos.CosmosInvoker
	neoInvoker    *neo.NeoInvoker
}

//NewTestFramework return a TestFramework instance
func NewTestFramework() *TestFramework {
	return &TestFramework{
		testCases:       make([]TestCase, 0),
		testCaseNameMap: make(map[string]string, 0),
		testCasesMap:    make(map[string]TestCase, 0),
		testCaseRes:     make(map[string]bool, 0),
	}
}

//RegTestCase register a test case to framework
func (this *TestFramework) RegTestCase(name string, testCase TestCase) {
	this.testCases = append(this.testCases, testCase)
	testCaseId := this.getTestCaseId(testCase)
	this.testCaseNameMap[testCaseId] = name
	this.testCasesMap[testCaseId] = testCase
}

//Start run test case
func (this *TestFramework) Start(testCases []string, loopNumber int) {
	if len(testCases) > 0 {
		taseCaseList := make([]TestCase, 0, len(testCases))
		for _, t := range testCases {
			if t == "" {
				continue
			}
			testCase := this.getTestCaseByName(t)
			if testCase != nil {
				taseCaseList = append(taseCaseList, testCase)
			}
		}
		if len(taseCaseList) > 0 {
			this.runTestList(taseCaseList, loopNumber)
		}
		log.Info("No test case to run")
		os.Exit(1)
	}

	this.runTestList(this.testCases, loopNumber)
}

func (this *TestFramework) RunOnce(testCases []string) {
	this.Start(testCases, 1)
}

func (this *TestFramework) Run(testCases []string, loopNum int) {
	this.Start(testCases, loopNum)
}

func (this *TestFramework) runTestList(testCaseList []TestCase, loopNumber int) {
	this.onTestStart()
	defer this.onTestFinish(testCaseList)

	ctx := NewTestFrameworkContext(this, testCaseList, this.rcSdk, this.ethInvoker, this.bscInvoker, this.btcInvoker,
		this.ontInvoker, this.cosmosInvoker, this.neoInvoker)
	if this.ontInvoker != nil {
		go MonitorOnt(ctx)
	}
	go MonitorRChain(ctx)
	if this.ethInvoker != nil {
		go MonitorEthLikeChain(ctx, config.DefConfig.EthChainID)
	}
	if this.bscInvoker != nil {
		go MonitorEthLikeChain(ctx, config.DefConfig.BscChainID)
	}
	if this.btcInvoker != nil {
		go MonitorBtc(ctx)
	}
	if this.cosmosInvoker != nil {
		go MonitorCosmos(ctx)
	}
	if this.neoInvoker != nil {
		go MonitorNeo(ctx)
	}
	go ReportPending(ctx)

	wg := &sync.WaitGroup{}
	for i, testCase := range testCaseList {
		wg.Add(1)
		go this.runTest(i+1, ctx, testCase, loopNumber, wg)
	}

	wg.Wait()
}

//Run a single test case
func (this *TestFramework) runTest(index int, ctx *TestFrameworkContext, testCase TestCase, loopNum int, wg *sync.WaitGroup) {
	for i := 0; i < loopNum; i++ {
		this.onBeforeTestCaseStart(index, loopNum, testCase)
		status := ctx.Status.AddCase(index)
		ok := testCase(ctx, status)
		this.onAfterTestCaseFinish(index, loopNum, testCase, ok)
		this.testCaseRes[this.getTestCaseId(testCase)] = ok
		if !ok {
			status.SetItSuccess(-1)
			log.Errorf("case %s failed (loop: %d)", this.getTestCaseName(testCase), i)
			break
		}
	}
	wg.Done()
}

//SetRcSdk relaye chain sdk instance to test framework
func (this *TestFramework) SetRcSdk(rcSdk *poly_go_sdk.PolySdk) {
	this.rcSdk = rcSdk
}

//SetETH instance to test framework
func (this *TestFramework) SetEthInvoker(invoker *eth.EInvoker) {
	this.ethInvoker = invoker
}

//SetBSC instance to test framework
func (this *TestFramework) SetBSCInvoker(invoker *eth.EInvoker) {
	this.bscInvoker = invoker
}

//SetBtcCli instance to test framework
func (this *TestFramework) SetBtcInvoker(invoker *btc.BtcInvoker) {
	this.btcInvoker = invoker
}

func (this *TestFramework) SetOntInvoker(invoker *ont.OntInvoker) {
	this.ontInvoker = invoker
}

func (this *TestFramework) SetCosmosInvoker(cmInvoker *cosmos.CosmosInvoker) {
	this.cosmosInvoker = cmInvoker
}

func (this *TestFramework) SetNeoInvoker(neoInvoker *neo.NeoInvoker) {
	this.neoInvoker = neoInvoker
}

//onTestStart invoke at the beginning of test
func (this *TestFramework) onTestStart() {
	version, _ := this.rcSdk.GetVersion()
	log.Info("===============================================================")
	log.Infof("-------CrossChain Test Start Version: %s", version)
	log.Info("===============================================================")
	log.Info("")
	this.startTime = time.Now()
	str := ""
	//if this.btcInvoker != nil {
	//	btcInfo, err := this.btcInvoker.GetAccInfo()
	//	if err != nil {
	//		panic(err)
	//	}
	//	str += btcInfo + "\n"
	//}
	//if this.ethInvoker != nil {
	//	ethInfo, err := this.ethInvoker.GetAccInfo()
	//	if err != nil {
	//		panic(err)
	//	}
	//	str += ethInfo + "\n"
	//}
	//if this.ontInvoker != nil {
	//	ontInfo, err := this.ontInvoker.GetAccInfo()
	//	if err != nil {
	//		panic(err)
	//	}
	//	str += ontInfo + "\n"
	//}
	//if this.cosmosInvoker != nil {
	//	cmInfo, err := this.cosmosInvoker.GetAccInfo()
	//	if err != nil {
	//		panic(err)
	//	}
	//	str += cmInfo + "\n"
	//}

	log.Infof("account info: {\n %s}", str)
}

//onTestStart invoke at the end of test
func (this *TestFramework) onTestFinish(testCaseList []TestCase) {
	failedList := make([]string, 0)
	successList := make([]string, 0)
	for testCase, ok := range this.testCaseRes {
		if ok {
			successList = append(successList, this.getTestCaseName(testCase))
		} else {
			failedList = append(failedList, this.getTestCaseName(testCase))
		}
	}

	skipList := make([]string, 0)
	for _, testCase := range testCaseList {
		_, ok := this.testCaseRes[this.getTestCaseId(testCase)]
		if !ok {
			skipList = append(skipList, this.getTestCaseName(testCase))
		}
	}

	succCount := len(successList)
	failedCount := len(failedList)

	log.Info("===============================================================")
	log.Infof("CrossChain Test Finish Total:%v Success:%v Failed:%v Skip:%v TimeCost:%.2f s.",
		len(this.testCases),
		succCount,
		failedCount,
		len(this.testCases)-succCount-failedCount,
		time.Now().Sub(this.startTime).Seconds())
	if succCount > 0 {
		log.Info("---------------------------------------------------------------")
		log.Info("Success list:")
		for i, succCase := range successList {
			log.Infof("%d.\t%s", i+1, succCase)
		}
	}
	if failedCount > 0 {
		log.Info("---------------------------------------------------------------")
		log.Info("Fail list:")
		for i, failCase := range failedList {
			log.Infof("%d.\t%s", i+1, failCase)
		}
	}
	if len(skipList) > 0 {
		log.Info("---------------------------------------------------------------")
		log.Info("Skip list:")
		for i, failCase := range skipList {
			log.Infof("%d.\t%s", i+1, failCase)
		}
	}
	log.Info("===============================================================")
	time.Sleep(time.Second * time.Duration(config.DefConfig.ReportInterval+1))
	os.Exit(0)
}

//onTestFailNow invoke when context.FailNow() was be called
func (this *TestFramework) onTestFailNow() {
	log.Info("Test Stop.")
}

//onBeforeTestCaseStart invoke before single test case
func (this *TestFramework) onBeforeTestCaseStart(index, loop int, testCase TestCase) {
	log.Info("===============================================================")
	log.Infof("%d. Start TestCase:%s (loop_time: %d)", index, this.getTestCaseName(testCase), loop)
	log.Info("---------------------------------------------------------------")
}

//onBeforeTestCaseStart invoke after single test case
func (this *TestFramework) onAfterTestCaseFinish(index, loop int, testCase TestCase, res bool) {
	log.Info("---------------------------------------------------------------")
	if res {
		log.Infof("TestCase: (index: %d, name: %s, loop_time: %d) success.", index,
			this.getTestCaseName(testCase), loop)
	} else {
		log.Infof("TestCase: (index: %d, name: %s, loop_time: %d) failed.", index,
			this.getTestCaseName(testCase), loop)
	}
	log.Info("===============================================================")
	log.Info("")
}

//getTestCaseName return the name of test case
func (this *TestFramework) getTestCaseName(testCase interface{}) string {
	testCaseStr, ok := testCase.(string)
	if !ok {
		testCaseStr = this.getTestCaseId(testCase)
	}
	name, ok := this.testCaseNameMap[testCaseStr]
	if ok {
		return name
	}
	return ""
}

//getTestCaseByName return test case by test case name
func (this *TestFramework) getTestCaseByName(name string) TestCase {
	testCaseId := ""
	for id, n := range this.testCaseNameMap {
		if n == name {
			testCaseId = id
			break
		}
	}
	if testCaseId == "" {
		return nil
	}
	return this.testCasesMap[testCaseId]
}

//getTestCaseId return the id of test case
func (this *TestFramework) getTestCaseId(testCase interface{}) string {
	return fmt.Sprintf("%v", reflect.ValueOf(testCase).Pointer())
}
