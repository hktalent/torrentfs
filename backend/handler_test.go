// Copyright 2020 The CortexTheseus Authors
// This file is part of the CortexTheseus library.
//
// The CortexTheseus library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The CortexTheseus library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the CortexTheseus library. If not, see <http://www.gnu.org/licenses/>.

package backend

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/CortexFoundation/torrentfs/params"
)

func TestGetFile(t *testing.T) {
	params.DefaultConfig.DataDir = "testdata"
	params.DefaultConfig.Port = 0
	params.DefaultConfig.Mode = "LAZY"
	ih := "aea5584d0cd3865e90c80eace3bfcb062473d966"
	fmt.Println(params.DefaultConfig)
	tm, _ := NewTorrentManager(&params.DefaultConfig, 1, false, false)
	//tm.Simulate()
	tm.Start()
	tm.Search(context.Background(), ih, 0)
	defer tm.Close()
	time.Sleep(30 * time.Second)
	//a, _, _, _ := tm.Available(ih, 100000000)
	//fmt.Println("available", a)
	file, _, _ := tm.GetFile(ih, "data")
	if file == nil {
		log.Fatal("failed to get file")
	}
	fmt.Println("file", file[:20])
}
