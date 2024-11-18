package core

import (
	"errors"

	"github.com/rs/zerolog/log"
	eventemitter "github.com/vansante/go-event-emitter"
)

// type PlayerManager interface {
// 	AddPlayer(p Player) error
// 	Exists(id string) bool
// 	Filter(fn func()) []Player
// 	GetBroadcastTargets() []Character
// 	GetPlayer(id string) (Player, error)
// 	GetPlayersAsArray() []Player
// 	Keyify(p Player) string
// 	LoadPlayer(state string, a Account, username string, force bool) Player
// 	LoadPlayers() error
// 	RemovePlayer(id string, killSocket bool)
// 	Save()
// 	SaveAll()
// 	TickAll()

// 	eventemitter.EventEmitter
// 	eventemitter.Observable
// }

/*
Listens to Events:

	PlayerManager#event:save
	PlayerManager#event:updateTick
*/
type PlayerManager struct {
	players map[string]*Player

	eventemitter.EventEmitter
	eventemitter.Observable
}

func NewPlayerManager(em eventemitter.EventEmitter, ob eventemitter.Observable) *PlayerManager {
	return &PlayerManager{
		players: make(map[string]*Player),

		EventEmitter: em,
		Observable:   ob,
	}
}

// AddPlayer implements PlayerManager.
func (pm *PlayerManager) AddPlayer(p *Player) error {
	if _, ok := pm.players[p.Name]; ok {
		log.Error().Msg("Player already registered")

		return errors.New("player already registered")
	}
	pm.players[p.Name] = p

	return nil
}

// Exists implements PlayerManager.
func (pm *PlayerManager) Exists(id string) bool {
	_, ok := pm.players[id]

	return ok
}

// Filter implements PlayerManager.
func (pm *PlayerManager) Filter(fn func()) []*Player {
	panic("unimplemented")
}

// GetBroadcastTargets implements PlayerManager.
func (pm *PlayerManager) GetBroadcastTargets() []Character {
	panic("unimplemented")
}

// GetPlayer implements PlayerManager.
func (pm *PlayerManager) GetPlayer(id string) (*Player, error) {
	if p, ok := pm.players[id]; ok {
		return p, nil
	}

	log.Error().Err(errors.New("Player not found"))

	return nil, errors.New("Player not found")
}

// GetPlayersAsArray implements PlayerManager.
func (pm *PlayerManager) GetPlayersAsArray() []*Player {
	panic("unimplemented")
}

// Keyify implements PlayerManager.
func (pm *PlayerManager) Keyify(p *Player) string {
	panic("unimplemented")
}

// LoadPlayer implements PlayerManager.
func (pm *PlayerManager) LoadPlayer(state string, a *Account, username string, force bool) *Player {
	panic("unimplemented")
}

func (pm *PlayerManager) LoadPlayers() error {
	log.Info().Msg("Loading players")
	pm.players = map[string]*Player{
		"test": {
			// EventEmitter: pm.EventEmitter,
			// Observable:   pm.Observable,
			// Name:         "test",
			Account: &Account{
				Id:         "1",
				Username:   "test",
				Characters: []string{"test"},
				Password:   "test",
				Banned:     false,
			},
		},
	}

	return nil
}

// RemovePlayer implements PlayerManager.
func (pm *PlayerManager) RemovePlayer(id string, killSocket bool) {
	delete(pm.players, id)
}

// Save implements PlayerManager.
// Emits Player#event:save
func (pm *PlayerManager) Save() {
	pm.EmitEvent("Player#event:save", "Save")
}

// SaveAll implements PlayerManager.
// Emits Player#event:saved
func (pm *PlayerManager) SaveAll() {
	pm.EmitEvent("Player#event:saved", "SaveAll")
	panic("unimplemented")
}

// Emits Player#event:updateTick
func (pm *PlayerManager) TickAll() {
	pm.EmitEvent("Player#event:updateTick", "TickAll")
}
