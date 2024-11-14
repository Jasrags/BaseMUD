package core

type RoomManager struct {
	Rooms map[int]*Room
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		Rooms: make(map[int]*Room),
	}
}

func (rm *RoomManager) AddRoom(r *Room) {
	rm.Rooms[r.ID] = r
}

func (rm *RoomManager) GetRoom(id int) *Room {
	return rm.Rooms[id]
}

func (rm *RoomManager) RemoveRoom(id int) {
	delete(rm.Rooms, id)
}
