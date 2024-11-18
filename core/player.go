package core

import (
	"github.com/gliderlabs/ssh"
	eventemitter "github.com/vansante/go-event-emitter"
)

const (
	PlayerAttributeUpdated eventemitter.EventType = "Player#event:attributeUpdated"
	PlayerCombatantAdded   eventemitter.EventType = "Player#event:combatantAdded"
	PlayerCombatantRemoved eventemitter.EventType = "Player#event:combatantRemoved"
	PlayerCombatEnd        eventemitter.EventType = "Player#event:combatEnd"
	PlayerCombatStart      eventemitter.EventType = "Player#event:combatStart"
	PlayerDamaged          eventemitter.EventType = "Player#event:damaged"
	PlayerEffectAdded      eventemitter.EventType = "Player#event:effectAdded"
	PlayerEffectRemoved    eventemitter.EventType = "Player#event:effectRemoved"
	PlayerEnterRoom        eventemitter.EventType = "Player#event:enterRoom"
	PlayerEquip            eventemitter.EventType = "Player#event:equip"
	PlayerFollowed         eventemitter.EventType = "Player#event:followed"
	PlayerGainedFollower   eventemitter.EventType = "Player#event:gainedFollower"
	PlayerHeal             eventemitter.EventType = "Player#event:heal"
	PlayerHealed           eventemitter.EventType = "Player#event:healed"
	PlayerHit              eventemitter.EventType = "Player#event:hit"
	PlayerLostFollower     eventemitter.EventType = "Player#event:lostFollower"
	PlayerSaved            eventemitter.EventType = "Player#event:saved"
	PlayerUnquip           eventemitter.EventType = "Player#event:unquip"
	PlayerUnfollowed       eventemitter.EventType = "Player#event:unfollowed"
	PlayerUpdateTick       eventemitter.EventType = "Player#event:updateTick"
)

type PlayerRole int

const (
	RolePlayer PlayerRole = iota
	RoleBuilder
	RoleAdmin
)

// type Player interface {
// 	Character

// 	AddPrompt(id string, renderer func(), removeOnRender bool)
// 	Emit(eventName string, args ...interface{})
// 	GetName() string
// 	GetSession() ssh.Session
// 	HasPrompt(id string) bool
// 	MoveTo(nextRoom Room, onMoved func())
// 	QueueCommand(executable string, lag int) // CommandExecutable
// 	RemovePrompt(id string)
// 	EnterRoom(room Room)

// 	eventemitter.EventEmitter
// 	eventemitter.Observable
// }

// Character
// name 	string
// inventory 	Inventory
// combatants 	Set
// level 	number
// attributes 	Attributes
// effects 	EffectList
// room 	Room

// Player
// account 	Account
// experience 	number
// password 	string
// prompt 	string
// socket 	net.Socket
// questTracker 	QuestTracker
// extraPrompts 	Map.<string, function()>
// questData 	Object

// Room the character is currently in

type Player struct {
	Account    *Account
	Experience int
	Password   string
	Prompt     string
	Session    ssh.Session
	//  QuestTracker
	// ExtraPrompts
	// QuestData

	// pubSub *gochannel.GoChannel

	Character
}

func NewPlayer(s ssh.Session, em eventemitter.EventEmitter, ob eventemitter.Observable) *Player {
	p := &Player{
		Session: s,
	}

	p.EventEmitter = em
	p.Observable = ob

	return p
}

// AddPrompt implements Player.
func (p *Player) AddPrompt(id string, renderer func(), removeOnRender bool) {
	panic("unimplemented")
}

// // GetBroadcastTargets implements Player.
// func (p *Player) GetBroadcastTargets() {
// 	panic("unimplemented")
// }

// HasPrompt implements Player.
func (p *Player) HasPrompt(id string) bool {
	panic("unimplemented")
}

// Hydrate implements Player.
func (p *Player) Hydrate(state string) {
	panic("unimplemented")
}

func (p *Player) InterpolatePrompt(promptStr, extraData string) {

}

func (p *Player) IsNpc() bool {
	return false
}

// MoveTo implements Player.
// Emits
// Room#event:playerLeave
// Room#event:playerEnter
// Player#event:enterRoom
func (p *Player) MoveTo(nextRoom *Room, onMoved func()) {
	prevRoom := p.Room

	if p.Room != nil && p.Room.GetID() != nextRoom.GetID() {
		p.EmitEvent(RoomPlayerLeave, prevRoom.GetID(), p.Name)
		// p.Emit(RoomPlayerLeave, prevRoom.GetID())
		// p.room.Emit(string(RoomPlayerLeave), prevRoom.GetID())
		p.Room.RemovePlayer(p)
	}

	p.Room = nextRoom
	nextRoom.AddPlayer(p)

	onMoved()

	nextRoom.EmitEvent(RoomPlayerEnter, p.Room.GetID(), p.Name)
	// nextRoom.Emit("Room#event:playerEnter", p.room.GetID())
	p.EmitEvent(PlayerEnterRoom, nextRoom.GetID())
	// p.Emit("Player#event:enterRoom", nextRoom.GetID())
}

// QueueCommand implements Player.
func (p *Player) QueueCommand(executable string, lag int) {
	panic("unimplemented")
}

// RemovePrompt implements Player.
func (p *Player) RemovePrompt(id string) {
	panic("unimplemented")
}

// Serialize implements Player.
func (p *Player) Serialize() string {
	return ""
}
