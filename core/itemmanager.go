package core

type ItemManager struct {
}

func NewItemManager() *ItemManager {
	return &ItemManager{}
}

// Emits Item#event:updateTick
func (im *ItemManager) TickAll() {

}
