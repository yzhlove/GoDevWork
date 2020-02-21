package app

import (
	"WorkSpace/GoDevWork/GiftServer/manager"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	"WorkSpace/GoDevWork/GiftServer/pubsub"
	"context"
)

func (p *app) Generate(_ context.Context, req *pb.Manager_GenReq) (*pb.Manager_Nil, error) {

	message, err := manager.GenerateCodeInfo(req)
	if err != nil {
		return nil, err
	}

	pubsub.Pub(req.ZoneIds, message)

	return &pb.Manager_Nil{}, nil
}

func (p *app) List(_ context.Context, _ *pb.Manager_Nil) (*pb.Manager_ListResp, error) {

	return &pb.Manager_ListResp{}, nil
}

func (p *app) Export(_ context.Context, req *pb.Manager_ExportReq) (*pb.Manager_ExportResp, error) {
	return &pb.Manager_ExportResp{}, nil
}
