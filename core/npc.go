package core

import (
	eventemitter "github.com/vansante/go-event-emitter"
)

type Npc struct {
	Id   int
	Area *Area
	// Behaviors map[string]Behavior

	Character

	listeners                 []*eventemitter.Listener `yaml:"-"`
	eventemitter.EventEmitter `yaml:"-"`
	eventemitter.Observable   `yaml:"-"`
}

// func NewNPC(em eventemitter.EventEmitter, ob eventemitter.Observable) *Npc {
// 	n := &Npc{}

// 	n.EventEmitter = em
// 	n.Observable = ob

// 	return n
// }

func (n *Npc) Init(em eventemitter.EventEmitter, ob eventemitter.Observable) {
	n.EventEmitter = em
	n.Observable = ob

	// Setup listeners
}

// Hydrate implements NPC.
func (n *Npc) Hydrate(state string) {
	panic("unimplemented")
}

// MoveTo implements NPC.
// Emits
// Room#event:npcLeave
// Room#event:npcEnter
// Npc#event:enterRoom
func (n *Npc) MoveTo(nextRoom Room, onMoved func()) {
	panic("unimplemented")
}

func (n *Npc) IsNpc() bool {
	return true
}
