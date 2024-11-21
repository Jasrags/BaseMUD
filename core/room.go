package core

import (
	"io"

	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/rs/zerolog/log"
	eventemitter "github.com/vansante/go-event-emitter"
)

const (
	RoomPlayerEnter eventemitter.EventType = "Room#event:playerEnter"
	RoomPlayerLeave eventemitter.EventType = "Room#event:playerLeave"
	RoomNpcEnter    eventemitter.EventType = "Room#event:npcEnter"
	RoomNpcLeave    eventemitter.EventType = "Room#event:npcLeave"
	RoomReady       eventemitter.EventType = "Room#event:ready"
	RoomSpawn       eventemitter.EventType = "Room#event:spawn"
	RoomUpdateTick  eventemitter.EventType = "Room#event:updateTick"
)

// type Room interface {
// 	AddItem(item Item)
// 	// AddListener2(behaviorName string)
// 	AddNpc(npc NPC)
// 	AddPlayer(p Player)
// 	CloseDoor(fromRoom *Room)
// 	// Emit(eventName string, args ...interface{})
// 	FindExit(exitName string) //false|object
// 	GetBroadcastTargets() []Character
// 	GetDoor(fromRoom *Room)       // object
// 	GetExits()                   // {Array.<{id: string, direction: string, inferred: boolean, room: Room=}>}
// 	GetExitToRoom(nextRoom Room) // {false|Object}
// 	GetID() string
// 	HasDoor(fromRoom *Room) bool
// 	IsDoorLocked(fromRoom *Room) bool
// 	LockDoor(fromRoom *Room)
// 	OpenDoor(fromRoom *Room)
// 	RemoveItem(item Item)
// 	RemoveNpc(n NPC, removeSpawn bool)
// 	RemovePlayer(p Player)
// 	Render(s ssh.Session)
// 	SpawnItem(state, id string) Item
// 	SpawnNpc(state, id string) NPC
// 	UnlockDoor(fromRoom *Room)

// 	// Event Handlers
// 	// HandleChannelReceiveEvent(arguments ...interface{})
// 	// HandleNpcEnterEvent(arguments ...interface{})
// 	// HandleNpcLeaveEvent(arguments ...interface{})
// 	HandlePlayerEnterEvent(arguments ...interface{})
// 	HandlePlayerLeaveEvent(arguments ...interface{})
// 	// HandleRoomReadyEvent(arguments ...interface{})
// 	// HandleRoomSpawnEvent(arguments ...interface{})
// 	// HandleUpdateTickEvent(arguments ...interface{})

// 	eventemitter.EventEmitter
// 	eventemitter.Observable
// }

type Room struct {
	Area        *Area `yaml:"area"`
	Coordinates []int `yaml:"coordinates"`
	// 	// DefaultItems
	// 	// DefaultNPSs []NPC
	// 	// Exits
	Id string `yaml:"id"`
	// 	// Items []Item
	// 	// NPCs  []NPC
	Players map[string]*Player `yaml:"-"`
	// 	// Script string
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	// // doors []Door
	// spawnedNpcs map[string]NPC

	listeners                 []*eventemitter.Listener `yaml:"-"`
	eventemitter.EventEmitter `yaml:"-"`
	eventemitter.Observable   `yaml:"-"`
}

func (r *Room) Init(em eventemitter.EventEmitter, ob eventemitter.Observable) {
	r.EventEmitter = em
	r.Observable = ob

	r.Players = make(map[string]*Player)

	// Setup listeners
	r.listeners = append(r.listeners,
		r.AddListener(
			eventemitter.EventType(RoomPlayerEnter),
			eventemitter.HandleFunc(r.HandlePlayerEnterEvent)))
	r.listeners = append(r.listeners,
		r.AddListener(
			eventemitter.EventType(RoomPlayerLeave),
			eventemitter.HandleFunc(r.HandlePlayerLeaveEvent)))
}

// HandlePlayerEnterEvent
// player, prevRoom
func (r *Room) HandlePlayerEnterEvent(arguments ...interface{}) {
	log.Debug().
		Str("event", string(RoomPlayerEnter)).
		Str("area_id", r.Area.Name).
		Str("room_id", r.Id).
		Msg("Player entered room:listener")

	if len(arguments) < 2 {
		log.Error().Msg("HandlePlayerEnterEvent: not enough arguments")
		return
	}

	// p := arguments[0].(*Player)
	// prevRoom := arguments[1].(*Room)

	// for _, player := range prevRoom.Players {
	// 	if player.Name != p.Name {
	// 		io.WriteString(player.Session, cfmt.Sprintf("{{%s}} has entered the room.\n", p.Name))
	// 	}
	// }
}

// HandlePlayerLeaveEvent
// player, nextRoom
func (r *Room) HandlePlayerLeaveEvent(arguments ...interface{}) {
	log.Debug().
		Str("event", string(RoomPlayerLeave)).
		Str("area_id", r.Area.Name).
		Str("room_id", r.Id).
		Msg("Player left room:listener")

	if len(arguments) < 2 {
		log.Error().Msg("HandlePlayerLeaveEvent: not enough arguments")
		return
	}

	// p := arguments[0].(*Player)
	// nextRoom := arguments[1].(*Room)

}

// AddItem implements Room.
func (r *Room) AddItem(item *Item) {
	panic("unimplemented")
}

// AddNpc implements Room.
func (r *Room) AddNpc(npc *Npc) {
	panic("unimplemented")
}

// AddPlayer implements Room.
func (r *Room) AddPlayer(p *Player) {
	r.Players[p.Name] = p
}

// CloseDoor implements Room.
func (r *Room) CloseDoor(fromRoom *Room) {
	panic("unimplemented")
}

func (r *Room) Emit(event string, args ...interface{}) {
	log.Debug().Str("source", "player").Str("event", event).Msg("Emitting event")
	r.EventEmitter.EmitEvent(eventemitter.EventType(event), args...)
}

// FindExit implements Room.
func (r *Room) FindExit(exitName string) {
	panic("unimplemented")
}

// GetBroadcastTargets implements Room.
func (r *Room) GetBroadcastTargets() []Character {
	panic("unimplemented")
}

// GetDoor implements Room.
func (r *Room) GetDoor(fromRoom *Room) {
	panic("unimplemented")
}

// GetExitToRoom implements Room.
func (r *Room) GetExitToRoom(nextRoom *Room) {
	panic("unimplemented")
}

// GetExits implements Room.
func (r *Room) GetExits() {
	panic("unimplemented")
}

// HasDoor implements Room.
func (r *Room) HasDoor(fromRoom *Room) bool {
	panic("unimplemented")
}

// IsDoorLocked implements Room.
func (r *Room) IsDoorLocked(fromRoom *Room) bool {
	panic("unimplemented")
}

// LockDoor implements Room.
func (r *Room) LockDoor(fromRoom *Room) {
	panic("unimplemented")
}

// OpenDoor implements Room.
func (r *Room) OpenDoor(fromRoom *Room) {
	panic("unimplemented")
}

// RemoveItem implements Room.
func (r *Room) RemoveItem(item *Item) {
	panic("unimplemented")
}

// RemoveNpc implements Room.
func (r *Room) RemoveNpc(n *Npc, removeSpawn bool) {
	panic("unimplemented")
}

// RemovePlayer implements Room.
func (r *Room) RemovePlayer(p *Player) {
	delete(r.Players, p.Name)
}

func (r *Room) Render(s ssh.Session) {
	io.WriteString(s, cfmt.Sprintf("{{%s}}::yellow|bold\n", r.Title))
	io.WriteString(s, cfmt.Sprintf("{{%s}}::bold\n", r.Description))
}

// SpawnItem implements Room.
func (r *Room) SpawnItem(state string, id string) Item {
	panic("unimplemented")
}

// SpawnNpc implements Room.
func (r *Room) SpawnNpc(state string, id string) *Npc {
	r.EmitEvent("Npc#event:spawn", state, id)

	return nil
}

// UnlockDoor implements Room.
func (r *Room) UnlockDoor(fromRoom *Room) {
	panic("unimplemented")
}

// var (
// 	TheVoid = Area{
// 		Name: "The Void",
// 		Rooms: map[int]*Room{
// 			0: {
// 				Players: make(map[string]*Player),
// 				ID:      0,
// 				// Coordinates: &Coordinate{X: 0, Y: 0, Z: 0},
// 				Title:       "The Void",
// 				Description: "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
// 			},
// 		},
// 	}
// )
