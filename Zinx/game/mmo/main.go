package main

import (
	"log"
	"zinx-game-example/mmo/api"
	"zinx-game-example/mmo/core"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	core.NewPlayer(nil)

	server := znet.NewTcpServer()
	server.ConnStartEvent(OnConnStart)
	server.Register(2, &api.WorldChatApi{})
	server.Run()

}

func OnConnStart(conn ziface.ConnImp) {
	player := core.NewPlayer(conn)
	player.SyncPid()
	player.BroadCastStartPosition()
	core.WorldMgr.AddPlayer(player)
	conn.SetAttr("pid", player.PID)
	log.Println("==> player pid:", player.PID, " active.")
}
