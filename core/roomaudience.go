package core

type RoomAudience struct {
}

func NewRoomAudience() *RoomAudience {
	return &RoomAudience{}
}

func (ra *RoomAudience) AlterMessage(msg string) string {
	return msg
}

func (ra *RoomAudience) Configure(config map[string]interface{}) {
}

func (ra *RoomAudience) GetBroadcastTargets() []*Player {
	return []*Player{}
}
