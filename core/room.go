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
	// area Area
	coordinates []int
	// 	// DefaultItems
	// 	// DefaultNPSs []NPC
	// 	// Exits
	id string
	// 	// Items []Item
	// 	// NPCs  []NPC
	players map[string]*Player
	// 	// Script string
	title       string
	description string
	// // doors []Door
	// spawnedNpcs map[string]NPC

	// listeners []*eventemitter.Listener

	eventemitter.EventEmitter
	eventemitter.Observable
}

func NewRoom(em eventemitter.EventEmitter, ob eventemitter.Observable) *Room {
	// ob.AddListener(
	// 	eventemitter.EventType("Room#event:playerEnter"),
	// 	eventemitter.HandleFunc(func(arguments ...interface{}) {
	// 		log.Info().Msg("Player left room")
	// 	}))
	// ob.AddListener(
	// 	eventemitter.EventType("Room#event:playerLeave"),
	// 	eventemitter.HandleFunc(func(arguments ...interface{}) {
	// 		log.Info().Msg("Player entered room")
	// 	}))
	// ob.AddListener(
	// 	eventemitter.EventType("Player#event:enterRoom"),
	// 	eventemitter.HandleFunc(func(arguments ...interface{}) {
	// 		log.Info().Msg("Player left room")
	// 	}))

	return &Room{
		players: make(map[string]*Player),

		EventEmitter: em,
		Observable:   ob,
	}
}

func (r *Room) HandlePlayerEnterEvent(arguments ...interface{}) {
	log.Info().Msg("Player entered room")
}

func (r *Room) HandlePlayerLeaveEvent(arguments ...interface{}) {
	log.Info().Msg("Player left room")
}

// AddItem implements Room.
func (r *Room) AddItem(item *Item) {
	panic("unimplemented")
}

// func (r *Room) AddListener2(behaviorName string) {
// 	log.Debug().Str("source", "room").Str("event", behaviorName).Msg("Adding listener")

// 	// r.AddListener(
// 	// 	eventemitter.EventType("Room#event:playerEnter"),
// 	// 	eventemitter.HandleFunc(func(arguments ...interface{}) {
// 	// 		log.Info().Msg("Player entered room")
// 	// 	}))
// 	// r.listeners = append(r.listeners, listener)
// }

// AddNpc implements Room.
func (r *Room) AddNpc(npc *Npc) {
	panic("unimplemented")
}

// AddPlayer implements Room.
func (r *Room) AddPlayer(p *Player) {
	r.players[p.Name] = p
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

func (r *Room) GetID() string {
	return r.id
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
	delete(r.players, p.Name)
}

func (r *Room) Render(s ssh.Session) {
	io.WriteString(s, cfmt.Sprintf("{{%s}}::yellow|bold\n", r.title))
	io.WriteString(s, cfmt.Sprintf("{{%s}}::bold\n", r.description))
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
