package core

import eventemitter "github.com/vansante/go-event-emitter"

// type EventManager interface {
// Add(eventName string, listener eventemitter.Listener)
// Attach(emitter eventemitter.EventEmitter, config string)
// Detach(emitter eventemitter.EventEmitter, events ...string)
// Get(name string)
// }

const (
	ChannelReceive eventemitter.EventType = "channelReceive"
	UpdateTick     eventemitter.EventType = "updateTick"
)

type EventManager struct {
	// listeners map[string]eventemitter.Listener
}

func NewEventManager() *EventManager {
	return &EventManager{}
}

// type Listener func(data interface{})

// type EventEmitter interface {
// 	On(eventName string, listener Listener)
// 	Emit(eventName string, data interface{})
// }

// func (em *eventManager) On(eventName string, listener Listener) {
// 	if em.listeners == nil {
// 		em.listeners = make(map[string][]Listener)
// 	}
// 	em.listeners[eventName] = append(em.listeners[eventName], listener)
// }

// func (em *eventManager) Emit(eventName string, data interface{}) {
// 	if listeners, found := em.listeners[eventName]; found {
// 		for _, listener := range listeners {
// 			listener(data)
// 		}
// 	}
// }

// type eventManager struct {
// 	listeners map[string][]Listener
// 	mu        sync.RWMutex
// }

// func NewEventManager() EventManager {
// 	return &eventManager{
// 		listeners: make(map[string][]Listener),
// 	}
// }
