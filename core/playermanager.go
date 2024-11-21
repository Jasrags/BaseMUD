package core

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
	Players map[string]*Player

	eventemitter.EventEmitter
	eventemitter.Observable
}

func NewPlayerManager(em eventemitter.EventEmitter, ob eventemitter.Observable) *PlayerManager {
	return &PlayerManager{
		Players: make(map[string]*Player),

		EventEmitter: em,
		Observable:   ob,
	}
}

// AddPlayer implements PlayerManager.
func (pm *PlayerManager) AddPlayer(p *Player) error {
	if _, ok := pm.Players[p.Name]; ok {
		log.Error().Msg("Player already registered")

		return errors.New("player already registered")
	}
	pm.Players[p.Name] = p

	return nil
}

// Exists implements PlayerManager.
func (pm *PlayerManager) Exists(id string) bool {
	_, ok := pm.Players[id]

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
func (pm *PlayerManager) GetPlayer(name string) (*Player, error) {
	if p, ok := pm.Players[strings.ToLower(name)]; ok {
		return p, nil
	}

	log.Error().
		Str("name", name).
		Err(errors.New("Player not found"))

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
	playersPath := viper.GetString("data.players_path")
	log.Info().
		Str("path", playersPath).
		Msg("Loading players")

	files, err := os.ReadDir(playersPath)
	if err != nil {
		log.Error().
			Str("path", playersPath).
			Err(err).
			Msg("Failed to read players directory")

		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(playersPath, file.Name())

			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Error().
					Str("file_path", filePath).
					Err(err).
					Msg("Failed to read file")
				continue
			}

			var p Player
			if err := json.Unmarshal(data, &p); err != nil {
				log.Error().
					Err(err).
					Str("file_path", filePath).
					Msg("Failed to unmarshal player data from file")
				continue
			}

			pm.Players[strings.ToLower(p.Name)] = &p
			log.Debug().
				Str("name", p.Name).
				Msg("Loaded player")
		}
	}

	log.Info().
		Int("count", len(pm.Players)).
		Msg("Players loaded")

	return nil
}

// RemovePlayer implements PlayerManager.
func (pm *PlayerManager) RemovePlayer(id string, killSocket bool) {
	delete(pm.Players, id)
}

// Save implements PlayerManager.
// Emits Player#event:save
func (pm *PlayerManager) Save() {
	pm.EmitEvent("Player#event:save", "Save")
	panic("unimplemented")
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
	panic("unimplemented")
}
