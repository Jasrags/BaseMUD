package core

// type Item interface {
// 	AddItem(item Item)
// 	Close()
// 	FindCarrier() //→ {Character|Item|null}
// 	InitializeInventory(inventory string)
// 	IsInventoryFull() bool
// 	Lock()
// 	Open()
// 	RemoveItem(item Item)
// 	Unlock()
// }

type Item struct {
	// area Area
	// // metadata 	object
	// // behaviors 	Array
	// description string
	// id          string
	// isEquipped  bool
	// equippedBy  Character
	// inventory   map[string]Item
	// name        string
	// room        Room
	// roomDesc    string
	// script      string
	// // itemType ItemType | string
	// uuid      string
	// closeable bool
	// closed    bool
	// locked    bool
	// lockedBy  string
	// // carriedBy  Character | Item
}

func NewItem() *Item {
	return &Item{}
}

// AddItem implements Item.
func (i *Item) AddItem(item *Item) {
	panic("unimplemented")
}

// Close implements Item.
func (i *Item) Close() {
	panic("unimplemented")
}

// FindCarrier implements Item.
func (i *Item) FindCarrier() {
	panic("unimplemented")
}

// InitializeInventory implements Item.
func (i *Item) InitializeInventory(inventory string) {
	panic("unimplemented")
}

// IsInventoryFull implements Item.
func (i *Item) IsInventoryFull() bool {
	panic("unimplemented")
}

// Lock implements Item.
func (i *Item) Lock() {
	panic("unimplemented")
}

// Open implements Item.
func (i *Item) Open() {
	panic("unimplemented")
}

// RemoveItem implements Item.
func (i *Item) RemoveItem(item *Item) {
	panic("unimplemented")
}

// Unlock implements Item.
func (i *Item) Unlock() {
	panic("unimplemented")
}
