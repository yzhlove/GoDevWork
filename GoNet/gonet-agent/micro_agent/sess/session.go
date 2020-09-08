package sess

import (
	"crypto/rc4"
	"micro_agent/proto"
	"net"
	"time"
)

const (
	SESS_KEYEXCG    = 0x1 //是否已经交换完毕key
	SESS_ENCRYPT    = 0x2 //是否可以开始加密
	SESS_KICKED_OUT = 0x4 //踢掉
	SESS_AUTHORIZED = 0x8 //已授权访问
)

type Session struct {
	IP              net.IP
	MQ              chan proto.Game_Frame          //返回给客户端的异步消息
	Encoder         *rc4.Cipher                    //加密器
	Decoder         *rc4.Cipher                    //解密器
	UserID          uint64                         //玩家ID
	GSID            string                         //游戏服ID
	Stream          proto.GameService_StreamClient //grpc流
	Die             chan struct{}                  //会话关闭信号
	Flag            int32                          //会话标记
	ConnectTime     time.Time                      //TCP链接建立时间
	PacketTime      time.Time                      //当前包到达时间
	LastPacketTime  time.Time                      //前一个包的到达时间
	PacketCount     uint32                         //对收到的包进行统计，避免恶意发包
	PacketCount1Min int                            //每分钟的包统计，用于RPM判断
}

