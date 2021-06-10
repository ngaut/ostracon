package main

import (
	"flag"
	"github.com/nathanieltornow/ostracon/recordshard"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	storagePath  = flag.String("storagePath", "tmp", "Path to storage directory")
	ipAddr       = flag.String("ipAddr", "", "Ip-Address of seqshard")
	parentIpAddr = flag.String("parentIpAddr", "", "Address of parent seqshard")
)

func main() {
	flag.Parse()
	recShard, err := recordshard.NewRecordShard(*storagePath, time.Microsecond*100)
	if err != nil {
		logrus.Fatalln("Failed creating seqshard")
	}

	err = recShard.Start(*ipAddr, *parentIpAddr)
	if err != nil {
		logrus.Fatalln("Failed starting seqshard", err)
	}
}
