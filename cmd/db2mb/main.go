package main

import (
	"github.com/NeowayLabs/logger"
	"github.com/RideLink-carshare/db2mb"
	"github.com/RideLink-carshare/db2mb/mysql"
)

func main() {
	logger.Info("Trying to connect to Remote Event Log...")

	var eventlog db2mb.EventLog
	eventlog = mysql.NewRemoteEventLog("127.0.0.1", "root", "", 3306)

	err := eventlog.Connect()
	if err != nil {
		logger.Fatal("Cannot connect to Remote Event Log: %v", err)
	}

	logger.Info("Connected to Remote Event Log!")

	err = eventlog.SetPosition(4) // TODO we need to save which was last position
	if err != nil {
		logger.Fatal("Cannot set start position: %v", err)
	}

	logger.Info("Waiting for new events...\n")

	ignoreEvents := map[db2mb.EventType]struct{}{
		db2mb.EventType{ID: 9800704}:     struct{}{},
		db2mb.EventType{ID: 9796608}:     struct{}{},
		mysql.EventTypeAnonymousGTIDLog:  struct{}{},
		mysql.EventTypeXID:               struct{}{},
		mysql.EventTypeFormatDescription: struct{}{},
		mysql.EventTypePreviousGTIDSLog:  struct{}{},

		// mysql.EventTypeQuery: struct{}{},
		// mysql.EventTypeTableMap:   struct{}{},
		// mysql.EventTypeWriteRows:  struct{}{},
		// mysql.EventTypeUpdateRows: struct{}{},
		// mysql.EventTypeDeleteRows: struct{}{},
	}

	for {
		event, err := eventlog.WaitForEvent()
		if err != nil {
			logger.Fatal("Cannot read event: %v", err)
		}

		eventType := event.GetType()
		if _, ok := ignoreEvents[eventType]; ok {
			continue
		}

		logger.Info("-> New event: #%02d - %s", eventType.ID, eventType.Name)
		logger.Info("Payload:\n%s", event.String())
	}
}
