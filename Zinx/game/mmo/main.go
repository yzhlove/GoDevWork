package main

import (
	"log"
	"zinx-game-example/mmo/api"
	"zinx-game-example/mmo/core"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	server := znet.NewTcpServer()
	server.ConnStartEvent(OnConnStart)
	server.ConnStopEvent(OnConnStop)
	server.Register(2, &api.WorldChatApi{})
	server.Register(3, &api.MoveApi{})
	server.Run()

}

func OnConnStart(conn ziface.ConnImp) {
	if conn != nil {
		player := core.NewPlayer(conn)
		player.SyncPid()
		player.BroadCastStartPosition()
		core.WorldMgr.AddPlayer(player)
		conn.SetAttr("pid", player.PID)
		player.SyncRangePlayers()
		log.Println("==> player pid:", player.PID, " active.")
	}
}

func OnConnStop(conn ziface.ConnImp) {
	pid, ok := conn.GetAttr("pid")
	if !ok {
		log.Println("get attr pid err")
		return
	}
	if player := core.WorldMgr.GetPlayer(pid.(int32)); player != nil {
		player.LostConn()
	}
	log.Println("player leave ==> ", pid)
}
