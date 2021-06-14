package seqshard

import (
	pb "github.com/nathanieltornow/ostracon/seqshard/seqshardpb"
	"io"
)

func (s *SeqShard) GetOrder(stream pb.Shard_GetOrderServer) error {
	if s.isRoot {
		s.snMu.Lock()
		s.sn = 0
		s.snMu.Unlock()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			s.snMu.Lock()
			res := pb.OrderResponse{StartGsn: s.sn, StartLsn: req.StartLsn, NumOfRecords: req.NumOfRecords}
			s.sn += req.NumOfRecords
			s.snMu.Unlock()

			if err := stream.Send(&res); err != nil {
				return err
			}
		}
	}

	s.orderRespCsMu.Lock()
	s.orderRespCs[stream] = make(chan *orderResponse, 4096)
	s.orderRespCsMu.Unlock()

	go s.sendOrderResponses(stream)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.snMu.Lock()
		oR := orderRequest{stream: stream, numOfRecords: req.NumOfRecords, startLsn: req.StartLsn}
		s.waitingOrderReqs[s.sn] = &oR
		s.sn += req.NumOfRecords
		s.snMu.Unlock()
		s.orderReqsC <- &oR
	}

}

// sendOrderResponses handles a specific order-stream to send back all finished order-requests in its channel
func (s *SeqShard) sendOrderResponses(stream pb.Shard_GetOrderServer) {
	for finishedOR := range s.orderRespCs[stream] {
		res := pb.OrderResponse{StartLsn: finishedOR.startLsn, StartGsn: finishedOR.startGsn, NumOfRecords: finishedOR.numOfRecords}
		if err := stream.Send(&res); err != nil {
			return
		}
	}
}
