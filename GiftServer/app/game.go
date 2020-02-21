package app

import (
	"WorkSpace/GoDevWork/GiftServer/entity"
	"WorkSpace/GoDevWork/GiftServer/manager"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	"WorkSpace/GoDevWork/GiftServer/pubsub"
	"context"
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

	channel, ok := pubsub.Sub(req.Zone)
	if !ok {
		go func() {
			for _, code := range entity.Entity.CodesMap {
				channel <- manager.GeneratePtoCodeMessage(code)
			}
		}()
	}

	defer pubsub.CloseChan(req.Zone)

	for msg := range channel {
		if code, ok := msg.(*pb.Manager_CodeInfo); ok {
			if err := stream.Send(code); err != nil {
				return err
			}
		}
	}

	return nil
}
