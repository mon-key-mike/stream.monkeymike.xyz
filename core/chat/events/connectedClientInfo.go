package events

import "github.com/owncast/owncast/core/user"

// ConnectedClientInfo represents the information about a connected client.
type ConnectedClientInfo struct {
	User *user.User `json:"user"`
	Event
}
