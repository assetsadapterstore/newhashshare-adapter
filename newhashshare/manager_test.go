/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package newhashshare

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"path/filepath"
	"testing"
	"time"

	"github.com/codeskyblue/go-sh"
	"github.com/shopspring/decimal"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "conf.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	//wm.ExplorerClient.Debug = false
	wm.WalletClient.Debug = true
	return wm
}

func TestWalletManager(t *testing.T) {

	t.Log("Symbol:", tw.Config.Symbol)
	t.Log("ServerAPI:", tw.Config.ServerAPI)
}

//func TestImportPrivKey(t *testing.T) {
//
//	tests := []struct {
//		seed []byte
//		name string
//		tag  string
//	}{
//		{
//			seed: tw.GenerateSeed(),
//			name: "Chance",
//			tag:  "first",
//		},
//		{
//			seed: tw.GenerateSeed(),
//			name: "Chance",
//			tag:  "second",
//		},
//	}
//
//	for i, test := range tests {
//		key, err := keystore.NewHDKey(test.seed, test.name, "m/44'/88'")
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		privateKey, err := key.MasterKey.ECPrivKey()
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		publicKey, err := key.MasterKey.ECPubKey()
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		wif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		t.Logf("Privatekey wif[%d] = %s\n", i, wif.String())
//
//		address, err := btcutil.NewAddressPubKey(publicKey.SerializeCompressed(), &chaincfg.MainNetParams)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		t.Logf("Privatekey address[%d] = %s\n", i, address.EncodeAddress())
//
//		//解锁钱包
//		err = tw.UnlockWallet("1234qwer", 120)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//		}
//
//		//导入私钥
//		err = tw.ImportPrivKey(wif.String(), test.name)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//		} else {
//			t.Logf("ImportPrivKey[%d] success \n", i)
//		}
//	}
//
//}

func TestGetCoreWalletinfo(t *testing.T) {
	tw.GetCoreWalletinfo()
}

func TestKeyPoolRefill(t *testing.T) {

	//解锁钱包
	err := tw.UnlockWallet("1234qwer", 120)
	if err != nil {
		t.Errorf("KeyPoolRefill failed unexpected error: %v\n", err)
	}

	err = tw.KeyPoolRefill(10000)
	if err != nil {
		t.Errorf("KeyPoolRefill failed unexpected error: %v\n", err)
	}
}

func TestCreateReceiverAddress(t *testing.T) {

	tests := []struct {
		account string
		tag     string
	}{
		{
			account: "john",
			tag:     "normal",
		},
		//{
		//	account: "Chance",
		//	tag:     "normal",
		//},
	}

	for i, test := range tests {

		a, err := tw.CreateReceiverAddress(test.account)
		if err != nil {
			t.Errorf("CreateReceiverAddress[%d] failed unexpected error: %v", i, err)
		} else {
			t.Logf("CreateReceiverAddress[%d] address = %v", i, a)
		}

	}

}

func TestGetAddressesByAccount(t *testing.T) {
	addresses, err := tw.GetAddressesByAccount("")
	if err != nil {
		t.Errorf("GetAddressesByAccount failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("GetAddressesByAccount address[%d] = %s\n", i, a)
	}
}

func TestCreateBatchAddress(t *testing.T) {
	_, _, err := tw.CreateBatchAddress("WNafusXqpykohfnSKqcsR1y5xgDCdqHYxh", "1234qwer", 2)
	if err != nil {
		t.Errorf("CreateBatchAddress failed unexpected error: %v\n", err)
		return
	}
}

func TestEncryptWallet(t *testing.T) {
	err := tw.EncryptWallet("11111111")
	if err != nil {
		t.Errorf("EncryptWallet failed unexpected error: %v\n", err)
		return
	}
}

func TestUnlockWallet(t *testing.T) {
	err := tw.UnlockWallet("1234qwer", 1)
	if err != nil {
		t.Errorf("UnlockWallet failed unexpected error: %v\n", err)
		return
	}
}

func TestCreateNewWallet(t *testing.T) {
	_, _, err := tw.CreateNewWallet("Block", "1234qwer")
	if err != nil {
		t.Errorf("CreateNewWallet failed unexpected error: %v\n", err)
		return
	}
}

func TestGetWalletKeys(t *testing.T) {
	wallets, err := tw.GetWallets()
	if err != nil {
		t.Errorf("GetWalletKeys failed unexpected error: %v\n", err)
		return
	}

	for i, w := range wallets {
		t.Logf("GetWalletKeys wallet[%d] = %v", i, w)
	}
}

func TestGetWalletBalance(t *testing.T) {

	tests := []struct {
		name string
		tag  string
	}{
		{
			name: "W4ruoAyS5HdBMrEeeHQTBxo4XtaAixheXQ",
			tag:  "first",
		},
		{
			name: "Wallet Test",
			tag:  "second",
		},
		{
			name: "*",
			tag:  "all",
		},
		{
			name: "llllll",
			tag:  "account not exist",
		},
	}

	for i, test := range tests {
		balance := tw.GetWalletBalance(test.name)
		t.Logf("GetWalletBalance[%d] %s balance = %s \n", i, test.name, balance)
	}

}

func TestCreateNewPrivateKey(t *testing.T) {

	test := struct {
		name     string
		password string
		tag      string
	}{
		name:     "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek",
		password: "1234qwer",
	}

	count := 100

	w, err := tw.GetWalletInfo(test.name)
	if err != nil {
		t.Errorf("CreateNewPrivateKey failed unexpected error: %v\n", err)
		return
	}

	key, err := w.HDKey(test.password)
	if err != nil {
		t.Errorf("CreateNewPrivateKey failed unexpected error: %v\n", err)
		return
	}

	timestamp := 1
	t.Logf("CreateNewPrivateKey timestamp = %v \n", timestamp)

	derivedPath := fmt.Sprintf("%s/%d", key.RootPath, timestamp)
	childKey, _ := key.DerivedKeyWithPath(derivedPath, tw.Config.CurveType)

	for i := 0; i < count; i++ {

		wif, a, err := tw.CreateNewPrivateKey(key.KeyID, childKey, derivedPath, uint64(i))
		if err != nil {
			t.Errorf("CreateNewPrivateKey[%d] failed unexpected error: %v\n", i, err)
			continue
		}

		t.Logf("CreateNewPrivateKey[%d] wif = %v \n", i, wif)
		t.Logf("CreateNewPrivateKey[%d] address = %v \n", i, a.Address)
	}
}

func TestGetWalleInfo(t *testing.T) {
	w, err := tw.GetWalletInfo("Zhiquan Test")
	if err != nil {
		t.Errorf("GetWalletInfo failed unexpected error: %v\n", err)
		return
	}

	t.Logf("GetWalletInfo wallet = %v \n", w)
}

//func TestCreateBatchPrivateKey(t *testing.T) {
//
//	w, err := tw.GetWalletInfo("Zhiquan Test")
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	key, err := w.HDKey("1234qwer")
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	wifs, err := tw.CreateBatchPrivateKey(key, 10000)
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	for i, wif := range wifs {
//		t.Logf("CreateBatchPrivateKey[%d] wif = %v \n", i, wif)
//	}
//
//}

//func TestImportMulti(t *testing.T) {
//
//	addresses := []string{
//		"1CoRcQGjPEyWmB1ZyG6CEDN3SaMsaD3ERa",
//		"1ESGCsXkNr3h5wvWScdCpVHu2GP3KJtCdV",
//	}
//
//	keys := []string{
//		"L5k8VYSvuZxC5FCczGVC8MmnKKix3Mcs6t185eUJVKTzZb1f6bsX",
//		"L3RVDjPVBSc7DD4WtmzbHkAHJW4kDbyXbw4vBppZ4DRtPt5u8Naf",
//	}
//
//	UnlockWallet("1234qwer", 120)
//	failed, err := ImportMulti(addresses, keys, "Zhiquan Test")
//	if err != nil {
//		t.Errorf("ImportMulti failed unexpected error: %v\n", err)
//	} else {
//		t.Errorf("ImportMulti result: %v\n", failed)
//	}
//}

func TestBackupWallet(t *testing.T) {

	backupFile, err := tw.BackupWallet("W9JyC464XAZEJgdiAZxUXbPpsZZ2JeAujV")
	if err != nil {
		t.Errorf("BackupWallet failed unexpected error: %v\n", err)
	} else {
		t.Errorf("BackupWallet filePath: %v\n", backupFile)
	}
}

func TestBackupWalletData(t *testing.T) {
	tw.Config.WalletDataPath = "/home/www/btc/testdata/testnet3/"
	tmpWalletDat := fmt.Sprintf("tmp-walllet-%d.dat", time.Now().Unix())
	backupFile := filepath.Join(tw.Config.WalletDataPath, tmpWalletDat)
	err := tw.BackupWalletData(backupFile)
	if err != nil {
		t.Errorf("BackupWallet failed unexpected error: %v\n", err)
	} else {
		t.Errorf("BackupWallet filePath: %v\n", backupFile)
	}
}

func TestDumpWallet(t *testing.T) {
	tw.UnlockWallet("1234qwer", 120)
	file := filepath.Join(".", "openwallet", "")
	err := tw.DumpWallet(file)
	if err != nil {
		t.Errorf("DumpWallet failed unexpected error: %v\n", err)
	} else {
		t.Errorf("DumpWallet filePath: %v\n", file)
	}
}

func TestGOSH(t *testing.T) {
	//text, err := sh.Command("go", "env").Output()
	//text, err := sh.Command("wmd", "version").Output()
	text, err := sh.Command("wmd", "Config", "see", "-s", "btm").Output()
	if err != nil {
		t.Errorf("GOSH failed unexpected error: %v\n", err)
	} else {
		t.Errorf("GOSH output: %v\n", string(text))
	}
}

func TestGetBlockChainInfo(t *testing.T) {
	b, err := tw.GetBlockChainInfo()
	if err != nil {
		t.Errorf("GetBlockChainInfo failed unexpected error: %v\n", err)
	} else {
		t.Logf("GetBlockChainInfo info: %v\n", b)
	}
}

func TestListUnspent(t *testing.T) {
	utxos, err := tw.ListUnspent(0, "GJ2o7XyQbTYkXC3VmKfQoThZuAZqytZATj")
	if err != nil {
		t.Errorf("ListUnspent failed unexpected error: %v\n", err)
		return
	}
	totalBalance := decimal.Zero
	for _, u := range utxos {
		t.Logf("ListUnspent %s: %s = %s\n", u.Address, u.AccountID, u.Amount)
		amount, _ := decimal.NewFromString(u.Amount)
		totalBalance = totalBalance.Add(amount)
	}

	t.Logf("totalBalance: %s \n", totalBalance.String())
}

func TestEstimateFee(t *testing.T) {
	feeRate, _ := tw.EstimateFeeRate()
	t.Logf("EstimateFee feeRate = %s\n", feeRate.StringFixed(8))
	fees, _ := tw.EstimateFee(10, 2, feeRate)
	t.Logf("EstimateFee fees = %s\n", fees.StringFixed(8))
}

func TestGetNetworkInfo(t *testing.T) {
	tw.GetNetworkInfo()
}

func TestPrintConfig(t *testing.T) {
	tw.Config.PrintConfig()
}

func TestRestoreWallet(t *testing.T) {
	keyFile := "/myspace/workplace/go-workspace/projects/bin/data/btc/key/MacOS-W9JyC464XAZEJgdiAZxUXbPpsZZ2JeAujV.key"
	dbFile := "/myspace/workplace/go-workspace/projects/bin/data/btc/db/MacOS-W9JyC464XAZEJgdiAZxUXbPpsZZ2JeAujV.db"
	datFile := "/myspace/workplace/go-workspace/projects/bin/testdatfile/wallet.dat"
	tw.LoadConfig()
	err := tw.RestoreWallet(keyFile, dbFile, datFile, "1234qwer")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
	}

}

func TestWalletManager_ImportAddress(t *testing.T) {
	addr := "Ga2thK76EF4Y1q14RtmCfBZepC2GYBvaCy"
	err := tw.ImportAddress(addr, "")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
		return
	}
	log.Info("imported success")
}

func TestWalletManager_ListAddresses(t *testing.T) {
	addresses, err := tw.ListAddresses()
	if err != nil {
		t.Errorf("GetAddressesByAccount failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("ListAddresses address[%d] = %s\n", i, a)
	}
}

func TestWalletManager_GetInfo(t *testing.T) {
	tw.GetInfo()
}