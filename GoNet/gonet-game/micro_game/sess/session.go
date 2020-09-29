package sess

const (
	SESS_KICKED_OUT = 0x1 //踢掉
)

// Session 会话是一个单独玩家的上下文，在连入后到退出前的整个生命周期内存在，可以根据
// 业务自行扩展
type Session struct {
	Flag      int32 //会话标记
	UserId    int32
	LastReqId int16 //上一次会话使用的code
}
