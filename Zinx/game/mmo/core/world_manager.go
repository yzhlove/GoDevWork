package core

import "sync"

//世界管理模块

type WorldManager struct {
	aoiMgr  *AOIMgr
	players map[int32]*Player
	sync.RWMutex
}

var WorldMgr *WorldManager

func init() {
	WorldMgr = &WorldManager{
		aoiMgr:  NewAOIMgr(AOI_MIN_X, AOI_MIN_Y, AOI_MAX_X, AOI_MAX_Y, AOI_CNT_X, AOI_CNT_Y),
		players: make(map[int32]*Player),
	}
}

func (m *WorldManager) AddPlayer(player *Player) {
	m.Lock()
	m.players[player.PID] = player
	m.Unlock()
	m.aoiMgr.AddFormGridByPos(int(player.PID), player.X, player.Y)
}

func (m *WorldManager) DelPlayer(pid int32) {
	m.Lock()
	defer m.Unlock()
	delete(m.players, pid)
}

func (m *WorldManager) GetPlayer(pid int32) *Player {
	m.RLock()
	defer m.RUnlock()
	return m.players[pid]
}

func (m *WorldManager) GetPlayers() []*Player {
	m.RLock()
	defer m.RUnlock()
	players := make([]*Player, 0, len(m.players))
	for _, player := range m.players {
		players = append(players, player)
	}
	return players
}
