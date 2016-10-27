package db2mb

import "time"

type EventLog interface {
	Connect() error
	Disconnect() error
	SetPosition(pos uint) error
	WaitForEvent() (Event, error)
}

type EventType struct {
	ID   uint
	Name string
}

type Event interface {
	GetWhen() time.Time
	GetType() EventType
	JSON() map[string]interface{}
	String() string
}
