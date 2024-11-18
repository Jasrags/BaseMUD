package core

import eventemitter "github.com/vansante/go-event-emitter"

const (
	AreaUpdateTick eventemitter.EventType = "Area#event:updateTick"
)

type AreaManager struct {
	Areas map[string]*Area
}

func NewAreaManager() *AreaManager {
	return &AreaManager{}
}

func (am *AreaManager) AddArea(area *Area) {
}

func (am *AreaManager) GetAreaByReference(entityRef string) *Area {
	return nil
}

func (am *AreaManager) GetPlaceholderArea() *Area {
	return nil
}

func (am *AreaManager) RemoveArea(area *Area) {
}

// Emits Area#event:updateTick
func (am *AreaManager) TickAll(state string) {
}
