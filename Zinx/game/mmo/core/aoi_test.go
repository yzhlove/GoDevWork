package core

import "testing"

func Test_NewAOIMgr(t *testing.T) {
	aoiMgr := NewAOIMgr(100, 200, 300, 450, 4, 5)
	t.Log(aoiMgr)
}

func Test_RandGrid(t *testing.T) {
	aoiMgr := NewAOIMgr(0, 0, 250, 250, 5, 5)
	for k := range aoiMgr.grids {
		grids := aoiMgr.GetRangeGrids(k)
		t.Log("gid:", k, " rand grid length:", len(grids))
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}
		t.Log("grid:", k, " rand id:", gIDs)
		t.Log()
	}
}
