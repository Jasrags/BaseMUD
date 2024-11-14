package core

type Character struct {
	ID   string
	Name string
	// Inventory Inventory
	// Combatants set
	// Level int
	// Attributes Attributes
	// Effects    EffectList
	Room *Room
	// IsNPC bool
}

type C interface {
	AddAttribute(name string)
	AddCombatant(target C)
	AddEffect(effect Effect)
	AddFollower(follower C)
	AddItem(item Item)
	Emit(event string, args []string)
	Equip(item Item, slot string)
	EvaluateIncomingDamage(damage Damage)
	EvaluateOutgoingDamage(damage Damage, currentAmount int)
	Follow(target C)
	GetAttribute(attr string) int
	GetBaseAttribute(attr string) int
	GetBroadcastTargets()
	GetMaxAttribute(attr string) int
	HasAttribute(attr string) bool
	HasEffectType(effectType string) bool
	HasFollower(follower C) bool
	HasItem(itemReference string) Item
	Hydrate(state string)
	initiateCombat(target C, lag int)
	IsFollowing(target C) bool
	IsInCombat(target C) bool
	LowerAttribute(attr string, amount int)
	RaiseAttribute(attr string, amount int)
	RemoveComabatant(target C)
	RemoveEfffect(effect Effect)
	RemoveFollower(follower C)
	RemoveFromCombat()
	Serialize() string
	SetBaseAttribute(attr string, newBase int)
	SetAttributeToMax(attr string)
	Unequip(slot string)
	Unfollow()
}
