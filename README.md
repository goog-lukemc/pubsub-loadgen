# Cloud Pubsub Volume Tester

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
  >>  Create a prefix to the message attribute.  Default is msg-(random). (Optional) (default "msg-")
  * -p string
  >> 	The topic's projectid. (Required)
  * -r float
  >>  The number of message you would like to generate per second. (Optional) (default 1000)
  * -s int
  >>  The size of the data body for the message in bytes. (Optional) (default 1000)
  * -t string
  >> Name of the topic to connect to.  If the topic is not found it will be created.(Required)

## Limitations
pubsubvt is only message volume is bound by the CPU, RAM, Network resource avalibility.  A 2 core system with 1.4 gigs of RAM can generate roughly 12k message per second.  If your requirements are higher, add resources to the host OS or add worker nodes and adjust commandline paramters accordingly. 