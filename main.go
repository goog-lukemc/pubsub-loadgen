// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"cloud.google.com/go/pubsub"
)

var inbound chan *pubsub.Message
var numbers metrics
var wg sync.WaitGroup

func init() {
	inbound = make(chan *pubsub.Message, 1000000)
}

func main() {
	// Create close context
	cctx, close := context.WithCancel(context.Background())

	// Look for a control c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			close()
		}
	}()

	// Validate the commandline.
	valCmdLine()
	wg.Add(1)
	go senderMgr(cctx)

	// Stats Writer
	//statsWriter()
	//setupOpenCen()
	wg.Wait()

}
