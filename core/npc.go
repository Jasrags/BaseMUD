package core

// extends Character
type NPC struct {
	Character
	ID   int
	Area Area
	// Behaviors map[string]
}
