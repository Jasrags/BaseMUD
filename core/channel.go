package core

type (
	Audience string
)

const (
// AreaAudience    Audience = "area"
// PartyAudience   Audience = "party"
// PrivateAudience Audience = "private"
// RoomAudience    Audience = "room"
// WorldAudience   Audience = "world"
)

// type Channel interface {
// 	Name() string
// 	Color() []string
// 	Description() string
// 	MinRequiredRole() PlayerRole
// 	Audience() Audience
// 	SenderFormat() string
// 	TargetFormat() string

// 	FormatToReceipient(sender, target, message string) string
// 	FormatToSender(sender, message string) string
// 	Send(state, sender, message string)
// }

type Channel struct {
}
