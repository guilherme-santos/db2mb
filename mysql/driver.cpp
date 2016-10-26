#include <binlog.h>
#include "driver.h"

extern "C" {

using namespace binary_log;
using namespace binary_log::system;

Decoder decoder(0);

Driver *NewDriver(const char *url) {
    Binary_log_driver *driver = create_transport(url);
    return static_cast<Driver*>(driver);
}

struct resp_t Driver_GetNextEvent(Driver *self) {
    Binary_log_driver *driver = static_cast<Binary_log_driver*>(self);
    struct resp_t resp;
    std::pair<unsigned char *, size_t> buffer_buflen;

    int code = driver->get_next_event(&buffer_buflen);
    if (code == ERR_OK) {
        Binary_log_event *event;
        const char *error = NULL;

        event = decoder.decode_event((char *) buffer_buflen.first, buffer_buflen.second, &error, 1);
        if (event != NULL) {
            resp.ok = true;

            if (event->get_event_type() == binary_log::INCIDENT_EVENT ||
                    (event->get_event_type() == binary_log::ROTATE_EVENT &&
                    event->header()->log_pos == 0 ||
                    event->header()->log_pos == 0))
            {
                return resp;
            }

            resp.data = event;

            return resp;
        }

        resp.ok = false;
        resp.message = error;
    } else {
        resp.ok = false;
        resp.message = str_error(code);
    }

    return resp;
}

} // extern "C"
