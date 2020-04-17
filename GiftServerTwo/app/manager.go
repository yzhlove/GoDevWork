package app

import (
	"WorkSpace/GoDevWork/GiftServerTwo/manager"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
	"context"
	"errors"
	"fmt"
)

var ServerBusyErr = errors.New("server is busy")

func (p *app) Generate(_ context.Context, req *pb.Manager_GenReq) (*pb.Manager_Nil, error) {

	if p.m.Get() {
		return &pb.Manager_Nil{}, ServerBusyErr
	}

	p.m.Lock()
	defer p.m.Unlock()

	code, err := manager.GenerateCodes(p.h.entity.AutoId, req)
	if err != nil {
		return &pb.Manager_Nil{}, err
	}

	p.h.UpdateEntity(code)
	p.h.Sync(code)

	return &pb.Manager_Nil{}, nil
}

func (p *app) List(_ context.Context, _ *pb.Manager_Nil) (*pb.Manager_ListResp, error) {

	resp := &pb.Manager_ListResp{
		Details: make([]*pb.Manager_CodeInfo, 0, len(p.h.entity.Infos)),
	}
	for _, code := range p.h.entity.Infos {
		resp.Details = append(resp.Details, manager.GeneratePtoCodeInfo(code))
	}
	return resp, nil
}

func (p *app) Export(_ context.Context, req *pb.Manager_ExportReq) (*pb.Manager_ExportResp, error) {

	if p.m.Get() {
		return &pb.Manager_ExportResp{}, ServerBusyErr
	}

	p.m.Lock()
	defer p.m.Unlock()

	code, ok := p.h.entity.Infos[req.Id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("not found to id (%d)", req.Id))
	}

	respCodes, err := manager.GetCodes(req.Id, code.FixCode)
	if err != nil {
		return nil, err
	}

	return &pb.Manager_ExportResp{Details: respCodes}, nil
}
