# Cloud Pubsub Volume Tester

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Prerequisites

Go 1.9 or higher

## Installing

Clone this repo

```
go build -o pubsubvt
```

```
pubsubvt -p <<projectid>> -t <<my-test-topic>>
```
## Commandline options
  * -e
  >>  Generates and prints an example message (Optional)
  * -m int
  >> 	Used if you would like to use a message attribute for routing simulation.  It will randomly select a number from 0 to this value. (Optional) (default 1)
  * -n string
  >>  Create a prefix to the message attribute.  Default is msg-(random). (Optional) (default "msg-%v")
  * -p string
  >> 	The topic's projectid. (Required)
  * -r float
  >>  The number of message you would like to generate per second. (Optional) (default 1000)
  * -s int
  >>  The size of the data body for the message in bytes. (Optional) (default 1000)
  * -t string
  >> Name of the topic to connect to.  If the topic is not found it will be created.(Required)
