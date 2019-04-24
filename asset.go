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
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
)

// Static text resources
const (
	msgRoute                 = "msgroute"
	topicNameUsage           = "Name of the topic to connect to.  If the topic is not found it will be created.(Required)"
	msgNamePrefixUsage       = "Create a prefix to the message attribute.  Default is msg-(random). (Optional)"
	maxMsgRoutesUsage        = "Used if you would like to use a message attribute for routing simulation. (Optional)"
	bytesPerMessageBodyUsage = "The size of the data body for the message in bytes. (Optional)"
	messagesPerSecondUsage   = "The number of message you would like to generate per second. (Optional)"
	projectIdUsage           = "The topic's projectid. (Required)"
	exampleMessageUsage      = "Generates and prints an example message (Optional)"
	attributeTagsUsage       = "A comma seperated list of attribute names:value (Optional)\n Example  -t myattribute:myvalue,myattribute2:value2"
)

// Globals
var (
	topicName           string
	msgNamePrefix       string
	maxMsgRoutes        int
	bytesPerMessageBody = 1000
	src                 = rand.NewSource(time.Now().UnixNano())
	messagesPerSecond   float64
	projectid           string
	attributeTags       string
	prepedAttributes    map[string]string
)

// Commandline checking
func valCmdLine() {
	flag.StringVar(&topicName, "t", "", topicNameUsage)
	flag.StringVar(&msgNamePrefix, "n", "msg", msgNamePrefixUsage)
	flag.IntVar(&maxMsgRoutes, "m", 1, maxMsgRoutesUsage)
	flag.IntVar(&bytesPerMessageBody, "s", 1000, bytesPerMessageBodyUsage)
	flag.Float64Var(&messagesPerSecond, "r", 1000, messagesPerSecondUsage)
	flag.StringVar(&projectid, "p", "", projectIdUsage)
	flag.StringVar(&attributeTags, "g", "", attributeTagsUsage)
	exampleMsg := flag.Bool("e", false, exampleMessageUsage)
	flag.Parse()
	parseTags()
	if *exampleMsg {
		e := &pubsub.Message{Data: []byte(randStringBytesMaskImprSrcSB(bytesPerMessageBody)), Attributes: prepedAttributes}
		bts, _ := json.Marshal(e)
		fmt.Println(string(bts))
		os.Exit(0)
	}
	if projectid == "" {
		fmt.Println("The ProjectId is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if topicName == "" {
		fmt.Println("The topic must be set")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func parseTags() {
	// Make a map to populate
	attributes := make(map[string]string)

	// If it is an empty paramater return now
	if len(attributeTags) == 0 {
		prepedAttributes = attributes
		return
	}

	//Split the parameter tag on a comma
	tags := strings.Split(attributeTags, ",")

	//Loop through them and see get the process the values
	for _, tag := range tags {
		data := strings.Split(tag, ":")
		if len(data) == 2 {
			attributes[data[0]] = data[1]
		} else {
			attributes[data[0]] = ""
		}

	}

	prepedAttributes = attributes
	return
}
