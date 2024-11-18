package core

import (
	"time"

	eventemitter "github.com/vansante/go-event-emitter"
)

const (
	CharacterAttributeUpdate  eventemitter.EventType = "Character#attributeUpdate"
	CharacterCombatantAdded   eventemitter.EventType = "Character#combatantAdded"
	CharacterCombatantRemoved eventemitter.EventType = "Character#combatantRemoved"
	CharacterCombatEnd        eventemitter.EventType = "Character#combatEnd"
	CharacterCombatStart      eventemitter.EventType = "Character#combatStart"
	CharacterDamaged          eventemitter.EventType = "Character#damaged"
	CharacterEffectAdded      eventemitter.EventType = "Character#effectAdded"
	CharacterEffectRemoved    eventemitter.EventType = "Character#effectRemoved"
	CharacterEquip            eventemitter.EventType = "Character#equip"
	CharacterFollowed         eventemitter.EventType = "Character#followed"
	CharacterGainedFollower   eventemitter.EventType = "Character#gainedFollower"
	CharacterHeal             eventemitter.EventType = "Character#heal"
	CharacterHealed           eventemitter.EventType = "Character#healed"
	CharacterHit              eventemitter.EventType = "Character#hit"
	CharacterLostFollower     eventemitter.EventType = "Character#lostFollower"
	CharacterUnquip           eventemitter.EventType = "Character#unquip"
	CharacterUnfollowed       eventemitter.EventType = "Character#unfollowed"
)

// type Character interface {
// 	AddAttribute(name string)
// 	AddCombatant(target Character)
// 	AddEffect(effect Effect) bool
// 	AddFollower(follower Character)
// 	AddItem(item Item)
// 	Equip(item Item, slot string) error
// 	EvaluateIncomingDamage(damage *Damage) int
// 	EvaluateOutgoingDamage(damage *Damage, currentAmount int)
// 	Follow(target Character)
// 	GetAttribute(attr string) int
// 	GetBaseAttribute(attr string) int
// 	GetBroadcastTargets()
// 	GetMaxAttribute(attr string) int
// 	HasAttribute(attr string) bool
// 	HasEffectType(effectType string) bool
// 	HasFollower(follower Character) bool
// 	HasItem(itemReference string) Item
// 	Hydrate(state string)
// 	InitiateCombat(target Character, lag time.Duration)
// 	IsFollowing(target Character) bool
// 	IsInCombat(target Character) bool
// 	IsInventoryFull() bool
// 	IsNpc() bool
// 	LowerAttribute(attr string, amount int)
// 	RaiseAttribute(attr string, amount int)
// 	RemoveComabatant(target Character)
// 	RemoveEfffect(effect Effect)
// 	RemoveFollower(follower Character)
// 	RemoveFromCombat()
// 	RemoveItem(item Item)
// 	Serialize() string
// 	SetAttributeBase(attr string, newBase int)
// 	SetAttributeToMax(attr string)
// 	Unequip(slot string) error
// 	Unfollow()
// }

type Character struct {
	Name       string
	Inventory  *Inventory
	Combatants []*Character
	Level      int
	// Attributes Attributes
	// Effects    EffectList
	Room *Room
	// Id string
	// account Account
	//
	//	experience int
	//	password string
	//	prompt string
	//
	// socker ssh.Session
	//
	//	questTracker
	//
	// extraPrompts
	// questData

	eventemitter.EventEmitter
	eventemitter.Observable
}

func NewCharacter() *Character {
	return &Character{}
}

// AddAttribute implements Character.
func (c *Character) AddAttribute(name string) {
	panic("unimplemented")
}

// AddCombatant implements Character.
// Emits Character#event:combatantAdded
func (c *Character) AddCombatant(target *Character) {
	panic("unimplemented")
}

// AddEffect implements Character.
func (c *Character) AddEffect(effect Effect) bool {
	panic("unimplemented")
}

// AddFollower implements Character.
// Emits
// Character#event:gainedFollower
func (c *Character) AddFollower(follower *Character) {
	panic("unimplemented")
}

// AddItem implements Character.
func (c *Character) AddItem(item *Item) {
	panic("unimplemented")
}

// Equip implements Character.
// Emits
// Character#event:equip
// Item#event:equip
func (c *Character) Equip(item *Item, slot string) error {
	panic("unimplemented")
}

// EvaluateIncomingDamage implements Character.
func (c *Character) EvaluateIncomingDamage(damage *Damage) int {
	panic("unimplemented")
}

// EvaluateOutgoingDamage implements Character.
func (c *Character) EvaluateOutgoingDamage(damage *Damage, currentAmount int) {
	panic("unimplemented")
}

// Follow implements Character.
func (c *Character) Follow(target *Character) {
	panic("unimplemented")
}

// GetAttribute implements Character.
func (c *Character) GetAttribute(attr string) int {
	panic("unimplemented")
}

// GetBaseAttribute implements Character.
func (c *Character) GetBaseAttribute(attr string) int {
	panic("unimplemented")
}

// GetBroadcastTargets implements Character.
func (c *Character) GetBroadcastTargets() {
	panic("unimplemented")
}

// GetMaxAttribute implements Character.
func (c *Character) GetMaxAttribute(attr string) int {
	panic("unimplemented")
}

// HasAttribute implements Character.
func (c *Character) HasAttribute(attr string) bool {
	panic("unimplemented")
}

// HasEffectType implements Character.
func (c *Character) HasEffectType(effectType string) bool {
	panic("unimplemented")
}

// HasFollower implements Character.
func (c *Character) HasFollower(follower *Character) bool {
	panic("unimplemented")
}

// HasItem implements Character.
func (c *Character) HasItem(itemReference string) *Item {
	panic("unimplemented")
}

// InitiateCombat implements Character.
// Emits Character#event:combatStart
func (c *Character) InitiateCombat(target *Character, lag time.Duration) {
	panic("unimplemented")
}

// IsFollowing implements Character.
func (c *Character) IsFollowing(target Character) bool {
	panic("unimplemented")
}

// IsInCombat implements Character.
func (c *Character) IsInCombat(target Character) bool {
	panic("unimplemented")
}

// IsInventoryFull implements Character.
func (c *Character) IsInventoryFull() bool {
	panic("unimplemented")
}

// LowerAttribute implements Character.
// Emits Character#event:attributeUpdate
func (c *Character) LowerAttribute(attr string, amount int) {
	panic("unimplemented")
}

// RaiseAttribute implements Character.
// Emits Character#event:attributeUpdate
func (c *Character) RaiseAttribute(attr string, amount int) {
	panic("unimplemented")
}

// RemoveComabatant implements Character.
// Emits
// Character#event:combatantRemoved
// Character#event:combatEnd
func (c *Character) RemoveComabatant(target Character) {
	panic("unimplemented")
}

// RemoveEfffect implements Character.
func (c *Character) RemoveEfffect(effect Effect) {
	panic("unimplemented")
}

// RemoveFollower implements Character.
// Emits Character#event:lostFollower
func (c *Character) RemoveFollower(follower Character) {
	panic("unimplemented")
}

// RemoveFromCombat implements Character.
func (c *Character) RemoveFromCombat() {
	panic("unimplemented")
}

// RemoveItem implements Character.
func (c *Character) RemoveItem(item *Item) {
	panic("unimplemented")
}

// Serialize implements Character.
func (c *Character) Serialize() string {
	panic("unimplemented")
}

// SetAttributeBase implements Character.
// Emits Character#event:attributeUpdate
func (c *Character) SetAttributeBase(attr string, newBase int) {
	panic("unimplemented")
}

// SetAttributeToMax implements Character.
// Emits Character#event:attributeUpdate
func (c *Character) SetAttributeToMax(attr string) {
	panic("unimplemented")
}

// Unequip implements Character.
// Emits
// Item#event:unequip
// Character#event:unequip
func (c *Character) Unequip(slot string) error {
	panic("unimplemented")
}

// Unfollow implements Character.
// Emits Character#event:unfollowed
func (c *Character) Unfollow() {
	panic("unimplemented")
}
