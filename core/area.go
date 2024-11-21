package core

import eventemitter "github.com/vansante/go-event-emitter"

// type Area interface {
// 	AddNpc(n NPC)
// 	AddRoom(r Room)
// 	AddRoomToMap(r Room)
// 	GetBroadcastTargets() // npcs, players, rooms, and the area itself
// 	GetRoomAtCoordinates(x, y, z int) Room
// 	GetRoomByID(id string) Room
// 	RemoveNpc(n NPC)
// 	RemoveRoom(r Room)
// 	Update(state string)
// }

type Area struct {
	Name  string `yaml:"name"`
	Title string `yaml:"title"`
	// Script string
	// Map   map[int]Coordinate
	Rooms map[string]*Room `yaml:"-"`
	Npcs  []*Npc           `yaml:"-"`
	// Info Object
	// LastRespawnTick int

	// areaPath
	// floors
}

const (
	AreaRoomAdded   eventemitter.EventType = "Area#event:roomAdded"
	AreaRoomRemoved eventemitter.EventType = "Area#event:roomRemoved"
)

func NewArea() *Area {
	return &Area{
		Rooms: make(map[string]*Room),
	}
}

// AddNpc implements Area.
func (a *Area) AddNpc(n *Npc) {
	a.Npcs = append(a.Npcs, n)
}

// AddRoom implements Area.
// Emits
// Area#event:roomAdded
func (a *Area) AddRoom(r *Room) {
	a.Rooms[r.Id] = r
}

// AddRoomToMap implements Area.
func (a *Area) AddRoomToMap(r *Room) {
	panic("unimplemented")
}

// GetBroadcastTargets implements Area.
func (a *Area) GetBroadcastTargets() {
	panic("unimplemented")
}

// GetRoomAtCoordinates implements Area.
func (a *Area) GetRoomAtCoordinates(x int, y int, z int) *Room {
	panic("unimplemented")
}

// GetRoomByID implements Area.
func (a *Area) GetRoomByID(id string) *Room {
	if room, ok := a.Rooms[id]; ok {
		return room
	}
	return nil
}

// RemoveNpc implements Area.
func (a *Area) RemoveNpc(n *Npc) {
	// for i, npc := range a.Npcs {
	// 	if npc.Id == n.Id {
	// 		a.Npcs = append(a.Npcs[:i], a.Npcs[i+1:]...)
	// 		return
	// 	}
	// }
}

// RemoveRoom implements Area.
// Emits Area#event:roomRemoved
func (a *Area) RemoveRoom(r *Room) {
	delete(a.Rooms, r.Id)
}

// Update implements Area.
// Emits
// Room#event:updateTick
// Npc#event:updateTick
func (a *Area) Update(state string) {
	panic("unimplemented")
}
