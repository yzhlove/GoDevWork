package app

import (
	"WorkSpace/GoDevWork/GiftServerTwo/manager"
	"WorkSpace/GoDevWork/GiftServerTwo/misc/helper"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
	"context"
	"errors"
	"strconv"
)

func (p *app) CodeVerify(_ context.Context, req *pb.VerifyReq) (*pb.VerifyResp, error) {

	resp := &pb.VerifyResp{Status: 1}
	var (
		id uint32
		ok bool
	)

	if id, ok = helper.ToDecodeNumber(helper.ToDecodeStr(req.Code)); !ok {
		if id, ok = p.h.entity.Fixed[req.Code]; !ok {
			return resp, errors.New("not found code:" + req.Code)
		}
	}

	code, ok := p.h.entity.Infos[id]
	if !ok {
		return resp, errors.New("not found id :" + strconv.Itoa(int(id)))
	}

	var lock *MutexType
	if lock = p.ms.Get(req.Code); lock.Get() {
		return resp, ServerBusyErr
	}

	//验证过程加锁
	lock.Lock()
	defer lock.Unlock()

	var err error
	if ok, err = manager.CodeVerify(code, req.UserId, req.Zone, req.Code); err != nil {
		return resp, err
	}
	if ok {
		resp.Status = 0
	}
	return resp, nil
}

func (p *app) Sync(req *pb.SyncReq, stream pb.GiftService_SyncServer) error {

	p.h.setConn(req.Zone, stream)
	p.h.reqChan <- req.Zone
	<-stream.Context().Done()
	return nil
}
