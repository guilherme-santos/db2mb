#ifndef _MYSQL_EVENT_H_
#define _MYSQL_EVENT_H_

#ifdef __cplusplus
extern "C" {
#endif

typedef void Event;

unsigned int Event_GetEventType(Event *self);
unsigned long Event_GetWhen(Event *self);
void Event_PrintEventInfo(Event *self);

const char *QueryEvent_GetDBName(Event *self);
const char *QueryEvent_GetQuery(Event *self);

const char *TableEvent_GetDBName(Event *self);
const char *TableEvent_GetTableName(Event *self);
unsigned int TableEvent_GetTableID(Event *self);

unsigned int RowsEvent_GetTableID(Event *self);


#ifdef __cplusplus
}
#endif
#endif
