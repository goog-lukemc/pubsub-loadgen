package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	// SENTON is used to tag the message with a sent timestamp
	SENTON = "sent_on"
)

func senderMgr(ctx context.Context) {
	// Create seed for text creation
	rand.Seed(time.Now().UnixNano())
	psSendClient, err := pubsub.NewClient(ctx, projectid, setupOpenCen())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Check if the topic exists
	exist, err := psSendClient.Topic(topicName).Exists(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//If the topic doesnt exist create it.
	if !exist {
		_, err := psSendClient.CreateTopic(context.Background(), topicName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	numbers.updateSendTimer(int64(float64(time.Second.Nanoseconds()) / messagesPerSecond))
	schedule := time.NewTicker(time.Nanosecond * time.Duration(numbers.sendTimer))
	go sendWorker(schedule, psSendClient)
	for {
		select {
		case <-ctx.Done():
			schedule.Stop()
			defer wg.Done()
			return
		}

	}
}

func sendWorker(sched *time.Ticker, client *pubsub.Client) {
	// Create the message payload
	data := randStringBytesMaskImprSrcSB(bytesPerMessageBody)
	d := &data
	firstRun := true
	pubsub.DefaultPublishSettings.DelayThreshold = time.Second
	pubsub.DefaultPublishSettings.CountThreshold = int(messagesPerSecond)
	for tick := range sched.C {
		mp := make(map[string]string)
		for k, v := range prepedAttributes {
			mp[k] = v
		}
		mp[SENTON] = tick.String()
		if firstRun {
			// publish the generated message
			rs := client.Topic(topicName).Publish(context.Background(), &pubsub.Message{
				Data:       []byte(*d),
				Attributes: mp,
			})

			_, err := rs.Get(context.Background())
			if err != nil {
				panic(err)
			}
			firstRun = false
		} else {
			client.Topic(topicName).Publish(context.Background(), &pubsub.Message{
				Data:       []byte(*d),
				Attributes: mp,
			})
		}
		numbers.addSent()
	}
}
