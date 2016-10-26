package main

import (
	"github.com/NeowayLabs/logger"
	"github.com/RideLink-carshare/db2mb/mysql"
)

func main() {
	logger.Info("Trying to connect to Remote Event Log...")

	eventlog := mysql.NewRemoteEventLog("127.0.0.1", "root", "", 3306)

	err := eventlog.Connect()
	if err != nil {
		logger.Fatal("Cannot connect to Remote Event Log: %v", err)
	}

	logger.Info("Connected to Remote Event Log!")

	// TODO set initial position (we need to save which was last position)
	// TODO wait for new events
}
