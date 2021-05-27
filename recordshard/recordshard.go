package recordshard

import (
	"context"
	rpb "github.com/nathanieltornow/ostracon/recordshard/recordshardpb"
	"github.com/nathanieltornow/ostracon/recordshard/storage"
	spb "github.com/nathanieltornow/ostracon/shard/shardpb"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

type record struct {
	lsn    chan int64
	record string
}

type RecordShard struct {
	rpb.UnimplementedRecordShardServer
	disk             *storage.Storage
	parentConn       *grpc.ClientConn
	parentClient     *spb.ShardClient
	batchingInterval time.Duration
	curLsn           int64
	curLsnMu         sync.Mutex
	waitCMap         map[int64]chan int64
	waitCMapMu       sync.Mutex
	writeC           chan *record
}

func NewRecordShard(diskPath string, batchingInterval time.Duration) (*RecordShard, error) {
	disk, err := storage.NewStorage(diskPath, 0, 1, 1000)
	if err != nil {
		return nil, err
	}
	s := RecordShard{}
	s.disk = disk
	s.batchingInterval = batchingInterval
	s.writeC = make(chan *record, 4096)
	s.waitCMap = make(map[int64]chan int64)
	s.curLsn = -1
	return &s, nil
}

func (rs *RecordShard) Start(ipAddr string, parentIpAddr string) error {
	lis, err := net.Listen("tcp", ipAddr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	rpb.RegisterRecordShardServer(grpcServer, rs)

	if parentIpAddr != "" {
		time.Sleep(3 * time.Second)
		err = rs.ConnectToParent(parentIpAddr)
		if err != nil {
			return err
		}
	}
	go rs.writeAppends()
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (rs *RecordShard) ConnectToParent(parentIpAddr string) error {
	conn, err := grpc.Dial(parentIpAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := spb.NewShardClient(conn)
	rs.parentClient = &client
	rs.parentConn = conn

	stream, err := client.GetOrder(context.Background())
	if err != nil {
		return err
	}
	go rs.sendOrderRequests(stream)
	go rs.receiveOrderResponses(stream)
	return nil
}