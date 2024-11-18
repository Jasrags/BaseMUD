package core

import (
	"time"

	eventemitter "github.com/vansante/go-event-emitter"
)

const (
	EffectActivated   eventemitter.EventType = "Effect#Activated"
	EffectAdded       eventemitter.EventType = "Effect#Added"
	EffectDeactivated eventemitter.EventType = "Effect#Deactivated"
	EffectRefreshed   eventemitter.EventType = "Effect#Refreshed"
	EffectStackAdded  eventemitter.EventType = "Effect#StackAdded"
)

type Effect struct {
	Description string
	Duration    time.Duration
	Elapsed     time.Duration
	Name        string
	Remaining   time.Duration
}

func NewEffect() *Effect {
	return &Effect{}
}

// Effect#event:effectActivated
func (e *Effect) Activate() {
}

// Effect#event:effectDeactivated
func (e *Effect) Deactivate() {
}

func (e *Effect) Hydrate(state, data string) {
}

func (e *Effect) IsCurrent() bool {
	return false
}

func (e *Effect) ModifyAttribute(attrName string, currentValue int) int {
	return 0
}

func (e *Effect) ModifyIncomingDamage(damage *Damage, currentAmount int) *Damage {
	return nil
}

func (e *Effect) ModifyOutgoingDamage(damage *Damage, currentAmount int) *Damage {
	return nil
}

func (e *Effect) Pause() {

}

func (e *Effect) Remove() {

}

// Effect#event:remove
func (e *Effect) Resume() {

}

func (e *Effect) Serialize() string {
	return ""
}
