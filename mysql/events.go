package mysql

import "github.com/RideLink-carshare/db2mb"

var (
	EventTypeUnknown            = db2mb.EventType{ID: 0, Name: "Unknown"}
	EventTypeStartV3            = db2mb.EventType{ID: 1, Name: "Start_v3"}
	EventTypeQuery              = db2mb.EventType{ID: 2, Name: "Query"}
	EventTypeStop               = db2mb.EventType{ID: 3, Name: "Stop"}
	EventTypeRotate             = db2mb.EventType{ID: 4, Name: "Rotate"}
	EventTypeIntvar             = db2mb.EventType{ID: 5, Name: "Intvar"}
	EventTypeLoad               = db2mb.EventType{ID: 6, Name: "Load"}
	EventTypeSlave              = db2mb.EventType{ID: 7, Name: "Slave"}
	EventTypeCreateFile         = db2mb.EventType{ID: 8, Name: "Create_file"}
	EventTypeAppendBlock        = db2mb.EventType{ID: 9, Name: "Append_block"}
	EventTypeExecLoad           = db2mb.EventType{ID: 10, Name: "Exec_load"}
	EventTypeDeleteFile         = db2mb.EventType{ID: 11, Name: "Delete_file"}
	EventTypeNewLoad            = db2mb.EventType{ID: 12, Name: "New_load"}
	EventTypeRand               = db2mb.EventType{ID: 13, Name: "RAND"}
	EventTypeUserVar            = db2mb.EventType{ID: 14, Name: "User_var"}
	EventTypeFormatDescription  = db2mb.EventType{ID: 15, Name: "Format_desc"}
	EventTypeXID                = db2mb.EventType{ID: 16, Name: "Xid"}
	EventTypeBeginLoadQuery     = db2mb.EventType{ID: 17, Name: "Begin_load_query"}
	EventTypeExecuteLoadQuery   = db2mb.EventType{ID: 18, Name: "Execute_load_query"}
	EventTypeTableMap           = db2mb.EventType{ID: 19, Name: "Table_map"}
	EventTypePreGAWriteRows     = db2mb.EventType{ID: 20, Name: "Write_rows_event_old"}
	EventTypePreGAUpdateRows    = db2mb.EventType{ID: 21, Name: "Update_rows_event_old"}
	EventTypePreGADeleteRows    = db2mb.EventType{ID: 22, Name: "Delete_rows_event_old"}
	EventTypeWriteRowsV1        = db2mb.EventType{ID: 23, Name: "Write_rows_v1"}
	EventTypeUpdateRowsV1       = db2mb.EventType{ID: 24, Name: "Update_rows_v1"}
	EventTypeDeleteRowsV1       = db2mb.EventType{ID: 25, Name: "Delete_rows_v1"}
	EventTypeIncident           = db2mb.EventType{ID: 26, Name: "Incident"}
	EventTypeHeartbeatLog       = db2mb.EventType{ID: 27, Name: "Heartbeat"}
	EventTypeIgnorableLog       = db2mb.EventType{ID: 28, Name: "Ignorable"}
	EventTypeRowsQueryLog       = db2mb.EventType{ID: 29, Name: "Rows_query"}
	EventTypeWriteRows          = db2mb.EventType{ID: 30, Name: "Write_rows"}
	EventTypeUpdateRows         = db2mb.EventType{ID: 31, Name: "Update_rows"}
	EventTypeDeleteRows         = db2mb.EventType{ID: 32, Name: "Delete_rows"}
	EventTypeGTIDLog            = db2mb.EventType{ID: 33, Name: "Gtid"}
	EventTypeAnonymousGTIDLog   = db2mb.EventType{ID: 34, Name: "Anonymous_Gtid"}
	EventTypePreviousGTIDSLog   = db2mb.EventType{ID: 35, Name: "Previous_gtids"}
	EventTypeTransactionContext = db2mb.EventType{ID: 36, Name: "Transaction_context"}
	EventTypeViewChange         = db2mb.EventType{ID: 37, Name: "View_change"}
)

var events = map[uint]db2mb.EventType{
	EventTypeUnknown.ID:            EventTypeUnknown,
	EventTypeStartV3.ID:            EventTypeStartV3,
	EventTypeQuery.ID:              EventTypeQuery,
	EventTypeStop.ID:               EventTypeStop,
	EventTypeRotate.ID:             EventTypeRotate,
	EventTypeIntvar.ID:             EventTypeIntvar,
	EventTypeLoad.ID:               EventTypeLoad,
	EventTypeSlave.ID:              EventTypeSlave,
	EventTypeCreateFile.ID:         EventTypeCreateFile,
	EventTypeAppendBlock.ID:        EventTypeAppendBlock,
	EventTypeExecLoad.ID:           EventTypeExecLoad,
	EventTypeDeleteFile.ID:         EventTypeDeleteFile,
	EventTypeNewLoad.ID:            EventTypeNewLoad,
	EventTypeRand.ID:               EventTypeRand,
	EventTypeUserVar.ID:            EventTypeUserVar,
	EventTypeFormatDescription.ID:  EventTypeFormatDescription,
	EventTypeXID.ID:                EventTypeXID,
	EventTypeBeginLoadQuery.ID:     EventTypeBeginLoadQuery,
	EventTypeExecuteLoadQuery.ID:   EventTypeExecuteLoadQuery,
	EventTypeTableMap.ID:           EventTypeTableMap,
	EventTypePreGAWriteRows.ID:     EventTypePreGAWriteRows,
	EventTypePreGAUpdateRows.ID:    EventTypePreGAUpdateRows,
	EventTypePreGADeleteRows.ID:    EventTypePreGADeleteRows,
	EventTypeWriteRowsV1.ID:        EventTypeWriteRowsV1,
	EventTypeUpdateRowsV1.ID:       EventTypeUpdateRowsV1,
	EventTypeDeleteRowsV1.ID:       EventTypeDeleteRowsV1,
	EventTypeIncident.ID:           EventTypeIncident,
	EventTypeHeartbeatLog.ID:       EventTypeHeartbeatLog,
	EventTypeIgnorableLog.ID:       EventTypeIgnorableLog,
	EventTypeRowsQueryLog.ID:       EventTypeRowsQueryLog,
	EventTypeWriteRows.ID:          EventTypeWriteRows,
	EventTypeUpdateRows.ID:         EventTypeUpdateRows,
	EventTypeDeleteRows.ID:         EventTypeDeleteRows,
	EventTypeGTIDLog.ID:            EventTypeGTIDLog,
	EventTypeAnonymousGTIDLog.ID:   EventTypeAnonymousGTIDLog,
	EventTypePreviousGTIDSLog.ID:   EventTypePreviousGTIDSLog,
	EventTypeTransactionContext.ID: EventTypeTransactionContext,
	EventTypeViewChange.ID:         EventTypeViewChange,
}
