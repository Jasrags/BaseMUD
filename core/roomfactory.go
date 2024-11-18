package core

import eventemitter "github.com/vansante/go-event-emitter"

type RoomFactory interface {
	// Create(a Area, entityRef string) Room

	// EntityFactory
	AddScriptListener(id string, event eventemitter.EventType, listener eventemitter.HandleFunc)
	// clone(entity) → {Item|Npc|Room|Area}
	Create(area Area, id string) Room
	// createByType(area, entityRef, Type) → {type}
	// createEntityRef(areaName, id) → {string}
	// getDefinition(entityRef) → {Object}
	// setDefinition(entityRef, def)
}

type roomFactory struct {
}

// AddScriptListener implements RoomFactory.
func (r *roomFactory) AddScriptListener(id string, event eventemitter.EventType, listener eventemitter.HandleFunc) {
	panic("unimplemented")
}

// Create implements RoomFactory.
func (r *roomFactory) Create(area Area, id string) Room {
	panic("unimplemented")
}

func NewRoomFactory() RoomFactory {
	return &roomFactory{}
}
