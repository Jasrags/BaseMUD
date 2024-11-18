package core

type Inventory struct {
	IsFull  bool
	MaxSize int
}

func NewInventory(items []Item, max int) *Inventory {
	return &Inventory{}
}

func (i *Inventory) AddItem(item *Item) {
}

func (i *Inventory) GetMax() int {
	return i.MaxSize
}

func (i *Inventory) Hydrate(state string, carriedBy Character) {
}

func (i *Inventory) RemoveItem(item *Item) {
}

func (i *Inventory) SetMax(max int) {
	i.MaxSize = max
}
