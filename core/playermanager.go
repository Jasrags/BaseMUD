package core

type PlayerManager struct {
	Players map[string]*Player
	//Events EventManager
	//Loader EntityLoader
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		Players: make(map[string]*Player),
	}
}

func (pm *PlayerManager) AddListener(BehaviorName string, Listener func()) {

}

func (pm *PlayerManager) AddPlayer(p *Player) {
	pm.Players[p.ID] = p
}

func (pm *PlayerManager) Exists(id string) bool {
	_, ok := pm.Players[id]
	return ok
}

// Filter

// getBroadcastTargets

func (pm *PlayerManager) GetPlayer(id string) *Player {
	return pm.Players[id]
}

func (pm *PlayerManager) GetPlayers() []*Player {
	var players []*Player
	for _, p := range pm.Players {
		players = append(players, p)
	}
	return players
}

func (pm *PlayerManager) LoadPlayer(state string, account Account, username string, force bool) *Player {
	return nil
}

func (pm *PlayerManager) RemovePlayer(id string, killSocket bool) {
	delete(pm.Players, id)
}
