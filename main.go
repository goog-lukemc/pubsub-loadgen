package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	// Validate the commandline.
	valCmdLine()

	// Create seed for text creation
	rand.Seed(time.Now().UnixNano())
	psClient, err := pubsub.NewClient(context.Background(), projectid)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Check if the topic exists
	exist, err := psClient.Topic(topicName).Exists(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//If the topic doesnt exist create it.
	if !exist {
		_, err := psClient.CreateTopic(context.Background(), topicName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	// Set a counter so we can adjust our speed control
	cnt := float64(0)

	// Set the start time so we can calcualte the number of messages per minute.
	start := time.Now()

	// Create the message payload
	data := randStringBytesMaskImprSrcSB(bytesPerMessageBody)
	d := &data
	// Loop and send messages
	for {
		// Create a specifice attribute in the message so we can route if needed.
		msgRoute := fmt.Sprintf(msgNamePrefix, rand.Intn(maxMsgRoutes))

		// publish the generated message
		psClient.Topic(topicName).Publish(context.Background(), &pubsub.Message{
			Data:       []byte(*d),
			Attributes: map[string]string{msgRoute: msgRoute},
		})

		// Incremen the counter
		cnt++

		// calculate the elasped time
		elapSec := time.Since(start).Seconds()

		// caculate the message per second so far
		perSecond := cnt / elapSec

		// Check is our message rate per second is higher than what we can normall sustain.
		if perSecond > messagesPerSecond {
			// If it is levelize by causing the loop to sleep
			time.Sleep(time.Millisecond * time.Duration(perSecond-messagesPerSecond))
		}

	}
}

// Borrowed heavely from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()_+"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
