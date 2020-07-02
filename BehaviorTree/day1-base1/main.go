package main

import (
	"behavior_day1_base1/bt"
	"fmt"
	"sync"
	"time"
)

func main() {

	if err := bt.Source("tree.txt"); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go Run(&wg)
	wg.Wait()
}

func Run(wg *sync.WaitGroup) {
	npc := NewNpc()
	ai, err := bt.NewTask("tree", npc)
	if err != nil {
		panic(err)
	}
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			npc.hp -= 5
			ai.Update(33)
			if npc.hp <= 0 {
				wg.Done()
				return
			}
		}
	}
}

type Npc struct {
	hp    int
	mp    int
	isEnd bool
}

func NewNpc() *Npc {
	return &Npc{
		hp:    100,
		mp:    100,
		isEnd: false,
	}
}

func (n *Npc) GetCurHp() int {
	return n.hp
}

func (n *Npc) GetCurMp() int {
	return n.mp
}

func (n *Npc) CastSkill(sid, mp int) {
	n.mp -= mp
	n.isEnd = true
	fmt.Println("npc cast skill id:", sid, " used mp:", mp)
}

func (n *Npc) IsSkillEnd() bool {
	return n.isEnd
}

func (n *Npc) Run() {
	fmt.Println("npc running...")
}

func (n *Npc) Eat(id int) {
	fmt.Println("eat item:", id)
}
