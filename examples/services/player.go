package services

//玩家
type Player struct {
	Pid uint32
}

//创建玩家对象
func NewPlayer(pid uint32) *Player {
	return &Player{
		Pid: pid,
	}
}

//获取玩家id
func (p *Player) GetPlayerId() uint32 {
	return p.Pid
}
