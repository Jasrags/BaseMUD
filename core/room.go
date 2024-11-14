package core

import "github.com/rs/zerolog/log"

type Room struct {
	Area Area
	// Coordinates *Coordinate
	// 	// DefaultItems
	// 	// DefaultNPSs []NPC
	// 	// Exits
	ID int
	// 	// Items []Item
	// 	// NPCs  []NPC
	// Players []*Player
	Players map[string]*Player
	// 	// Script string
	Title       string
	Description string
	// // doors []Door
}

type R interface {
	AddPlayer(p *Player)
	Emit(event string, args ...interface{})
	GetBroadcastTargets() (*Room, []*Player, []*NPC)
	RemovePlayer(p *Player)
}

func NewRoom() *Room {
	return &Room{
		Players: make(map[string]*Player),
	}
}

func (r *Room) AddPlayer(p *Player) {
	r.Players[p.ID] = p
}

func (r *Room) RemovePlayer(p *Player) {
	delete(r.Players, p.ID)
}

func (r *Room) Emit(event string, args ...interface{}) {
	log.Debug().
		Str("source", "room").
		Str("event", event).
		Msg("Emitting event")
}

func (r *Room) GetBroadcastTargets() (*Room, []*Player, []*NPC) {
	return r, []*Player{}, []*NPC{}

}

type EventRoomPlayerEnter struct {
	Player *Player
	Room   *Room
}

type EventRoomPlayerLeave struct {
	Player   *Player
	NextRoom *Room
}

var (
	TheVoid = Area{
		Name: "The Void",
		Rooms: map[int]*Room{
			0: {
				Players: make(map[string]*Player),
				ID:      0,
				// Coordinates: &Coordinate{X: 0, Y: 0, Z: 0},
				Title:       "The Void",
				Description: "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
			},
		},
	}
)
