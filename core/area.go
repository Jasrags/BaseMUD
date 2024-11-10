package core

import "github.com/rs/zerolog/log"

// type Coordinate struct {
// X, Y, Z int
// }

type Area struct {
	Name  string
	Title string
	// Script string
	// Map   map[int]Coordinate
	Rooms map[int]*Room
	// NPCs []NPC
	// Info Object
	// LastRespawnTick int
}

func NewArea() *Area {
	return &Area{
		Rooms: make(map[int]*Room),
	}
}

func (a *Area) AddRoom(r *Room) {
	a.Rooms[r.ID] = r

	// if a.Map == nil {
	// a.AddRoomToMap(r)
	// }

	a.Emit("area:roomAdded", r)

}

// func (a *Area) AddRoomToMap(r *Room) {
// a.Map[r.Coordinates.Z] = *r.Coordinates
// }

func (a *Area) Emit(event string, args ...interface{}) {
	log.Debug().
		Str("source", "area").
		Str("event", event).
		Msg("Emitting event")
}
