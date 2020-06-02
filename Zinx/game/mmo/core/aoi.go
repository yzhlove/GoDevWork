package core

import (
	"fmt"
	"log"
)

const (
	AOI_MIN_X = 85
	AOI_MAX_X = 410
	AOI_MIN_Y = 75
	AOI_MAX_Y = 400
	AOI_CNT_X = 10
	AOI_CNT_Y = 20
)

type AOIMgr struct {
	MinX, MinY, MaxX, MaxY int           //区域的上下左右坐标
	CntX, CntY             int           //x,y轴方向格子的数量
	grids                  map[int]*Grid //当前区域中的格子的数量
}

func NewAOIMgr(minX, minY, maxX, maxY, cntX, cntY int) *AOIMgr {
	aoiMgr := &AOIMgr{
		MinX:  minX,
		MinY:  minY,
		MaxX:  maxX,
		MaxY:  maxY,
		CntX:  cntX,
		CntY:  cntY,
		grids: make(map[int]*Grid),
	}
	//aoi初始化
	for y := 0; y < cntY; y++ {
		for x := 0; x < cntX; x++ {
			//格子ID
			gid := y*cntX + x
			//初始化一个格子
			aoiMgr.grids[gid] = NewGrid(
				gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridHeight(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridHeight(),
			)
		}
	}
	return aoiMgr
}

func (m *AOIMgr) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntX
}

func (m *AOIMgr) gridHeight() int {
	return (m.MaxY - m.MinY) / m.CntY
}

// 根据gid算出周边的格子
func (m *AOIMgr) GetRangeGrids(gID int) (grids []*Grid) {

	//判断当前gid是否存在
	if _, ok := m.grids[gID]; !ok {
		return
	}
	//将当前gid加入到当前格子
	grids = append(grids, m.grids[gID])
	//判断当前gID左边是否有格子
	idx := gID % m.CntX
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	//判断当前gID右边是否有格子
	if idx < m.CntX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	//将x轴的格子全部取出，进行遍历，再分别得到每个格子的上下是否有格子
	idsToX := make([]int, 0, len(grids))
	for _, grid := range grids {
		idsToX = append(idsToX, grid.GID)
	}
	//遍历X轴的格子
	for _, id := range idsToX {
		//计算该格子处在第几列
		idy := id / m.CntX
		//判断idy上边是否有格子
		if idy > 0 {
			grids = append(grids, m.grids[id-m.CntX])
		}
		//判断当前的idy下边是否有格子
		if idy < m.CntY-1 {
			grids = append(grids, m.grids[id+m.CntX])
		}
	}
	return
}

//更具横纵坐标算出对应格子的ID
func (m *AOIMgr) GetGridID(x, y float32) int {
	gx := (int(x) - m.MinX) / m.gridWidth()
	gy := (int(y) - m.MinY) / m.gridHeight()
	return gy*m.CntX + gx
}

//根据玩家的坐标求出周围所有玩家的
func (m *AOIMgr) GetPlayerIDS(x, y float32) (playerIds []int) {
	//获取gID
	gID := m.GetGridID(x, y)
	//获取周围格子
	grids := m.GetRangeGrids(gID)
	for _, grid := range grids {
		playerIds = append(playerIds, grid.GetPlayIds()...)
		log.Printf("==> grid :%d pids:%v \n", grid.GID, grid.GetPlayIds())
	}
	return
}

func (m *AOIMgr) GetPlayers(gid int) []int {
	if grid, ok := m.grids[gid]; ok {
		return grid.GetPlayIds()
	}
	return nil
}

func (m *AOIMgr) DelPlayer(gid, pid int) {
	if grid, ok := m.grids[gid]; ok {
		grid.Del(pid)
	}
}

func (m *AOIMgr) AddPlayer(gid, pid int) {
	if grid, ok := m.grids[gid]; ok {
		grid.Add(pid)
	}
}

func (m *AOIMgr) AddFormGridByPos(pid int, x, y float32) {
	if grid, ok := m.grids[m.GetGridID(x, y)]; ok {
		grid.Add(pid)
	}
}

func (m *AOIMgr) DelFormGridByPos(pid int, x, y float32) {
	if grid, ok := m.grids[m.GetGridID(x, y)]; ok {
		grid.Del(pid)
	}
}

func (m *AOIMgr) String() string {
	s := fmt.Sprintf("AOIMgr:\n\tminX:%d maxX:%d minY:%d maxY:%d cntX:%d cntY:%d \n",
		m.MinX, m.MaxX, m.MinY, m.MaxY, m.CntX, m.CntY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}
