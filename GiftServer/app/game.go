package app

import (
	"WorkSpace/GoDevWork/GiftServer/manager"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	"WorkSpace/GoDevWork/GiftServer/pubsub"
	"context"
	"errors"
	"io"
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

	stream.Context()

	for msg := range pubsub.Sub(req.Zone) {
		if code, ok := msg.(*pb.Manager_CodeInfo); ok {
			if err := stream.Send(code); err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return err
			}
		}
	}
	return nil
}
