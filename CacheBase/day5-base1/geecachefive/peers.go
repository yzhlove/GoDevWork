package geecachefive

type PeerPick interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(group, key string) ([]byte, error)
}
