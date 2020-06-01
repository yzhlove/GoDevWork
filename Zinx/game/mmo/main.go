package main

import (
	"log"
	"zinx-game-example/mmo/core"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	core.NewPlayer(nil)

	server := znet.NewTcpServer()
	server.ConnStartEvent(OnConnStart)
	server.Run()

}

func OnConnStart(conn ziface.ConnImp) {
	player := core.NewPlayer(conn)
	player.SyncPid()
	player.BroadCastStartPostion()
	log.Println("==> player pid:", player.PID, " active.")
}
