package geecacheseven

import "geecacheseven/pb"

type PeerPick interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(in *pb.Cache_Req, out *pb.Cache_Resp) error
}
