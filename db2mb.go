package db2mb

import "time"

type EventLog interface {
	Connect() error
	Disconnect() error
	GetPosition() uint
	SetPosition(pos uint) error
	WaitForEvent() (Event, error)
}

type EventType struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Event interface {
	GetWhen() time.Time
	GetType() EventType
	JSON() map[string]interface{}
	String() string
}
