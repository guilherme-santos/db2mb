package mysql

/*
#include "event.h"
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"time"
	"unsafe"

	"github.com/RideLink-carshare/db2mb"
)

type MySQLSchema struct {
	DB    string
	Table string
}

type MySQLEvent struct {
	ptr       unsafe.Pointer
	when      time.Time
	eventType db2mb.EventType
	tableID   uint
	MySQLSchema
}

var tables map[uint]MySQLSchema = make(map[uint]MySQLSchema)

func NewEvent(event unsafe.Pointer) db2mb.Event {
	ev := &MySQLEvent{}
	ev.ptr = event

	when := int64(C.Event_GetWhen(ev.ptr))
	ev.when = time.Unix(when, 0)

	var ok bool

	eventType := uint(C.Event_GetEventType(ev.ptr))
	ev.eventType, ok = events[eventType]
	if !ok {
		ev.eventType = db2mb.EventType{ID: eventType}
	}

	if ev.eventType == EventTypeTableMap {
		ev.tableID = uint(C.TableEvent_GetTableID(ev.ptr))

		schema := MySQLSchema{
			DB:    C.GoString(C.TableEvent_GetDBName(ev.ptr)),
			Table: C.GoString(C.TableEvent_GetTableName(ev.ptr)),
		}

		ev.MySQLSchema = schema
		tables[ev.tableID] = schema
	} else if ev.eventType == EventTypeQuery {
		ev.MySQLSchema = MySQLSchema{
			DB: C.GoString(C.QueryEvent_GetDBName(ev.ptr)),
		}
	} else if ev.eventType == EventTypePreGAWriteRows || ev.eventType == EventTypeWriteRows || ev.eventType == EventTypeWriteRowsV1 ||
		ev.eventType == EventTypePreGAUpdateRows || ev.eventType == EventTypeUpdateRows || ev.eventType == EventTypeUpdateRowsV1 ||
		ev.eventType == EventTypePreGADeleteRows || ev.eventType == EventTypeDeleteRows || ev.eventType == EventTypeDeleteRowsV1 {

		ev.tableID = uint(C.RowsEvent_GetTableID(ev.ptr))
		ev.MySQLSchema = tables[ev.tableID]
	}

	return ev
}

func (ev *MySQLEvent) GetWhen() time.Time {
	return ev.when
}

func (ev *MySQLEvent) GetType() db2mb.EventType {
	return ev.eventType
}

func (ev *MySQLEvent) JSON() map[string]interface{} {
	result := map[string]interface{}{
		"when": ev.when,
		"type": ev.eventType,
	}

	if ev.MySQLSchema.DB != "" {
		result["database"] = ev.MySQLSchema.DB
	}
	if ev.MySQLSchema.Table != "" {
		result["table"] = ev.MySQLSchema.Table
	}

	if ev.eventType == EventTypeQuery {
		result["query"] = C.GoString(C.QueryEvent_GetQuery(ev.ptr))
	} else if ev.eventType == EventTypePreGAWriteRows || ev.eventType == EventTypeWriteRows || ev.eventType == EventTypeWriteRowsV1 {
		result["operation"] = "insert"
	} else if ev.eventType == EventTypePreGAUpdateRows || ev.eventType == EventTypeUpdateRows || ev.eventType == EventTypeUpdateRowsV1 {
		result["operation"] = "update"
	} else if ev.eventType == EventTypePreGADeleteRows || ev.eventType == EventTypeDeleteRows || ev.eventType == EventTypeDeleteRowsV1 {
		result["operation"] = "delete"
	}

	return result
}

func (ev *MySQLEvent) String() string {
	if _, ok := events[ev.eventType.ID]; ok {
		C.Event_PrintEventInfo(ev.ptr)
	}

	j, _ := json.MarshalIndent(ev.JSON(), "", "    ")

	return fmt.Sprintf("%s\n--------------------------------------", string(j))
}
