package core

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	eventemitter "github.com/vansante/go-event-emitter"
)

// type RoomManager interface {
// 	AddRoom(r Room) error
// 	GetRoom(id string) (Room, error)
// 	RemoveRoom(id string)
// 	LoadRoom(id string) (Room, error)
// 	LoadRooms() error
// }

type RoomManager struct {
	rooms map[string]*Room

	eventEmitter  eventemitter.EventEmitter
	eventObserver eventemitter.Observable
}

func NewRoomManager(em eventemitter.EventEmitter, ob eventemitter.Observable) *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),

		eventEmitter:  em,
		eventObserver: ob,
	}
}

// AddRoom implements RoomManager.
func (rm *RoomManager) AddRoom(room *Room) error {
	if _, ok := rm.rooms[room.GetID()]; ok {
		log.Error().Msg("Room already registered")
		return errors.New("room already registered")
	}

	rm.rooms[room.GetID()] = room

	return nil
}

// GetRoom implements RoomManager.
func (rm *RoomManager) GetRoom(id string) (*Room, error) {
	if room, ok := rm.rooms[id]; ok {
		return room, nil
	}

	log.Error().Msg("Room not found")
	return nil, errors.New("room not found")
}

func (rm *RoomManager) LoadRoom(id string) (*Room, error) {
	panic("unimplemented")
}

func (rm *RoomManager) LoadRooms() error {
	areaPath := viper.GetString("data.area_path")
	log.Info().Str("area_path", areaPath).Msg("Loading rooms")

	files, err := os.ReadDir(areaPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read area directory")
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			manifestPath := filepath.Join(areaPath, file.Name(), "manifest.yaml")
			roomsPath := filepath.Join(areaPath, file.Name(), "rooms.yaml")

			manifestData, err := os.ReadFile(manifestPath)
			if err != nil {
				log.Error().Err(err).Str("path", manifestPath).Msg("Failed to read manifest file")
				continue
			}

			var area Area
			err = yaml.Unmarshal(manifestData, &area)
			if err != nil {
				log.Error().Err(err).Str("path", manifestPath).Msg("Failed to unmarshal manifest file")
				continue
			}

			roomsData, err := os.ReadFile(roomsPath)
			if err != nil {
				log.Error().Err(err).Str("path", roomsPath).Msg("Failed to read rooms file")
				continue
			}

			var rooms map[string]*Room
			err = yaml.Unmarshal(roomsData, &rooms)
			if err != nil {
				log.Error().Err(err).Str("path", roomsPath).Msg("Failed to unmarshal rooms file")
				continue
			}

			for id, room := range rooms {
				room.EventEmitter = rm.eventEmitter
				room.Observable = rm.eventObserver
				rm.rooms[id] = room
			}
		}
	}

	rm.rooms = map[string]*Room{
		"the_void:the_void": {
			EventEmitter: rm.eventEmitter,
			Observable:   rm.eventObserver,
			id:           "the_void:the_void",
			players:      make(map[string]*Player),
			title:        "The Void",
			coordinates:  []int{0, 0, 0},
			description:  "You don't think that you are not floating in nothing.",
			// npcs: make(map[string]NPC),
			// items: make(map[string]Item),
			// exits: make(map[string]Room),
		},
		"the_void:limbo": {
			EventEmitter: rm.eventEmitter,
			Observable:   rm.eventObserver,
			id:           "the_void:limbo",
			players:      make(map[string]*Player),
			title:        "Limbo",
			coordinates:  []int{0, 0, 1},
			description:  "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
		},
	}

	return nil
}

// RemoveRoom implements RoomManager.
func (rm *RoomManager) RemoveRoom(id string) {
	delete(rm.rooms, id)
}
