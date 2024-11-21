package core

import (
	"errors"
	"fmt"
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
	Rooms map[string]*Room

	eventEmitter  eventemitter.EventEmitter
	eventObserver eventemitter.Observable
}

func NewRoomManager(em eventemitter.EventEmitter, ob eventemitter.Observable) *RoomManager {
	return &RoomManager{
		Rooms: make(map[string]*Room),

		eventEmitter:  em,
		eventObserver: ob,
	}
}

func (rm *RoomManager) AddRoom(room *Room) error {
	roomId := fmt.Sprintf("%s:%s", room.Area.Name, room.Id)

	log.Debug().Str("id", roomId).Msg("Adding room")
	if _, ok := rm.Rooms[roomId]; ok {
		log.Error().Msg("Room already registered")
		return errors.New("room already registered")
	}

	rm.Rooms[roomId] = room

	return nil
}

// GetRoom implements RoomManager.
func (rm *RoomManager) GetRoom(id string) (*Room, error) {
	log.Debug().Str("id", id).Msg("Getting room")

	if room, ok := rm.Rooms[id]; ok {
		return room, nil
	}

	log.Error().Msg("Room not found")
	return nil, errors.New("room not found")
}

func (rm *RoomManager) LoadRoom(id string) (*Room, error) {
	panic("unimplemented")
}

func (rm *RoomManager) LoadRooms() error {
	areaPath := viper.GetString("data.areas_path")
	log.Info().Str("area_path", areaPath).Msg("Loading rooms")

	files, err := os.ReadDir(areaPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read area directory")
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			areaName := file.Name()
			log.Info().Str("area", areaName).Msg("Loading area")

			manifestPath := filepath.Join(areaPath, file.Name(), viper.GetString("data.manifest_file"))
			roomsPath := filepath.Join(areaPath, file.Name(), viper.GetString("data.rooms_file"))

			manifestData, err := os.ReadFile(manifestPath)
			if err != nil {
				log.Error().Err(err).Str("path", manifestPath).Msg("Failed to read manifest file")
				continue
			}

			var area Area
			if err = yaml.Unmarshal(manifestData, &area); err != nil {
				log.Error().Err(err).Str("path", manifestPath).Msg("Failed to unmarshal manifest file")
				continue
			}

			roomsData, err := os.ReadFile(roomsPath)
			if err != nil {
				log.Error().Err(err).Str("path", roomsPath).Msg("Failed to read rooms file")
				continue
			}

			// var rooms map[string]*Room
			var rooms []*Room
			if err = yaml.Unmarshal(roomsData, &rooms); err != nil {
				log.Error().Err(err).Str("path", roomsPath).Msg("Failed to unmarshal rooms file")
				continue
			}

			for _, room := range rooms {
				log.Debug().Str("room", room.Id).Msg("Loading room")
				room.Init(rm.eventEmitter, rm.eventObserver)
				room.Area = &area

				if err := rm.AddRoom(room); err != nil {
					log.Error().Err(err).Str("room", room.Id).Msg("Failed to add room")
				}
				// rm.Rooms[fmt.Sprintf("%s:%s", room.Area.Name, room.Id)] = room
			}
			log.Info().Int("room_count", len(rm.Rooms)).Msg("Area loaded")
		}
	}

	return nil
}

func (rm *RoomManager) Save() error {
	areaPath := viper.GetString("data.area_path")
	roomsData, err := yaml.Marshal(rm.Rooms)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal rooms data")
		return err
	}

	roomsFilePath := filepath.Join(areaPath, "the_void/rooms.yaml")
	err = os.WriteFile(roomsFilePath, roomsData, 0644)
	if err != nil {
		log.Error().Err(err).Str("path", roomsFilePath).Msg("Failed to write rooms file")
		return err
	}

	log.Info().Str("path", roomsFilePath).Msg("Rooms data saved successfully")

	return nil
}

// RemoveRoom implements RoomManager.
func (rm *RoomManager) RemoveRoom(id string) {
	delete(rm.Rooms, id)
}
