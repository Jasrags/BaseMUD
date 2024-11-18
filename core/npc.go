package core

import (
	eventemitter "github.com/vansante/go-event-emitter"
)

type Npc struct {
	Id   int
	Area *Area
	// Behaviors map[string]Behavior

	Character
}

func NewNPC(em eventemitter.EventEmitter, ob eventemitter.Observable) *Npc {
	n := &Npc{}

	n.EventEmitter = em
	n.Observable = ob

	return n
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
