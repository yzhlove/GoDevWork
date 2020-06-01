package core

import (
	"fmt"
	"sync"
)

// 一个地图中的格子
type Grid struct {
	GID                    int              //格子ID
	MinX, MinY, MaxX, MaxY int              //格子的上下左右坐标
	playerIds              map[int]struct{} //格子内的玩家或者物体成员的ID
	mutex                  sync.RWMutex
}

func NewGrid(gid, minX, minY, maxX, maxY int) *Grid {
	return &Grid{
		GID:  gid,
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (g *Grid) Add(playID int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.playerIds[playID] = struct{}{}
}

func (g *Grid) Del(playID int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	delete(g.playerIds, playID)
}

func (g *Grid) GetPlayIds() []int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	ids := make([]int, 0, len(g.playerIds))
	for id := range g.playerIds {
		ids = append(ids, id)
	}
	return ids
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid : %d ,minX %d maxX %d minY %d maxY %d playerIds:%v ",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIds)
}
