package core

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/gliderlabs/ssh"
	"github.com/rs/zerolog/log"
	eventemitter "github.com/vansante/go-event-emitter"
)

type Player struct {
	Character
	eventemitter.EventEmitter
	eventemitter.Observable
	// EventEmitter eventemitter.EventEmitter
	Account *Account
	// Experience int
	// Password string
	// Prompt   string
	Session ssh.Session
	// QuestTracker *QuestTracker
	// ExtraPrompts
	// QuestData

	PubSub *gochannel.GoChannel
}

type P interface {
	AddPrompt(id string, renderer func(), removeOnRender bool)
	HasPrompt(id string) bool
	MoveTo(nextRoom R, onMoved func())
	QueueCommand(executable string, lag int) // CommandExecutable
	RemovePrompt(id string)
	EnterRoom(room R)
}

func NewPlayer() *Player {
	var em eventemitter.EventEmitter
	var ob eventemitter.Observable

	e := eventemitter.NewEmitter(true)
	em = e
	ob = e
	return &Player{
		EventEmitter: em,
		Observable:   ob,
	}
}

/*
Room#event:playerLeave
Room#event:playerEnter
Player#event:enterRoom
*/
func (p *Player) MoveTo(nextRoom *Room, onMoved func()) {
	prevRoom := p.Room
	if p.Room != nil && p.Room.ID != nextRoom.ID {
		p.Emit("playerLeave", prevRoom)
		p.Room.RemovePlayer(p)
	}

	p.Room = nextRoom
	nextRoom.AddPlayer(p)

	onMoved()

	nextRoom.Emit("room:playerEnter", p, prevRoom)
	// p.Emit("player:enterRoom", nextRoom)
	p.EventEmitter.EmitEvent(EventPlayerEnterRoom, nextRoom)
}

const (
	EventPlayerEnterRoom eventemitter.EventType = "player:enterRoom"
)

func (p *Player) Emit(event string, args ...interface{}) {
	log.Debug().
		Str("source", "player").
		Str("event", event).
		Msg("Emitting event")

	p.PubSub.Publish("player", message.NewMessage(watermill.NewUUID(), []byte(event)))
}

// type EventPlayerEnterRoom struct {
// 	PrevRoom *Room
// }
