package db2mb

type EventLog interface {
	Connect() error
	Disconnect() error
	SetPosition(pos uint) error
	WaitForEvent() (Event, error)
}

type EventType

type Event interface {
    GetType() int
}
