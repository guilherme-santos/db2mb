#include <ostream>
#include <binlog.h>
#include "event.h"

extern "C" {

using namespace binary_log;

unsigned int Event_GetEventType(Event *self) {
    Binary_log_event *event = static_cast<Binary_log_event*>(self);

    return event->get_event_type();
}

unsigned long Event_GetWhen(Event *self) {
    Binary_log_event *event = static_cast<Binary_log_event*>(self);

    return event->header()->when.tv_sec;
}

void Event_PrintEventInfo(Event *self) {
    Binary_log_event *event = static_cast<Binary_log_event*>(self);
    if (event->get_event_type() == 9683968) {
        return;
    }

    std::cout << "--------------------------------------" << std::endl;
    std::cout << "Event info: " << event->get_event_type() << std::endl;
    event->print_long_info(std::cout);
    std::cout << std::endl;
    std::cout << "----" << std::endl;
}

const char *TableEvent_GetDBName(Event *self) {
    Table_map_event *event = static_cast<Table_map_event*>(self);

    return event->m_dbnam.c_str();
}

const char *TableEvent_GetTableName(Event *self) {
    Table_map_event *event = static_cast<Table_map_event*>(self);

    return event->m_tblnam.c_str();
}

unsigned int TableEvent_GetTableID(Event *self) {
    Table_map_event *event = static_cast<Table_map_event*>(self);

    return event->get_table_id();
}

unsigned int RowsEvent_GetTableID(Event *self) {
    Rows_event *event = static_cast<Rows_event*>(self);

    return event->get_table_id();
}

} // extern "C"
