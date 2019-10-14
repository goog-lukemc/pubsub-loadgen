package main

import (
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

//var jexp JSONExporter

func setupOpenCen() option.ClientOption {

	// Register stats and trace exporters to export
	// the collected data to the console.
	// view.RegisterExporter(exp)
	// Register the view to collect gRPC client stats.

	view.RegisterExporter(newPubSubExporter())

	//Set the repoting period to 1 second
	view.SetReportingPeriod(time.Second * 2)
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Fatal(err)
	}

	if err := view.Register(pubsub.DefaultPublishViews...); err != nil {
		log.Fatal(err)
	}

	//view.SetReportingPeriod(time.Second)
	grpcDialOption := grpc.WithStatsHandler(&ocgrpc.ClientHandler{})

	go writeEvents()

	return option.WithGRPCDialOption(grpcDialOption)

}
