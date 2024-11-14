package core

type RoomFactory struct {
}

func NewRoomFactory() *RoomFactory {
	return &RoomFactory{}
}

// Create a new instance of a room
func (rf *RoomFactory) Create(area Area, id int) *Room {
	return NewRoom()
}
