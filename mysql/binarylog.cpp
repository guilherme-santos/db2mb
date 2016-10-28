#include <binlog.h>
#include "binarylog.h"

extern "C" {

using namespace binary_log;

BinaryLog *NewBinaryLog(Driver *driver) {
    system::Binary_log_driver *drv = static_cast<system::Binary_log_driver*>(driver);
    Binary_log *binlog = new Binary_log(drv);

    return static_cast<BinaryLog*>(binlog);
}

struct resp_t BinaryLog_Connect(BinaryLog *self) {
    Binary_log *binlog = static_cast<Binary_log*>(self);
    struct resp_t resp;

    int code = binlog->connect();
    if (code == ERR_OK) {
        resp.ok = true;
        resp.message = NULL;
    } else {
        resp.ok = false;
        resp.message = str_error(code);
    }

    return resp;
}

struct resp_t BinaryLog_Disconnect(BinaryLog *self) {
    Binary_log *binlog = static_cast<Binary_log*>(self);
    struct resp_t resp;

    int code = binlog->disconnect();
    if (code == ERR_OK) {
        resp.ok = true;
        resp.message = NULL;
    } else {
        resp.ok = false;
        resp.message = str_error(code);
    }

    return resp;
}

unsigned long BinaryLog_GetPosition(BinaryLog *self) {
    Binary_log *binlog = static_cast<Binary_log*>(self);
    struct resp_t resp;

    return binlog->get_position();
}

struct resp_t BinaryLog_SetPosition(BinaryLog *self, unsigned long position) {
    Binary_log *binlog = static_cast<Binary_log*>(self);
    struct resp_t resp;

    int code = binlog->set_position(position);
    if (code == ERR_OK) {
        resp.ok = true;
        resp.message = NULL;
    } else {
        resp.ok = false;
        resp.message = str_error(code);
    }

    return resp;
}

} // extern "C"
