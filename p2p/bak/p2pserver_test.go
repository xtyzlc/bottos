// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

// This program is free software: you can distribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Bottos.  If not, see <http://www.gnu.org/licenses/>.

// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exec provides functions for executing WebAssembly bytecode.

/*
 * file description: the interface for WASM execution
 * @Author: Stewart Li
 * @Date:   2018-02-08
 * @Last Modified by: Stewart Li
 * @Last Modified time: 2018-05-30
 */

package p2pserver

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/bottos-project/bottos/common/types"
	"github.com/bottos-project/bottos/config"
)

func TestP2PServ(t *testing.T) {
	fmt.Println("p2p_server::TestP2PServ")

	if TST == 0 {
		err := config.LoadConfig()
		if err != nil {
			fmt.Println("Load config fail")
			os.Exit(1)
		}
	}

	p2p := NewServ()
	p2p.Start()
}

func TestTrxSend(t *testing.T) {

	p2pconfig := ReadFile(CONF_FILE)

	addrPort := p2pconfig.PeerLst[0] + ":" + fmt.Sprint(p2pconfig.ServPort)
	conn, err := net.Dial("tcp", addrPort)
	if err != nil {
		fmt.Println("*ERROR* Failed to create a connection for remote server !!! err: ", err)
		return
	}

	type message struct {
		Src     string
		Dst     string
		MsgType uint8
		Content []byte
	}

	trx := &types.Transaction{
		Version:     1,
		CursorNum:   1,
		CursorLabel: 1,
		Lifetime:    1,
		Sender:      "Trump",
		Contract:    "Check",
		Method:      "Func1",
		Param:       nil,
		SigAlg:      1,
		Signature:   []byte{},
	}

	byteTrx, err := json.Marshal(trx)
	if err != nil {
		fmt.Println("*ERROR* Failed to package the message : ", err)
		return
	}

	msg := message{
		Src:     p2pconfig.ServAddr,
		Dst:     p2pconfig.PeerLst[0],
		MsgType: CRX_BROADCAST,
		Content: byteTrx,
	}

	byteMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("*ERROR* Failed to package the message : ", err)
	}

	len, err := conn.Write(byteMsg)
	if err != nil {
		fmt.Println("*ERROR* Failed to send data to the remote server addr !!! err: ", err)
	} else if len < 0 {
		fmt.Println("*ERROR* Failed to send data to the remote server addr !!! err: ", err)
	}
}
