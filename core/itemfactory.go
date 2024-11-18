package core

type ItemFactory struct {
}

func NewItemFactory() *ItemFactory {
	return &ItemFactory{}
}

func (ifac *ItemFactory) CreateItem(area *Area, entityRef string) *Item {
	return &Item{}
}
