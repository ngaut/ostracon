package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/nathanieltornow/ostracon/recordshard/recordshardpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

var (
	parentIpAddr = flag.String("parentIpAddr", "", "")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*parentIpAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln("Failed making connection to shard")
	}
	defer conn.Close()

	shardClient := pb.NewRecordShardClient(conn)
	time.Sleep(3 * time.Second)

	stream, err := shardClient.Subscribe(context.Background(), &pb.Empty{})
	if err != nil {
		logrus.Fatalln(err)
	}
	for {
		in, err := stream.Recv()
		if err != nil {
			logrus.Fatalln(err)
		}
		fmt.Println(in)
	}
}
