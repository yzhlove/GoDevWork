//Code generated script: gen_proto.sh
package handler

import "micro_game/misc/packet"

// AutoId 公共结构，用于只传id或一个数字的结构
type AutoId struct {
	Id int32
}

func (p AutoId) Pack(w *packet.Packet) {
	w.WriteS32(p.Id)
}

func PacketAutoId(reader *packet.Packet) (tbl AutoId, err error) {
	if tbl.Id, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	return
}

// ErrorInfo 一般性回复payload,0代表成功
type ErrorInfo struct {
	Code int32
	Msg  string
}

func (p ErrorInfo) Pack(w *packet.Packet) {
	w.WriteS32(p.Code)
	w.WriteString(p.Msg)
}

func PacketErrorInfo(reader *packet.Packet) (tbl ErrorInfo, err error) {
	if tbl.Code, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	if tbl.Msg, err = reader.ReadString(); err != nil {
		panic(err)
	}
	return
}

// UserLoginInfo 用户登陆发包 1代表使用uuid登陆 2代表使用证书登陆
type UserLoginInfo struct {
	LoginWay          int32
	OpenUid           string
	ClientCertificate string
	ClientVersion     int32
	UserLang          string
	AppId             string
	OsVersion         string
	DeviceName        string
	DeviceId          string
	DeviceIdType      int32
	LoginIp           string
}

func (p UserLoginInfo) Pack(w *packet.Packet) {
	w.WriteS32(p.LoginWay)
	w.WriteString(p.OpenUid)
	w.WriteString(p.ClientCertificate)
	w.WriteS32(p.ClientVersion)
	w.WriteString(p.UserLang)
	w.WriteString(p.AppId)
	w.WriteString(p.OsVersion)
	w.WriteString(p.DeviceName)
	w.WriteString(p.DeviceId)
	w.WriteS32(p.DeviceIdType)
	w.WriteString(p.LoginIp)
}

func PacketUserLoginInfo(reader *packet.Packet) (tbl UserLoginInfo, err error) {
	if tbl.LoginWay, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	if tbl.OpenUid, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.ClientCertificate, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.ClientVersion, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	if tbl.UserLang, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.AppId, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.OsVersion, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.DeviceName, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.DeviceId, err = reader.ReadString(); err != nil {
		panic(err)
	}
	if tbl.DeviceIdType, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	if tbl.LoginIp, err = reader.ReadString(); err != nil {
		panic(err)
	}
	return
}

// SeedInfo 通信加密种子
type SeedInfo struct {
	ClientSendSeed    int32
	ClientReceiveSeed int32
}

func (p SeedInfo) Pack(w *packet.Packet) {
	w.WriteS32(p.ClientSendSeed)
	w.WriteS32(p.ClientReceiveSeed)
}

func PacketSeedInfo(reader *packet.Packet) (tbl SeedInfo, err error) {
	if tbl.ClientSendSeed, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	if tbl.ClientReceiveSeed, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	return
}

// UserSnapshot 用户信息包
type UserSnapshot struct {
	Uid int32
}

func (p UserSnapshot) Pack(w *packet.Packet) {
	w.WriteS32(p.Uid)
}

func PacketUserSnapshot(reader *packet.Packet) (tbl UserSnapshot, err error) {
	if tbl.Uid, err = reader.ReadS32(); err != nil {
		panic(err)
	}
	return
}
