#include <binlog.h>
#include "binarylog.h"

extern "C" {

using namespace binary_log;

BinaryLog *NewBinaryLog(Driver *driver) {
    system::Binary_log_driver *drv = static_cast<system::Binary_log_driver*>(driver);
    Binary_log *binlog = new Binary_log(drv);

    return static_cast<BinaryLog*>(binlog);
}

struct resp_error_t BinaryLog_Connect(BinaryLog *self) {
    Binary_log *binlog = static_cast<Binary_log*>(self);
    struct resp_error_t error;

    int code = binlog->connect();
    if (code == ERR_OK) {
        error.ok = true;
        error.message = NULL;
    } else {
        error.ok = false;
        error.message = str_error(code);
    }

    return error;
}

struct resp_error_t BinaryLog_Disconnect(BinaryLog *self) {
    Binary_log *binlog = static_cast<Binary_log*>(self);
    struct resp_error_t error;

    int code = binlog->disconnect();
    if (code == ERR_OK) {
        error.ok = true;
        error.message = NULL;
    } else {
        error.ok = false;
        error.message = str_error(code);
    }

    return error;
}

} // extern "C"
