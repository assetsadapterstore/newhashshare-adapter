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
	"encoding/hex"
	"github.com/assetsadapterstore/newhashshare-adapter/newhashshare_addrdec"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	newhashshare_addrdec.Default.IsTestNet = false

	p2pk, _ := hex.DecodeString("2a54d3864b81df355fd4b7a8e873d2785f1059ee")
	p2pkAddr, _ := newhashshare_addrdec.Default.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)

	p2sh, _ := hex.DecodeString("131a861f0609944596e2d618e41ba8ce07b281d0")
	p2shAddr, _ := newhashshare_addrdec.Default.AddressEncode(p2sh, newhashshare_addrdec.NHSS_mainnetAddressP2SH)
	t.Logf("p2shAddr: %s", p2shAddr)
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	newhashshare_addrdec.Default.IsTestNet = false

	p2pkAddr := "SgZ5WzsQEkjAS9dMxeg42XUTWdyb39J4Kk"
	p2pkHash, _ := newhashshare_addrdec.Default.AddressDecode(p2pkAddr)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))

	//p2shAddr := "sQMG5PncvvxVMrVwXpFfBoi3JFHvPiA9aw"
	//
	//p2shHash, _ := newhashshare_addrdec.Default.AddressDecode(p2shAddr, newhashshare_addrdec.NHSS_mainnetAddressP2SH)
	//t.Logf("p2shHash: %s", hex.EncodeToString(p2shHash))
}

func TestAddressDecoder_ScriptPubKeyToBech32Address(t *testing.T) {

	scriptPubKey, _ := hex.DecodeString("002079db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1")

	addr, err := tw.Decoder.ScriptPubKeyToBech32Address(scriptPubKey)
	if err != nil {
		t.Errorf("ScriptPubKeyToBech32Address failed unexpected error: %v\n", err)
		return
	}
	t.Logf("addr: %s", addr)


	t.Logf("addr: %s", addr)
}


func TestAddressDecoder_VerifyAddress(t *testing.T) {
	newhashshare_addrdec.Default.IsTestNet = false
	check := newhashshare_addrdec.Default.AddressVerify("2GMhkR1uYueeK5cbG5k3XHfW1aSdALjq7Bs")
	t.Logf("check: %v \n", check)
}