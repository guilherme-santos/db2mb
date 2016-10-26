package main

type EventLog interface {
	Connect() error
	Disconnect() error
}
