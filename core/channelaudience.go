package core

type ChannelAudience struct {
	State   string
	Sender  string
	Message string
}

func NewChannelAudience(state, sender, message string) *ChannelAudience {
	return &ChannelAudience{}
}

func (ca *ChannelAudience) Configure(state, sender, message string) {
	ca.State = state
	ca.Sender = sender
	ca.Message = message
}

func (ca *ChannelAudience) AlterMessage(msg string) string {
	return msg
}

func (ca *ChannelAudience) GetBroadcastTargets() []*Player {
	//     return this.state.PlayerManager.getPlayersAsArray();
	return []*Player{}
}
