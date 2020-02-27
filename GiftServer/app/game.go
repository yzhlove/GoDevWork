package app

import (
	"WorkSpace/GoDevWork/GiftServer/manager"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	"WorkSpace/GoDevWork/GiftServer/pubsub"
	"context"
	log "github.com/sirupsen/logrus"
)

func (p *app) CodeVerify(_ context.Context, req *pb.VerifyReq) (*pb.VerifyResp, error) {
	resp := &pb.VerifyResp{Status: 1}
	if ok, err := manager.CodeVerify(req.UserId, req.Zone, req.Code); err != nil {
		return nil, err
	} else if ok {
		resp.Status = 0
	}
	return resp, nil
}

func (p *app) Sync(req *pb.SyncReq, stream pb.GiftService_SyncServer) error {
	log.Info("new stream to ", req.Zone)
	c := pubsub.Sub(req.Zone)
	ctx := stream.Context()
	for {
		select {
		case msg, ok := <-c.MsgChan:
			if ok {
				if code, ok := msg.(*pb.Manager_CodeInfo); ok {
					if err := stream.Send(code); err != nil {
						return err
					}
				}
			} else {
				return nil
			}
		case <-ctx.Done():
			pubsub.Close(c)
			log.Error("close stream ...")
			return nil
		}
	}
}
