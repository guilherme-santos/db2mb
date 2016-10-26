#ifndef _MYSQL_BINARYLOG_H_
#define _MYSQL_BINARYLOG_H_

#include <stdbool.h>
#include "driver.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef void BinaryLog;

struct resp_error_t {
    bool ok;
    const char *message;
};

BinaryLog *NewBinaryLog(Driver *driver);

struct resp_error_t BinaryLog_Connect(BinaryLog *self);
struct resp_error_t BinaryLog_Disconnect(BinaryLog *self);

#ifdef __cplusplus
}
#endif
#endif
