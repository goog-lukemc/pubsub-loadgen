package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
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
)

// Commandline checking
func valCmdLine() {
	flag.StringVar(&topicName, "t", "", topicNameUsage)
	flag.StringVar(&msgNamePrefix, "n", "msg-%v", msgNamePrefixUsage)
	flag.IntVar(&maxMsgRoutes, "m", 1, maxMsgRoutesUsage)
	flag.IntVar(&bytesPerMessageBody, "s", 1000, bytesPerMessageBodyUsage)
	flag.Float64Var(&messagesPerSecond, "r", 1000, messagesPerSecondUsage)
	flag.StringVar(&projectid, "p", "", projectIdUsage)
	exampleMsg := flag.Bool("e", false, exampleMessageUsage)
	flag.Parse()
	if *exampleMsg {
		e := &pubsub.Message{Data: []byte(randStringBytesMaskImprSrcSB(bytesPerMessageBody)), Attributes: map[string]string{msgRoute: msgRoute}}
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
