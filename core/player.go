package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
	// Account    *Account
	// Experience int
	// Password string
	Id      string      `json:"id"`
	Prompt  string      `json:"prompt"`
	Session ssh.Session `json:"-"`
	Name    string      `json:"name"`
	RoomId  string      `json:"room_id"`
	Room    *Room       `json:"-"`
	// Inventory  *Inventory
	// Combatants []*Character
	// Level      int
	// Attributes Attributes
	// Effects    EffectList
	//  QuestTracker
	// ExtraPrompts
	// QuestData

	// pubSub *gochannel.GoChannel

	// Character

	listeners                 []*eventemitter.Listener `json:"-"`
	eventemitter.EventEmitter `json:"-"`
	eventemitter.Observable   `json:"-"`
}

func (p *Player) Init(em eventemitter.EventEmitter, ob eventemitter.Observable) {
	p.EventEmitter = em
	p.Observable = ob

	// Setup listeners
	// p.listeners = append(p.listeners,
	// 	p.AddListener(
	// 		eventemitter.EventType(RoomPlayerEnter),
	// 		eventemitter.HandleFunc(p.HandlePlayerEnterEvent)))
	// p.listeners = append(r.listeners,
	// 	p.AddListener(
	// 		eventemitter.EventType(RoomPlayerLeave),
	// 		eventemitter.HandleFunc(p.HandlePlayerLeaveEvent)))
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
	panic("unimplemented")
}

func (p *Player) IsNpc() bool {
	return false
}

// MoveTo
// Emits
// Room#event:playerLeave
// Room#event:playerEnter
// Player#event:enterRoom
func (p *Player) MoveTo(nextRoom *Room, onMoved func()) {
	prevRoom := p.Room

	if p.Room != nil && !strings.EqualFold(p.Room.Id, nextRoom.Id) {
		p.EventEmitter.EmitEvent(RoomPlayerLeave, p, nextRoom)
		p.Room.RemovePlayer(p)
	}

	p.Room = nextRoom
	nextRoom.AddPlayer(p)

	onMoved()

	nextRoom.EmitEvent(RoomPlayerEnter, p, prevRoom)
	p.EventEmitter.EmitEvent(PlayerEnterRoom, nextRoom)
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

func (p *Player) Save() {
	var sb strings.Builder

	e := json.NewEncoder(&sb)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")

	if err := e.Encode(p); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to encode player to JSON")
		return
	}

	filePath := fmt.Sprintf("%s/%s.json",
		viper.GetString("data.players_path"), strings.ToLower(p.Name))

	if err := os.WriteFile(filePath, []byte(sb.String()), 0644); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to write player to file")
		return
	}
}

// character

// AddAttribute implements Character.
func (p *Player) AddAttribute(name string) {
	panic("unimplemented")
}

// AddCombatant implements Character.
// Emits Character#event:combatantAdded
func (p *Player) AddCombatant(target *Character) {
	panic("unimplemented")
}

// AddEffect implements Character.
func (p *Player) AddEffect(effect Effect) bool {
	panic("unimplemented")
}

// AddFollower implements Character.
// Emits
// Character#event:gainedFollower
func (p *Player) AddFollower(follower *Character) {
	panic("unimplemented")
}

// AddItem implements Character.
func (p *Player) AddItem(item *Item) {
	panic("unimplemented")
}

// Equip implements Character.
// Emits
// Character#event:equip
// Item#event:equip
func (p *Player) Equip(item *Item, slot string) error {
	panic("unimplemented")
}

// EvaluateIncomingDamage implements Character.
func (p *Player) EvaluateIncomingDamage(damage *Damage) int {
	panic("unimplemented")
}

// EvaluateOutgoingDamage implements Character.
func (p *Player) EvaluateOutgoingDamage(damage *Damage, currentAmount int) {
	panic("unimplemented")
}

// Follow implements Character.
func (p *Player) Follow(target *Character) {
	panic("unimplemented")
}

// GetAttribute implements Character.
func (p *Player) GetAttribute(attr string) int {
	panic("unimplemented")
}

// GetBaseAttribute implements Character.
func (p *Player) GetBaseAttribute(attr string) int {
	panic("unimplemented")
}

// GetBroadcastTargets implements Character.
func (p *Player) GetBroadcastTargets() {
	panic("unimplemented")
}

// GetMaxAttribute implements Character.
func (p *Player) GetMaxAttribute(attr string) int {
	panic("unimplemented")
}

// HasAttribute implements Character.
func (p *Player) HasAttribute(attr string) bool {
	panic("unimplemented")
}

// HasEffectType implements Character.
func (p *Player) HasEffectType(effectType string) bool {
	panic("unimplemented")
}

// HasFollower implements Character.
func (p *Player) HasFollower(follower *Character) bool {
	panic("unimplemented")
}

// HasItem implements Character.
func (p *Player) HasItem(itemReference string) *Item {
	panic("unimplemented")
}

// InitiateCombat implements Character.
// Emits Character#event:combatStart
func (p *Player) InitiateCombat(target *Character, lag time.Duration) {
	panic("unimplemented")
}

// IsFollowing implements Character.
func (p *Player) IsFollowing(target Character) bool {
	panic("unimplemented")
}

// IsInCombat implements Character.
func (p *Player) IsInCombat(target Character) bool {
	panic("unimplemented")
}

// IsInventoryFull implements Character.
func (p *Player) IsInventoryFull() bool {
	panic("unimplemented")
}

// LowerAttribute implements Character.
// Emits Character#event:attributeUpdate
func (p *Player) LowerAttribute(attr string, amount int) {
	panic("unimplemented")
}

// RaiseAttribute implements Character.
// Emits Character#event:attributeUpdate
func (p *Player) RaiseAttribute(attr string, amount int) {
	panic("unimplemented")
}

// RemoveComabatant implements Character.
// Emits
// Character#event:combatantRemoved
// Character#event:combatEnd
func (p *Player) RemoveComabatant(target Character) {
	panic("unimplemented")
}

// RemoveEfffect implements Character.
func (p *Player) RemoveEfffect(effect Effect) {
	panic("unimplemented")
}

// RemoveFollower implements Character.
// Emits Character#event:lostFollower
func (p *Player) RemoveFollower(follower Character) {
	panic("unimplemented")
}

// RemoveFromCombat implements Character.
func (p *Player) RemoveFromCombat() {
	panic("unimplemented")
}

// RemoveItem implements Character.
func (p *Player) RemoveItem(item *Item) {
	panic("unimplemented")
}

// SetAttributeBase implements Character.
// Emits Character#event:attributeUpdate
func (p *Player) SetAttributeBase(attr string, newBase int) {
	panic("unimplemented")
}

// SetAttributeToMax implements Character.
// Emits Character#event:attributeUpdate
func (p *Player) SetAttributeToMax(attr string) {
	panic("unimplemented")
}

// Unequip implements Character.
// Emits
// Item#event:unequip
// Character#event:unequip
func (p *Player) Unequip(slot string) error {
	panic("unimplemented")
}

// Unfollow implements Character.
// Emits Character#event:unfollowed
func (p *Player) Unfollow() {
	panic("unimplemented")
}
