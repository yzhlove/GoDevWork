package geecacheseven

import "geecacheseven/pb"

type Req = *pb.Cache_Req
type Resp = *pb.Cache_Resp

type PeerPick interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(in *Req, out *Resp) error
}
